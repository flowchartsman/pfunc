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

package conf

import (
	"encoding/json"
	"errors"
	"flag"
	"io/fs"
	"os"

	log "andy.dev/pfunc/logutil"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	PulsarServiceURL       string `json:"pulsarServiceURL" yaml:"pulsarServiceURL"`
	StateStorageServiceURL string `json:"stateStorageServiceUrl" yaml:"stateStorageServiceUrl"`
	PulsarWebServiceURL    string `json:"pulsarWebServiceUrl" yaml:"pulsarWebServiceUrl"`
	InstanceID             int    `json:"instanceID" yaml:"instanceID"`
	FuncID                 string `json:"funcID" yaml:"funcID"`
	FuncVersion            string `json:"funcVersion" yaml:"funcVersion"`
	MaxBufTuples           int    `json:"maxBufTuples" yaml:"maxBufTuples"`
	Port                   int    `json:"port" yaml:"port"`
	ClusterName            string `json:"clusterName" yaml:"clusterName"`
	KillAfterIdleMs        int    `json:"killAfterIdleMs" yaml:"killAfterIdleMs"`
	// function details config
	Tenant               string `json:"tenant" yaml:"tenant"`
	NameSpace            string `json:"nameSpace" yaml:"nameSpace"`
	Name                 string `json:"name" yaml:"name"`
	LogTopic             string `json:"logTopic" yaml:"logTopic"`
	ProcessingGuarantees int32  `json:"processingGuarantees" yaml:"processingGuarantees"`
	SecretsMap           string `json:"secretsMap" yaml:"secretsMap"`
	Runtime              int32  `json:"runtime" yaml:"runtime"`
	// Authentication
	ClientAuthenticationPlugin     string `json:"clientAuthenticationPlugin" yaml:"clientAuthenticationPlugin"`
	ClientAuthenticationParameters string `json:"clientAuthenticationParameters" yaml:"clientAuthenticationParameters"`
	TLSTrustCertsFilePath          string `json:"tlsTrustCertsFilePath" yaml:"tlsTrustCertsFilePath"`
	TLSAllowInsecureConnection     bool   `json:"tlsAllowInsecureConnection" yaml:"tlsAllowInsecureConnection"`
	TLSHostnameVerificationEnable  bool   `json:"tlsHostnameVerificationEnable" yaml:"tlsHostnameVerificationEnable"`
	// Deprecated
	AutoACK     bool  `json:"autoAck" yaml:"autoAck"`
	Parallelism int32 `json:"parallelism" yaml:"parallelism"`
	// source config
	SubscriptionType     int32  `json:"subscriptionType" yaml:"subscriptionType"`
	TimeoutMs            uint64 `json:"timeoutMs" yaml:"timeoutMs"`
	SubscriptionName     string `json:"subscriptionName" yaml:"subscriptionName"`
	CleanupSubscription  bool   `json:"cleanupSubscription"  yaml:"cleanupSubscription"`
	SubscriptionPosition int32  `json:"subscriptionPosition" yaml:"subscriptionPosition"`
	// source input specs
	SourceInputSpecs map[string]string `json:"sourceInputSpecs" yaml:"sourceInputSpecs"`
	// for backward compatibility
	// Deprecated
	SourceSpecTopic string `json:"sourceSpecsTopic" yaml:"sourceSpecsTopic"`
	// Deprecated
	SourceSchemaType string `json:"sourceSchemaType" yaml:"sourceSchemaType"`
	// Deprecated
	IsRegexPatternSubscription bool `json:"isRegexPatternSubscription" yaml:"isRegexPatternSubscription"`
	// Deprecated
	ReceiverQueueSize int32 `json:"receiverQueueSize" yaml:"receiverQueueSize"`
	// sink spec config
	SinkSpecTopic  string `json:"sinkSpecsTopic" yaml:"sinkSpecsTopic"`
	SinkSchemaType string `json:"sinkSchemaType" yaml:"sinkSchemaType"`
	// resources config
	CPU  float64 `json:"cpu" yaml:"cpu"`
	RAM  int64   `json:"ram" yaml:"ram"`
	Disk int64   `json:"disk" yaml:"disk"`
	// retryDetails config
	MaxMessageRetries           int32  `json:"maxMessageRetries" yaml:"maxMessageRetries"`
	DeadLetterTopic             string `json:"deadLetterTopic" yaml:"deadLetterTopic"`
	ExpectedHealthCheckInterval int32  `json:"expectedHealthCheckInterval" yaml:"expectedHealthCheckInterval"`
	UserConfig                  string `json:"userConfig" yaml:"userConfig"`
	// metrics config
	MetricsPort int `json:"metricsPort" yaml:"metricsPort"`
}

var (
	help         bool
	confFilePath string
	confContent  string
)

func initConfig() {
	if flag.Parsed() {
		return
	}
	flag.BoolVar(&help, "help", false, "print help cmd")
	flag.StringVar(&confFilePath, "instance-conf-path", "", "config conf.yml filepath")
	flag.StringVar(&confContent, "instance-conf", "", "the string content of Conf struct")
	flag.Parse()

	if help {
		flag.Usage()
	}
}

func GetConf() *Conf {
	initConfig()
	// prefer, in order:
	// - manual configuration file
	// - manual configuration as flag
	// - automatic configuration file

	if confFilePath == "" {
		confFilePath = getDefaultConfigPath()
	}

	// no config provided, can't proceed
	if confFilePath == "" && confContent == "" {
		log.Errorf("neither --instance-conf-path nor --instance-conf provided, cannot continue")
		return nil
	}

	if flagProvided("instance-conf-path") && confFilePath != "" && confContent != "" {
		// both types of configs explicitly provided, this is something we should at least warn on
		log.Warnf("both --instance-conf-path and --instance-conf provided, defaulting to --instance-conf-path: %s", confFilePath)
		// wipe the flag content so that we load the file instead
		confContent = ""
	}

	c := &Conf{}

	// if set, use the command line flag value
	if confContent != "" {
		err := json.Unmarshal([]byte(confContent), c)
		if err != nil {
			log.Errorf("unmarshal config content error:%s", err.Error())
			return nil
		}
		return c
	}

	// otherwise load from the file
	yamlFile, err := os.ReadFile(confFilePath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			log.Errorf("supplied config file does not exist")
			return nil
		}
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Errorf("unmarshal yaml file error:%s", err.Error())
		return nil
	}
	return c
}

func getDefaultConfigPath() string {
	for _, configPath := range []string{
		"conf/conf.yaml",
		"conf.yaml",
	} {
		s, err := os.Stat(configPath)
		if err != nil || s.IsDir() {
			continue
		}
		return configPath
	}
	return ""
}

func flagProvided(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
