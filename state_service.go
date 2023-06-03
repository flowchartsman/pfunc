package pfunc

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"andy.dev/pfunc/internal/bookkeeper/kv"
	"andy.dev/pfunc/internal/pb/bookkeeper/storage"
	"andy.dev/pfunc/internal/pb/bookkeeper/stream"
)

type stateService interface {
	GetCounter(key string) (int64, error)
	IncrCounter(key string, amount int64) (int64, error)
	DelCounter(key string) error
	GetState(key string) ([]byte, error)
	PutState(key string, value []byte) error
	DelState(key string) error
}

type bkStateService struct {
	client *kv.Client
	table  *kv.Table
}

// need to call from init with instanceConfig if stateStorageServiceURL is set.
// else no-op
func newbkStateService(instanceConf *instanceConf) (*bkStateService, error) {
	u, err := url.Parse(instanceConf.stateServiceURL)
	if err != nil {
		return nil, fmt.Errorf("parsing state storage url: %v", err)
	}
	kvNamespace := strings.ReplaceAll(fmt.Sprintf("%s_%s", instanceConf.funcDetails.Tenant, instanceConf.funcDetails.Namespace), "-", "_")
	client, err := kv.NewClient(context.TODO(), u.Host, kvNamespace)
	if err != nil {
		return nil, fmt.Errorf("bookkeeper KV client init: %v", err)
	}

	table, err := client.GetTable(context.TODO(), instanceConf.funcDetails.Name)
	// TODO: need to fix this make a proper sentinel error
	if err != nil {
		if se, ok := err.(kv.StorageError); ok {
			if se.Code() == storage.StatusCode_STREAM_NOT_FOUND {
				// https://github.com/apache/pulsar/blob/v2.5.0/pulsar-functions/instance/src/main/java/org/apache/pulsar/functions/instance/JavaInstanceRunnable.java#L337-L340
				table, err = client.CreateTable(context.Background(), instanceConf.funcDetails.Name, &stream.StreamConfiguration{
					MinNumRanges:     4,
					InitialNumRanges: 4,
					StorageType:      stream.StorageType_TABLE,
					KeyType:          stream.RangeKeyType_HASH,
					RetentionPolicy: &stream.RetentionPolicy{
						TimePolicy: &stream.TimeBasedRetentionPolicy{
							RetentionMinutes: -1,
						},
					},
				})
				if err != nil {
					return nil, fmt.Errorf("creating BK table: %v", err)
				}
			}
		}
	}
	return &bkStateService{
		client: client,
		table:  table,
	}, nil
}

// TODO: Consider removing temlpate or using default?

func (bs *bkStateService) GetCounter(key string) (int64, error) {
	return bs.table.GetInt(context.Background(), key)
}

func (bs *bkStateService) IncrCounter(key string, amount int64) (int64, error) {
	return bs.table.Incr(context.Background(), key, amount)
}

func (bs *bkStateService) DelCounter(key string) error {
	return bs.table.Delete(context.Background(), key)
}

func (bs *bkStateService) GetState(key string) ([]byte, error) {
	return bs.table.Get(context.Background(), key)
}

func (bs *bkStateService) PutState(key string, value []byte) error {
	return bs.table.Put(context.Background(), key, value)
}

func (bs *bkStateService) DelState(key string) error {
	return bs.table.Delete(context.Background(), key)
}

var ErrStateServiceNotEnabled = errors.New("state service not enabled")

type noopStateService struct{}

func (_ noopStateService) GetCounter(_ string) (int64, error) {
	return 0, ErrStateServiceNotEnabled
}

func (_ noopStateService) IncrCounter(_ string, _ int64) (int64, error) {
	return 0, ErrStateServiceNotEnabled
}

func (_ noopStateService) DelCounter(_ string) error {
	return ErrStateServiceNotEnabled
}

func (_ noopStateService) GetState(_ string) ([]byte, error) {
	return nil, ErrStateServiceNotEnabled
}

func (_ noopStateService) PutState(_ string, _ []byte) error {
	return ErrStateServiceNotEnabled
}

func (_ noopStateService) DelState(_ string) error {
	return ErrStateServiceNotEnabled
}
