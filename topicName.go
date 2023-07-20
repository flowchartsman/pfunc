//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package pfunc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TopicName abstract a struct contained in a Topic
type TopicName struct {
	Domain       string
	Namespace    string
	Name         string
	LocalName    string
	Partition    int
	partitionIdx int
}

func (tn *TopicName) WithoutPartition() string {
	return tn.Name[:tn.partitionIdx]
}

const (
	publicTenant           = "public"
	defaultNamespace       = "default"
	partitionedTopicSuffix = "-partition-"
)

// ParseTopicName parse the given topic name and return TopicName.
func ParseTopicName(topic string) (*TopicName, error) {
	// The topic name can be in two different forms, one is fully qualified topic name,
	// the other one is short topic name
	if !strings.Contains(topic, "://") {
		// The short topic name can be:
		// - <topic>
		// - <tenant>/<namespace>/<topic>
		// - <tenant>/<cluster>/<namespace>/<topic>
		parts := strings.Split(topic, "/")
		if len(parts) == 3 || len(parts) == 4 {
			topic = "persistent://" + topic
		} else if len(parts) == 1 {
			topic = "persistent://" + publicTenant + "/" + defaultNamespace + "/" + parts[0]
		} else {
			return nil, errors.New(
				"Invalid short topic name '" + topic +
					"', it should be in the format of <tenant>/<namespace>/<topic> or <topic>")
		}
	}

	// The fully qualified topic name can be in two different forms:
	// new:    persistent://tenant/namespace/topic
	// legacy: persistent://tenant/cluster/namespace/topic
	parts := strings.SplitN(topic, "://", 2)
	domain := parts[0]
	if domain != "persistent" && domain != "non-persistent" {
		return nil, errors.New("Invalid topic domain: " + domain)
	}

	tn := &TopicName{
		Domain:       domain,
		Name:         topic,
		Partition:    -1,
		partitionIdx: len(topic),
	}

	rest := parts[1]
	var err error

	// The rest of the name can be in different forms:
	// new:    tenant/namespace/<localName>
	// legacy: tenant/cluster/namespace/<localName>
	// Examples of localName:
	// 1. some/name/xyz//
	// 2. /xyz-123/feeder-2
	parts = strings.SplitN(rest, "/", 4)
	if len(parts) == 3 {
		// New topic name without cluster name
		tn.Namespace = parts[0] + "/" + parts[1]
		tn.LocalName = parts[2]
	} else if len(parts) == 4 {
		// Legacy topic name that includes cluster name
		tn.Namespace = fmt.Sprintf("%s/%s/%s", parts[0], parts[1], parts[2])
		tn.LocalName = parts[3]
	} else {
		return nil, errors.New("invalid topic name: " + topic)
	}

	pIdx := strings.LastIndex(topic, partitionedTopicSuffix)
	if pIdx >= 0 {
		tn.partitionIdx = pIdx
		pIdx += len(partitionedTopicSuffix)
		if pIdx >= len(topic) {
			return nil, errors.New("missing partition number")
		}
		tn.Partition, err = strconv.Atoi(topic[pIdx:])
		if err != nil {
			return nil, fmt.Errorf("invalid partition number: %q", topic[pIdx:])
		}
		tn.LocalName = tn.LocalName[:strings.LastIndex(tn.LocalName, partitionedTopicSuffix)]
	}
	return tn, nil
}
