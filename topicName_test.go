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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTopicName(t *testing.T) {
	testCases := []struct {
		in        string
		name      string
		namespace string
		localName string
		noPart    string
		partition int
	}{
		{
			in:        "persistent://my-tenant/my-ns/my-topic1",
			name:      "persistent://my-tenant/my-ns/my-topic1",
			namespace: "my-tenant/my-ns",
			localName: "my-topic1",
			partition: -1,
		},
		{
			in:        "my-topic2",
			name:      "persistent://public/default/my-topic2",
			namespace: "public/default",
			localName: "my-topic2",
			partition: -1,
		},
		{
			in:        "my-tenant/my-namespace/my-topic3",
			name:      "persistent://my-tenant/my-namespace/my-topic3",
			namespace: "my-tenant/my-namespace",
			localName: "my-topic3",
			partition: -1,
		},
		{
			in:        "non-persistent://my-tenant/my-namespace/my-topic4",
			name:      "non-persistent://my-tenant/my-namespace/my-topic4",
			namespace: "my-tenant/my-namespace",
			localName: "my-topic4",
			partition: -1,
		},
		{
			in:        "my-topic5-partition-5",
			name:      "persistent://public/default/my-topic5-partition-5",
			noPart:    "persistent://public/default/my-topic5",
			namespace: "public/default",
			localName: "my-topic5",
			partition: 5,
		},
		{
			in:        "persistent://mytenant/mynamespace/my-topic6-partition-0",
			name:      "persistent://mytenant/mynamespace/my-topic6-partition-0",
			noPart:    "persistent://mytenant/mynamespace/my-topic6",
			namespace: "mytenant/mynamespace",
			localName: "my-topic6",
			partition: 0,
		},
		// V1 topic name
		{
			in:        "persistent://my-tenant/my-cluster/my-ns/my-topic7",
			name:      "persistent://my-tenant/my-cluster/my-ns/my-topic7",
			namespace: "my-tenant/my-cluster/my-ns",
			localName: "my-topic7",
			partition: -1,
		},
		{
			in:        "my-tenant/my-cluster/my-ns/my-topic8",
			name:      "persistent://my-tenant/my-cluster/my-ns/my-topic8",
			namespace: "my-tenant/my-cluster/my-ns",
			localName: "my-topic8",
			partition: -1,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.in, func(t *testing.T) {
			topic, err := ParseTopicName(testCase.in)
			assert.Nil(t, err)
			assert.Equal(t, testCase.name, topic.Name)
			assert.Equal(t, testCase.namespace, topic.Namespace)
			assert.Equal(t, testCase.partition, topic.Partition)
			assert.Equal(t, testCase.localName, topic.LocalName)
			if testCase.noPart == "" {
				assert.Equal(t, topic.WithoutPartition(), topic.Name)
			} else {
				assert.Equal(t, testCase.noPart, topic.WithoutPartition())
			}
		})
	}
}

func TestParseTopicNameErrors(t *testing.T) {
	testCases := []string{
		"invalid://my-tenant/my-ns/my-topic",
		"invalid://my-tenant/my-ns/my-topic-partition-xyz",
		"my-tenant/my-ns/my-topic-partition-xyz/invalid",
		"persistent://my-tenant",
		"persistent://my-tenant/my-namespace",
		"persistent://my-tenant/my-cluster/my-ns/my-topic-partition-xyz/invalid",
	}
	for _, testCase := range testCases {
		t.Run(testCase, func(t *testing.T) {
			_, err := ParseTopicName(testCase)
			assert.NotNil(t, err)
		})
	}
}
