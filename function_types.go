package pfunc

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/pulsar-client-go/pulsar"
)

// InputMessage is the incoming message from the selected topic.
type InputMessage = pulsar.Message

// TopicMsg is an outgoing message on a topic other than the primary
// topic.
type TopicMsg struct {
	topic string
	msg   pulsar.ProducerMessage
}

func TopicOutput(data any, topic string) *TopicMsg {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Output - unable to marshal value of type %T to JSON: %v", data, err))
	}
	return &TopicMsg{
		topic: topic,
		msg: pulsar.ProducerMessage{
			Payload: bytes,
		},
	}
}

func TopicMessage(payload []byte, topic string) *TopicMsg {
	return &TopicMsg{
		topic: topic,
		msg: pulsar.ProducerMessage{
			Payload: payload,
		},
	}
}

// WithKey sets the key on a TopicMessage
func (tm *TopicMsg) WithKey(key string) *TopicMsg {
	tm.msg.Key = key
	return tm
}

// WithProperty sets a single property on a TopicMessage
func (tm *TopicMsg) WithProperty(key, value string) *TopicMsg {
	tm.msg.Properties[key] = value
	return tm
}

// WithProperties sets multiple properties on a TopicMessage. WithProperties will panic if it receives an odd number of arguments.
func (tm *TopicMsg) WithProperties(keyvals ...string) *TopicMsg {
	if len(keyvals)%2 != 0 {
		panic("WithProperties - odd number of property keyvals")
	}
	for i := 0; i < len(keyvals); i += 2 {
		tm.msg.Properties[keyvals[i]] = keyvals[i+1]
	}
	return tm
}

// FnOutput is the output from a single function iteration.
type FnOutput struct {
	primary       *TopicMsg
	topicMessages []*TopicMsg
}

func Output(data any) *FnOutput {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Output - unable to marshal value of type %T to JSON: %v", data, err))
	}
	return &FnOutput{
		primary: &TopicMsg{
			msg: pulsar.ProducerMessage{
				Payload: bytes,
			},
		},
	}
}

func OutputMessage(payload []byte) *FnOutput {
	return &FnOutput{
		primary: &TopicMsg{
			msg: pulsar.ProducerMessage{
				Payload: payload,
			},
		},
	}
}

// WithKey sets the key on a the primary output message.
func (fo *FnOutput) WithKey(key string) *FnOutput {
	fo.primary.WithKey(key)
	return fo
}

// WithProperty sets a single property on the primary output message.
func (fo *FnOutput) WithProperty(key, value string) *FnOutput {
	fo.primary.WithProperty(key, value)
	return fo
}

// WithProperties sets multiple properties on the primary output message.
// WithProperties will panic if it receives an odd number of arguments.
func (fo *FnOutput) WithProperties(keyvals ...string) *FnOutput {
	fo.primary.WithProperties(keyvals...)
	return fo
}

func (fo *FnOutput) WithMessages(additionalMessages ...*TopicMsg) {
	fo.topicMessages = append(fo.topicMessages, additionalMessages...)
}

// type SourceFn func() ([])

type (
	FnProcessor func(ctx context.Context, msg InputMessage) (*FnOutput, error)
)

type Processor interface {
	~func(input []byte) error |
		~func(input []byte) (any, error) |
		~func(ctx context.Context, input []byte) error |
		~func(ctx context.Context, input []byte) (any, error) |
		~func(ctx context.Context, msg InputMessage) (*FnOutput, error)
}

func wrapFunction[P Processor](f P) FnProcessor {
	if f == nil {
		panic("cannot accept nil function")
	}
	switch fn := any(f).(type) {
	case func(input []byte) error:
		return func(ctx context.Context, msg InputMessage) (*FnOutput, error) {
			return nil, fn(msg.Payload())
		}
	case func(input []byte) (any, error):
		return func(_ context.Context, msg InputMessage) (*FnOutput, error) {
			data, err := fn(msg.Payload())
			if err != nil {
				return nil, err
			}
			switch d := data.(type) {
			case []byte:
				return OutputMessage(d), nil
			default:
				return Output(d), nil
			}
		}
	case func(ctx context.Context, input []byte) error:
		return func(ctx context.Context, msg InputMessage) (*FnOutput, error) {
			return nil, fn(ctx, msg.Payload())
		}
	case func(ctx context.Context, input []byte) (any, error):
		return func(ctx context.Context, msg InputMessage) (*FnOutput, error) {
			data, err := fn(ctx, msg.Payload())
			if err != nil {
				return nil, err
			}
			switch d := data.(type) {
			case []byte:
				return OutputMessage(d), nil
			default:
				return Output(d), nil
			}
		}
	case func(ctx context.Context, msg InputMessage) (*FnOutput, error):
		return FnProcessor(fn)
	default:
		panic("invalid function type")
	}
}
