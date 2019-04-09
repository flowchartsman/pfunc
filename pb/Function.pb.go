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

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Function.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProcessingGuarantees int32

const (
	ProcessingGuarantees_ATLEAST_ONCE     ProcessingGuarantees = 0
	ProcessingGuarantees_ATMOST_ONCE      ProcessingGuarantees = 1
	ProcessingGuarantees_EFFECTIVELY_ONCE ProcessingGuarantees = 2
)

var ProcessingGuarantees_name = map[int32]string{
	0: "ATLEAST_ONCE",
	1: "ATMOST_ONCE",
	2: "EFFECTIVELY_ONCE",
}
var ProcessingGuarantees_value = map[string]int32{
	"ATLEAST_ONCE":     0,
	"ATMOST_ONCE":      1,
	"EFFECTIVELY_ONCE": 2,
}

func (x ProcessingGuarantees) String() string {
	return proto.EnumName(ProcessingGuarantees_name, int32(x))
}
func (ProcessingGuarantees) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{0}
}

type SubscriptionType int32

const (
	SubscriptionType_SHARED   SubscriptionType = 0
	SubscriptionType_FAILOVER SubscriptionType = 1
)

var SubscriptionType_name = map[int32]string{
	0: "SHARED",
	1: "FAILOVER",
}
var SubscriptionType_value = map[string]int32{
	"SHARED":   0,
	"FAILOVER": 1,
}

func (x SubscriptionType) String() string {
	return proto.EnumName(SubscriptionType_name, int32(x))
}
func (SubscriptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{1}
}

type FunctionState int32

const (
	FunctionState_RUNNING FunctionState = 0
	FunctionState_STOPPED FunctionState = 1
)

var FunctionState_name = map[int32]string{
	0: "RUNNING",
	1: "STOPPED",
}
var FunctionState_value = map[string]int32{
	"RUNNING": 0,
	"STOPPED": 1,
}

func (x FunctionState) String() string {
	return proto.EnumName(FunctionState_name, int32(x))
}
func (FunctionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{2}
}

type FunctionDetails_Runtime int32

const (
	FunctionDetails_JAVA   FunctionDetails_Runtime = 0
	FunctionDetails_PYTHON FunctionDetails_Runtime = 1
	FunctionDetails_GO     FunctionDetails_Runtime = 3
)

var FunctionDetails_Runtime_name = map[int32]string{
	0: "JAVA",
	1: "PYTHON",
	3: "GO",
}
var FunctionDetails_Runtime_value = map[string]int32{
	"JAVA":   0,
	"PYTHON": 1,
	"GO":     3,
}

func (x FunctionDetails_Runtime) String() string {
	return proto.EnumName(FunctionDetails_Runtime_name, int32(x))
}
func (FunctionDetails_Runtime) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{2, 0}
}

type Resources struct {
	Cpu                  float64  `protobuf:"fixed64,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Ram                  int64    `protobuf:"varint,2,opt,name=ram,proto3" json:"ram,omitempty"`
	Disk                 int64    `protobuf:"varint,3,opt,name=disk,proto3" json:"disk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resources) Reset()         { *m = Resources{} }
func (m *Resources) String() string { return proto.CompactTextString(m) }
func (*Resources) ProtoMessage()    {}
func (*Resources) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{0}
}
func (m *Resources) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resources.Unmarshal(m, b)
}
func (m *Resources) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resources.Marshal(b, m, deterministic)
}
func (dst *Resources) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resources.Merge(dst, src)
}
func (m *Resources) XXX_Size() int {
	return xxx_messageInfo_Resources.Size(m)
}
func (m *Resources) XXX_DiscardUnknown() {
	xxx_messageInfo_Resources.DiscardUnknown(m)
}

var xxx_messageInfo_Resources proto.InternalMessageInfo

func (m *Resources) GetCpu() float64 {
	if m != nil {
		return m.Cpu
	}
	return 0
}

func (m *Resources) GetRam() int64 {
	if m != nil {
		return m.Ram
	}
	return 0
}

func (m *Resources) GetDisk() int64 {
	if m != nil {
		return m.Disk
	}
	return 0
}

type RetryDetails struct {
	MaxMessageRetries    int32    `protobuf:"varint,1,opt,name=maxMessageRetries,proto3" json:"maxMessageRetries,omitempty"`
	DeadLetterTopic      string   `protobuf:"bytes,2,opt,name=deadLetterTopic,proto3" json:"deadLetterTopic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetryDetails) Reset()         { *m = RetryDetails{} }
func (m *RetryDetails) String() string { return proto.CompactTextString(m) }
func (*RetryDetails) ProtoMessage()    {}
func (*RetryDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{1}
}
func (m *RetryDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RetryDetails.Unmarshal(m, b)
}
func (m *RetryDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RetryDetails.Marshal(b, m, deterministic)
}
func (dst *RetryDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetryDetails.Merge(dst, src)
}
func (m *RetryDetails) XXX_Size() int {
	return xxx_messageInfo_RetryDetails.Size(m)
}
func (m *RetryDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_RetryDetails.DiscardUnknown(m)
}

var xxx_messageInfo_RetryDetails proto.InternalMessageInfo

func (m *RetryDetails) GetMaxMessageRetries() int32 {
	if m != nil {
		return m.MaxMessageRetries
	}
	return 0
}

func (m *RetryDetails) GetDeadLetterTopic() string {
	if m != nil {
		return m.DeadLetterTopic
	}
	return ""
}

type FunctionDetails struct {
	Tenant               string                  `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	Namespace            string                  `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ClassName            string                  `protobuf:"bytes,4,opt,name=className,proto3" json:"className,omitempty"`
	LogTopic             string                  `protobuf:"bytes,5,opt,name=logTopic,proto3" json:"logTopic,omitempty"`
	ProcessingGuarantees ProcessingGuarantees    `protobuf:"varint,6,opt,name=processingGuarantees,proto3,enum=proto.ProcessingGuarantees" json:"processingGuarantees,omitempty"`
	UserConfig           string                  `protobuf:"bytes,7,opt,name=userConfig,proto3" json:"userConfig,omitempty"`
	SecretsMap           string                  `protobuf:"bytes,16,opt,name=secretsMap,proto3" json:"secretsMap,omitempty"`
	Runtime              FunctionDetails_Runtime `protobuf:"varint,8,opt,name=runtime,proto3,enum=proto.FunctionDetails_Runtime" json:"runtime,omitempty"`
	AutoAck              bool                    `protobuf:"varint,9,opt,name=autoAck,proto3" json:"autoAck,omitempty"`
	Parallelism          int32                   `protobuf:"varint,10,opt,name=parallelism,proto3" json:"parallelism,omitempty"`
	Source               *SourceSpec             `protobuf:"bytes,11,opt,name=source,proto3" json:"source,omitempty"`
	Sink                 *SinkSpec               `protobuf:"bytes,12,opt,name=sink,proto3" json:"sink,omitempty"`
	Resources            *Resources              `protobuf:"bytes,13,opt,name=resources,proto3" json:"resources,omitempty"`
	PackageUrl           string                  `protobuf:"bytes,14,opt,name=packageUrl,proto3" json:"packageUrl,omitempty"`
	RetryDetails         *RetryDetails           `protobuf:"bytes,15,opt,name=retryDetails,proto3" json:"retryDetails,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *FunctionDetails) Reset()         { *m = FunctionDetails{} }
func (m *FunctionDetails) String() string { return proto.CompactTextString(m) }
func (*FunctionDetails) ProtoMessage()    {}
func (*FunctionDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{2}
}
func (m *FunctionDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionDetails.Unmarshal(m, b)
}
func (m *FunctionDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionDetails.Marshal(b, m, deterministic)
}
func (dst *FunctionDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionDetails.Merge(dst, src)
}
func (m *FunctionDetails) XXX_Size() int {
	return xxx_messageInfo_FunctionDetails.Size(m)
}
func (m *FunctionDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionDetails.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionDetails proto.InternalMessageInfo

func (m *FunctionDetails) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *FunctionDetails) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *FunctionDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FunctionDetails) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *FunctionDetails) GetLogTopic() string {
	if m != nil {
		return m.LogTopic
	}
	return ""
}

func (m *FunctionDetails) GetProcessingGuarantees() ProcessingGuarantees {
	if m != nil {
		return m.ProcessingGuarantees
	}
	return ProcessingGuarantees_ATLEAST_ONCE
}

func (m *FunctionDetails) GetUserConfig() string {
	if m != nil {
		return m.UserConfig
	}
	return ""
}

func (m *FunctionDetails) GetSecretsMap() string {
	if m != nil {
		return m.SecretsMap
	}
	return ""
}

func (m *FunctionDetails) GetRuntime() FunctionDetails_Runtime {
	if m != nil {
		return m.Runtime
	}
	return FunctionDetails_JAVA
}

func (m *FunctionDetails) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *FunctionDetails) GetParallelism() int32 {
	if m != nil {
		return m.Parallelism
	}
	return 0
}

func (m *FunctionDetails) GetSource() *SourceSpec {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *FunctionDetails) GetSink() *SinkSpec {
	if m != nil {
		return m.Sink
	}
	return nil
}

func (m *FunctionDetails) GetResources() *Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *FunctionDetails) GetPackageUrl() string {
	if m != nil {
		return m.PackageUrl
	}
	return ""
}

func (m *FunctionDetails) GetRetryDetails() *RetryDetails {
	if m != nil {
		return m.RetryDetails
	}
	return nil
}

type ConsumerSpec struct {
	SchemaType           string                          `protobuf:"bytes,1,opt,name=schemaType,proto3" json:"schemaType,omitempty"`
	SerdeClassName       string                          `protobuf:"bytes,2,opt,name=serdeClassName,proto3" json:"serdeClassName,omitempty"`
	IsRegexPattern       bool                            `protobuf:"varint,3,opt,name=isRegexPattern,proto3" json:"isRegexPattern,omitempty"`
	ReceiverQueueSize    *ConsumerSpec_ReceiverQueueSize `protobuf:"bytes,4,opt,name=receiverQueueSize,proto3" json:"receiverQueueSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ConsumerSpec) Reset()         { *m = ConsumerSpec{} }
func (m *ConsumerSpec) String() string { return proto.CompactTextString(m) }
func (*ConsumerSpec) ProtoMessage()    {}
func (*ConsumerSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{3}
}
func (m *ConsumerSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerSpec.Unmarshal(m, b)
}
func (m *ConsumerSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerSpec.Marshal(b, m, deterministic)
}
func (dst *ConsumerSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerSpec.Merge(dst, src)
}
func (m *ConsumerSpec) XXX_Size() int {
	return xxx_messageInfo_ConsumerSpec.Size(m)
}
func (m *ConsumerSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerSpec proto.InternalMessageInfo

func (m *ConsumerSpec) GetSchemaType() string {
	if m != nil {
		return m.SchemaType
	}
	return ""
}

func (m *ConsumerSpec) GetSerdeClassName() string {
	if m != nil {
		return m.SerdeClassName
	}
	return ""
}

func (m *ConsumerSpec) GetIsRegexPattern() bool {
	if m != nil {
		return m.IsRegexPattern
	}
	return false
}

func (m *ConsumerSpec) GetReceiverQueueSize() *ConsumerSpec_ReceiverQueueSize {
	if m != nil {
		return m.ReceiverQueueSize
	}
	return nil
}

type ConsumerSpec_ReceiverQueueSize struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumerSpec_ReceiverQueueSize) Reset()         { *m = ConsumerSpec_ReceiverQueueSize{} }
func (m *ConsumerSpec_ReceiverQueueSize) String() string { return proto.CompactTextString(m) }
func (*ConsumerSpec_ReceiverQueueSize) ProtoMessage()    {}
func (*ConsumerSpec_ReceiverQueueSize) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{3, 0}
}
func (m *ConsumerSpec_ReceiverQueueSize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerSpec_ReceiverQueueSize.Unmarshal(m, b)
}
func (m *ConsumerSpec_ReceiverQueueSize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerSpec_ReceiverQueueSize.Marshal(b, m, deterministic)
}
func (dst *ConsumerSpec_ReceiverQueueSize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerSpec_ReceiverQueueSize.Merge(dst, src)
}
func (m *ConsumerSpec_ReceiverQueueSize) XXX_Size() int {
	return xxx_messageInfo_ConsumerSpec_ReceiverQueueSize.Size(m)
}
func (m *ConsumerSpec_ReceiverQueueSize) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerSpec_ReceiverQueueSize.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerSpec_ReceiverQueueSize proto.InternalMessageInfo

func (m *ConsumerSpec_ReceiverQueueSize) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SourceSpec struct {
	ClassName string `protobuf:"bytes,1,opt,name=className,proto3" json:"className,omitempty"`
	// map in json format
	Configs       string `protobuf:"bytes,2,opt,name=configs,proto3" json:"configs,omitempty"`
	TypeClassName string `protobuf:"bytes,5,opt,name=typeClassName,proto3" json:"typeClassName,omitempty"`
	// configs used only when source feeds into functions
	SubscriptionType SubscriptionType `protobuf:"varint,3,opt,name=subscriptionType,proto3,enum=proto.SubscriptionType" json:"subscriptionType,omitempty"`
	// @deprecated -- use topicsToSchema
	TopicsToSerDeClassName map[string]string `protobuf:"bytes,4,rep,name=topicsToSerDeClassName,proto3" json:"topicsToSerDeClassName,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	// *
	//
	InputSpecs    map[string]*ConsumerSpec `protobuf:"bytes,10,rep,name=inputSpecs,proto3" json:"inputSpecs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TimeoutMs     uint64                   `protobuf:"varint,6,opt,name=timeoutMs,proto3" json:"timeoutMs,omitempty"`
	TopicsPattern string                   `protobuf:"bytes,7,opt,name=topicsPattern,proto3" json:"topicsPattern,omitempty"` // Deprecated: Do not use.
	// If specified, this will refer to an archive that is
	// already present in the server
	Builtin              string   `protobuf:"bytes,8,opt,name=builtin,proto3" json:"builtin,omitempty"`
	SubscriptionName     string   `protobuf:"bytes,9,opt,name=subscriptionName,proto3" json:"subscriptionName,omitempty"`
	CleanupSubscription  bool     `protobuf:"varint,11,opt,name=cleanupSubscription,proto3" json:"cleanupSubscription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceSpec) Reset()         { *m = SourceSpec{} }
func (m *SourceSpec) String() string { return proto.CompactTextString(m) }
func (*SourceSpec) ProtoMessage()    {}
func (*SourceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{4}
}
func (m *SourceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceSpec.Unmarshal(m, b)
}
func (m *SourceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceSpec.Marshal(b, m, deterministic)
}
func (dst *SourceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceSpec.Merge(dst, src)
}
func (m *SourceSpec) XXX_Size() int {
	return xxx_messageInfo_SourceSpec.Size(m)
}
func (m *SourceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SourceSpec proto.InternalMessageInfo

func (m *SourceSpec) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *SourceSpec) GetConfigs() string {
	if m != nil {
		return m.Configs
	}
	return ""
}

func (m *SourceSpec) GetTypeClassName() string {
	if m != nil {
		return m.TypeClassName
	}
	return ""
}

func (m *SourceSpec) GetSubscriptionType() SubscriptionType {
	if m != nil {
		return m.SubscriptionType
	}
	return SubscriptionType_SHARED
}

// Deprecated: Do not use.
func (m *SourceSpec) GetTopicsToSerDeClassName() map[string]string {
	if m != nil {
		return m.TopicsToSerDeClassName
	}
	return nil
}

func (m *SourceSpec) GetInputSpecs() map[string]*ConsumerSpec {
	if m != nil {
		return m.InputSpecs
	}
	return nil
}

func (m *SourceSpec) GetTimeoutMs() uint64 {
	if m != nil {
		return m.TimeoutMs
	}
	return 0
}

// Deprecated: Do not use.
func (m *SourceSpec) GetTopicsPattern() string {
	if m != nil {
		return m.TopicsPattern
	}
	return ""
}

func (m *SourceSpec) GetBuiltin() string {
	if m != nil {
		return m.Builtin
	}
	return ""
}

func (m *SourceSpec) GetSubscriptionName() string {
	if m != nil {
		return m.SubscriptionName
	}
	return ""
}

func (m *SourceSpec) GetCleanupSubscription() bool {
	if m != nil {
		return m.CleanupSubscription
	}
	return false
}

type SinkSpec struct {
	ClassName string `protobuf:"bytes,1,opt,name=className,proto3" json:"className,omitempty"`
	// map in json format
	Configs       string `protobuf:"bytes,2,opt,name=configs,proto3" json:"configs,omitempty"`
	TypeClassName string `protobuf:"bytes,5,opt,name=typeClassName,proto3" json:"typeClassName,omitempty"`
	// configs used only when functions output to sink
	Topic          string `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	SerDeClassName string `protobuf:"bytes,4,opt,name=serDeClassName,proto3" json:"serDeClassName,omitempty"`
	// If specified, this will refer to an archive that is
	// already present in the server
	Builtin string `protobuf:"bytes,6,opt,name=builtin,proto3" json:"builtin,omitempty"`
	// *
	// Builtin schema type or custom schema class name
	SchemaType           string   `protobuf:"bytes,7,opt,name=schemaType,proto3" json:"schemaType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SinkSpec) Reset()         { *m = SinkSpec{} }
func (m *SinkSpec) String() string { return proto.CompactTextString(m) }
func (*SinkSpec) ProtoMessage()    {}
func (*SinkSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{5}
}
func (m *SinkSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SinkSpec.Unmarshal(m, b)
}
func (m *SinkSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SinkSpec.Marshal(b, m, deterministic)
}
func (dst *SinkSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SinkSpec.Merge(dst, src)
}
func (m *SinkSpec) XXX_Size() int {
	return xxx_messageInfo_SinkSpec.Size(m)
}
func (m *SinkSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SinkSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SinkSpec proto.InternalMessageInfo

func (m *SinkSpec) GetClassName() string {
	if m != nil {
		return m.ClassName
	}
	return ""
}

func (m *SinkSpec) GetConfigs() string {
	if m != nil {
		return m.Configs
	}
	return ""
}

func (m *SinkSpec) GetTypeClassName() string {
	if m != nil {
		return m.TypeClassName
	}
	return ""
}

func (m *SinkSpec) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SinkSpec) GetSerDeClassName() string {
	if m != nil {
		return m.SerDeClassName
	}
	return ""
}

func (m *SinkSpec) GetBuiltin() string {
	if m != nil {
		return m.Builtin
	}
	return ""
}

func (m *SinkSpec) GetSchemaType() string {
	if m != nil {
		return m.SchemaType
	}
	return ""
}

type PackageLocationMetaData struct {
	PackagePath          string   `protobuf:"bytes,1,opt,name=packagePath,proto3" json:"packagePath,omitempty"`
	OriginalFileName     string   `protobuf:"bytes,2,opt,name=originalFileName,proto3" json:"originalFileName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PackageLocationMetaData) Reset()         { *m = PackageLocationMetaData{} }
func (m *PackageLocationMetaData) String() string { return proto.CompactTextString(m) }
func (*PackageLocationMetaData) ProtoMessage()    {}
func (*PackageLocationMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{6}
}
func (m *PackageLocationMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageLocationMetaData.Unmarshal(m, b)
}
func (m *PackageLocationMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageLocationMetaData.Marshal(b, m, deterministic)
}
func (dst *PackageLocationMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageLocationMetaData.Merge(dst, src)
}
func (m *PackageLocationMetaData) XXX_Size() int {
	return xxx_messageInfo_PackageLocationMetaData.Size(m)
}
func (m *PackageLocationMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageLocationMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_PackageLocationMetaData proto.InternalMessageInfo

func (m *PackageLocationMetaData) GetPackagePath() string {
	if m != nil {
		return m.PackagePath
	}
	return ""
}

func (m *PackageLocationMetaData) GetOriginalFileName() string {
	if m != nil {
		return m.OriginalFileName
	}
	return ""
}

type FunctionMetaData struct {
	FunctionDetails      *FunctionDetails         `protobuf:"bytes,1,opt,name=functionDetails,proto3" json:"functionDetails,omitempty"`
	PackageLocation      *PackageLocationMetaData `protobuf:"bytes,2,opt,name=packageLocation,proto3" json:"packageLocation,omitempty"`
	Version              uint64                   `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	CreateTime           uint64                   `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	InstanceStates       map[int32]FunctionState  `protobuf:"bytes,5,rep,name=instanceStates,proto3" json:"instanceStates,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=proto.FunctionState"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *FunctionMetaData) Reset()         { *m = FunctionMetaData{} }
func (m *FunctionMetaData) String() string { return proto.CompactTextString(m) }
func (*FunctionMetaData) ProtoMessage()    {}
func (*FunctionMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{7}
}
func (m *FunctionMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionMetaData.Unmarshal(m, b)
}
func (m *FunctionMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionMetaData.Marshal(b, m, deterministic)
}
func (dst *FunctionMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionMetaData.Merge(dst, src)
}
func (m *FunctionMetaData) XXX_Size() int {
	return xxx_messageInfo_FunctionMetaData.Size(m)
}
func (m *FunctionMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionMetaData proto.InternalMessageInfo

func (m *FunctionMetaData) GetFunctionDetails() *FunctionDetails {
	if m != nil {
		return m.FunctionDetails
	}
	return nil
}

func (m *FunctionMetaData) GetPackageLocation() *PackageLocationMetaData {
	if m != nil {
		return m.PackageLocation
	}
	return nil
}

func (m *FunctionMetaData) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *FunctionMetaData) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *FunctionMetaData) GetInstanceStates() map[int32]FunctionState {
	if m != nil {
		return m.InstanceStates
	}
	return nil
}

type Instance struct {
	FunctionMetaData     *FunctionMetaData `protobuf:"bytes,1,opt,name=functionMetaData,proto3" json:"functionMetaData,omitempty"`
	InstanceId           int32             `protobuf:"varint,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{8}
}
func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (dst *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(dst, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetFunctionMetaData() *FunctionMetaData {
	if m != nil {
		return m.FunctionMetaData
	}
	return nil
}

func (m *Instance) GetInstanceId() int32 {
	if m != nil {
		return m.InstanceId
	}
	return 0
}

type Assignment struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	WorkerId             string    `protobuf:"bytes,2,opt,name=workerId,proto3" json:"workerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_Function_33c6e1841d2624f0, []int{9}
}
func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (dst *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(dst, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *Assignment) GetWorkerId() string {
	if m != nil {
		return m.WorkerId
	}
	return ""
}

func init() {
	proto.RegisterType((*Resources)(nil), "proto.Resources")
	proto.RegisterType((*RetryDetails)(nil), "proto.RetryDetails")
	proto.RegisterType((*FunctionDetails)(nil), "proto.FunctionDetails")
	proto.RegisterType((*ConsumerSpec)(nil), "proto.ConsumerSpec")
	proto.RegisterType((*ConsumerSpec_ReceiverQueueSize)(nil), "proto.ConsumerSpec.ReceiverQueueSize")
	proto.RegisterType((*SourceSpec)(nil), "proto.SourceSpec")
	proto.RegisterMapType((map[string]*ConsumerSpec)(nil), "proto.SourceSpec.InputSpecsEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto.SourceSpec.TopicsToSerDeClassNameEntry")
	proto.RegisterType((*SinkSpec)(nil), "proto.SinkSpec")
	proto.RegisterType((*PackageLocationMetaData)(nil), "proto.PackageLocationMetaData")
	proto.RegisterType((*FunctionMetaData)(nil), "proto.FunctionMetaData")
	proto.RegisterMapType((map[int32]FunctionState)(nil), "proto.FunctionMetaData.InstanceStatesEntry")
	proto.RegisterType((*Instance)(nil), "proto.Instance")
	proto.RegisterType((*Assignment)(nil), "proto.Assignment")
	proto.RegisterEnum("proto.ProcessingGuarantees", ProcessingGuarantees_name, ProcessingGuarantees_value)
	proto.RegisterEnum("proto.SubscriptionType", SubscriptionType_name, SubscriptionType_value)
	proto.RegisterEnum("proto.FunctionState", FunctionState_name, FunctionState_value)
	proto.RegisterEnum("proto.FunctionDetails_Runtime", FunctionDetails_Runtime_name, FunctionDetails_Runtime_value)
}

func init() { proto.RegisterFile("Function.proto", fileDescriptor_Function_33c6e1841d2624f0) }

var fileDescriptor_Function_33c6e1841d2624f0 = []byte{
	// 1220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x36, 0xad, 0xf3, 0xc8, 0xb6, 0xe8, 0x8d, 0x91, 0x10, 0xce, 0x8f, 0x40, 0xd1, 0xdf, 0x83,
	0xec, 0x24, 0x42, 0xe1, 0x5e, 0x34, 0xe8, 0x55, 0x15, 0x59, 0x4e, 0x54, 0xd8, 0x96, 0xba, 0x52,
	0x52, 0xe4, 0xaa, 0xd8, 0xd0, 0x63, 0x65, 0x21, 0x8a, 0x24, 0x76, 0x97, 0x69, 0xdc, 0x07, 0xe8,
	0x63, 0xf4, 0x49, 0xfa, 0x30, 0x7d, 0x92, 0xa2, 0xd8, 0x25, 0x29, 0x91, 0x94, 0xdc, 0xbb, 0x5e,
	0x89, 0xf3, 0xcd, 0x71, 0x67, 0xbf, 0x99, 0x15, 0x1c, 0x5c, 0x44, 0xbe, 0xab, 0x78, 0xe0, 0xf7,
	0x42, 0x11, 0xa8, 0x80, 0x54, 0xcc, 0x4f, 0x67, 0x00, 0x0d, 0x8a, 0x32, 0x88, 0x84, 0x8b, 0x92,
	0xd8, 0x50, 0x72, 0xc3, 0xc8, 0xb1, 0xda, 0x56, 0xd7, 0xa2, 0xfa, 0x53, 0x23, 0x82, 0x2d, 0x9d,
	0xdd, 0xb6, 0xd5, 0x2d, 0x51, 0xfd, 0x49, 0x08, 0x94, 0x6f, 0xb8, 0x5c, 0x38, 0x25, 0x03, 0x99,
	0xef, 0xce, 0x2d, 0xec, 0x51, 0x54, 0xe2, 0xee, 0x1c, 0x15, 0xe3, 0x9e, 0x24, 0xcf, 0xe1, 0x70,
	0xc9, 0x3e, 0x5f, 0xa1, 0x94, 0x6c, 0x8e, 0x5a, 0xc3, 0x51, 0x9a, 0xa8, 0x15, 0xba, 0xa9, 0x20,
	0x5d, 0x68, 0xdd, 0x20, 0xbb, 0xb9, 0x44, 0xa5, 0x50, 0xcc, 0x82, 0x90, 0xbb, 0x26, 0x5f, 0x83,
	0x16, 0xe1, 0xce, 0x1f, 0x15, 0x68, 0xa5, 0xc7, 0x48, 0x73, 0x3d, 0x84, 0xaa, 0x42, 0x9f, 0xf9,
	0xca, 0x24, 0x68, 0xd0, 0x44, 0x22, 0xff, 0x83, 0x86, 0xcf, 0x96, 0x28, 0x43, 0xe6, 0x62, 0x12,
	0x6f, 0x0d, 0xe8, 0x53, 0x68, 0xc1, 0x9c, 0xa2, 0x41, 0xcd, 0xb7, 0xf6, 0x70, 0x3d, 0x26, 0xe5,
	0xb5, 0x56, 0x94, 0x63, 0x8f, 0x15, 0x40, 0x8e, 0xa1, 0xee, 0x05, 0xf3, 0xb8, 0xbc, 0x8a, 0x51,
	0xae, 0x64, 0x32, 0x86, 0xa3, 0x50, 0x04, 0x2e, 0x4a, 0xc9, 0xfd, 0xf9, 0xeb, 0x88, 0x09, 0xe6,
	0x2b, 0x44, 0xe9, 0x54, 0xdb, 0x56, 0xf7, 0xe0, 0xec, 0x71, 0xdc, 0xf1, 0xde, 0x64, 0x8b, 0x09,
	0xdd, 0xea, 0x48, 0x9e, 0x00, 0x44, 0x12, 0xc5, 0x20, 0xf0, 0x6f, 0xf9, 0xdc, 0xa9, 0x99, 0x74,
	0x19, 0x44, 0xeb, 0x25, 0xba, 0x02, 0x95, 0xbc, 0x62, 0xa1, 0x63, 0xc7, 0xfa, 0x35, 0x42, 0x5e,
	0x42, 0x4d, 0x44, 0xbe, 0xe2, 0x4b, 0x74, 0xea, 0xa6, 0x86, 0x27, 0x49, 0x0d, 0x85, 0xee, 0xf5,
	0x68, 0x6c, 0x45, 0x53, 0x73, 0xe2, 0x40, 0x8d, 0x45, 0x2a, 0xe8, 0xbb, 0x0b, 0xa7, 0xd1, 0xb6,
	0xba, 0x75, 0x9a, 0x8a, 0xa4, 0x0d, 0xcd, 0x90, 0x09, 0xe6, 0x79, 0xe8, 0x71, 0xb9, 0x74, 0xc0,
	0x5c, 0x67, 0x16, 0x22, 0x27, 0x50, 0x8d, 0x99, 0xe4, 0x34, 0xdb, 0x56, 0xb7, 0x79, 0x76, 0x98,
	0x24, 0x9d, 0x1a, 0x70, 0x1a, 0xa2, 0x4b, 0x13, 0x03, 0xf2, 0x7f, 0x28, 0x4b, 0xee, 0x2f, 0x9c,
	0x3d, 0x63, 0xd8, 0x4a, 0x0d, 0xb9, 0xbf, 0x30, 0x66, 0x46, 0x49, 0x7a, 0xd0, 0x10, 0x29, 0x37,
	0x9d, 0x7d, 0x63, 0x69, 0x27, 0x96, 0x2b, 0xce, 0xd2, 0xb5, 0x89, 0xee, 0x4a, 0xc8, 0xdc, 0x05,
	0x9b, 0xe3, 0x5b, 0xe1, 0x39, 0x07, 0x71, 0x57, 0xd6, 0x08, 0xf9, 0x0e, 0xf6, 0x44, 0x86, 0xa6,
	0x4e, 0xcb, 0x84, 0x7c, 0xb0, 0x0a, 0xb9, 0x56, 0xd1, 0x9c, 0x61, 0xe7, 0x6b, 0xa8, 0x25, 0x8d,
	0x22, 0x75, 0x28, 0xff, 0xd8, 0x7f, 0xd7, 0xb7, 0x77, 0x08, 0x40, 0x75, 0xf2, 0x7e, 0xf6, 0x66,
	0x7c, 0x6d, 0x5b, 0xa4, 0x0a, 0xbb, 0xaf, 0xc7, 0x76, 0xa9, 0xf3, 0xb7, 0x05, 0x7b, 0x83, 0xc0,
	0x97, 0xd1, 0x12, 0x85, 0x3e, 0x88, 0xb9, 0x28, 0xf7, 0x23, 0x2e, 0xd9, 0xec, 0x2e, 0xc4, 0x84,
	0xa1, 0x19, 0x84, 0x7c, 0x05, 0x07, 0x12, 0xc5, 0x0d, 0x0e, 0x56, 0xc4, 0x8b, 0xa9, 0x5a, 0x40,
	0xb5, 0x1d, 0x97, 0x14, 0xe7, 0xf8, 0x79, 0xc2, 0xf4, 0x3c, 0xf8, 0x86, 0xb9, 0x75, 0x5a, 0x40,
	0xc9, 0x14, 0x0e, 0x05, 0xba, 0xc8, 0x3f, 0xa1, 0xf8, 0x29, 0xc2, 0x08, 0xa7, 0xfc, 0xb7, 0x98,
	0xcb, 0xcd, 0xb3, 0x2f, 0x93, 0x73, 0x66, 0xeb, 0xeb, 0xd1, 0xa2, 0x31, 0xdd, 0xf4, 0x3f, 0x3e,
	0x81, 0xc3, 0x0d, 0x3b, 0x72, 0x04, 0x95, 0x4f, 0xcc, 0x8b, 0x30, 0x99, 0xeb, 0x58, 0xe8, 0xfc,
	0x59, 0x01, 0x58, 0x5f, 0x77, 0x7e, 0xa4, 0xac, 0xe2, 0x48, 0x39, 0x50, 0x73, 0x0d, 0x9f, 0x65,
	0x72, 0xea, 0x54, 0x24, 0x5f, 0xc0, 0xbe, 0xba, 0x0b, 0x33, 0x5d, 0x89, 0x27, 0x2e, 0x0f, 0x92,
	0x01, 0xd8, 0x32, 0xfa, 0x20, 0x5d, 0xc1, 0x43, 0xcd, 0x69, 0xd3, 0xe2, 0x92, 0xa1, 0xfb, 0xa3,
	0x94, 0x50, 0x05, 0x35, 0xdd, 0x70, 0x20, 0x1c, 0x1e, 0x2a, 0x3d, 0xc4, 0x72, 0x16, 0x4c, 0x51,
	0x9c, 0x67, 0x72, 0x96, 0xdb, 0xa5, 0x6e, 0xf3, 0xec, 0xc5, 0x06, 0x89, 0x7b, 0xb3, 0xad, 0xf6,
	0x43, 0x5f, 0x89, 0xbb, 0x57, 0xbb, 0x8e, 0x45, 0xef, 0x09, 0x48, 0xfa, 0x00, 0xdc, 0x0f, 0x23,
	0xa5, 0x83, 0x48, 0x07, 0x4c, 0xf8, 0xa7, 0x9b, 0xe1, 0x47, 0x2b, 0x1b, 0x13, 0x92, 0x66, 0x9c,
	0x74, 0x43, 0x35, 0x0d, 0x83, 0x48, 0x5d, 0xc5, 0xeb, 0xa5, 0x4c, 0xd7, 0x00, 0xe9, 0xc2, 0x7e,
	0x9c, 0x3a, 0x25, 0x89, 0xd9, 0x1c, 0xa6, 0xa6, 0xbc, 0x42, 0xb7, 0xfe, 0x43, 0xc4, 0x3d, 0xc5,
	0x7d, 0xb3, 0x20, 0x1a, 0x34, 0x15, 0xc9, 0x69, 0xbe, 0xa9, 0xa6, 0x13, 0x0d, 0x63, 0xb2, 0x81,
	0x93, 0x6f, 0xe0, 0x81, 0xeb, 0x21, 0xf3, 0xa3, 0x30, 0xdb, 0x68, 0x33, 0xfd, 0x75, 0xba, 0x4d,
	0x75, 0x3c, 0x82, 0xc7, 0xff, 0xd2, 0x3d, 0xfd, 0xdc, 0x2c, 0xf0, 0x2e, 0x61, 0x8a, 0xfe, 0x5c,
	0xd3, 0x2c, 0x66, 0x48, 0x2c, 0x7c, 0xbf, 0xfb, 0xd2, 0x3a, 0xa6, 0xd0, 0x2a, 0x74, 0x6a, 0x8b,
	0xfb, 0x49, 0xd6, 0x7d, 0x3d, 0xeb, 0xd9, 0x19, 0xc8, 0xc4, 0xec, 0xfc, 0x65, 0x41, 0x3d, 0x5d,
	0x42, 0xff, 0x31, 0x79, 0x8f, 0xa0, 0x62, 0xae, 0x24, 0x79, 0x82, 0x62, 0x21, 0xd9, 0x07, 0x79,
	0x16, 0xa6, 0xfb, 0x20, 0x4b, 0xa5, 0xcc, 0xfd, 0x55, 0xf3, 0xf7, 0x97, 0xdf, 0x38, 0xb5, 0xe2,
	0xc6, 0xe9, 0xcc, 0xe1, 0xd1, 0x24, 0x5e, 0x89, 0x97, 0x81, 0xcb, 0xf4, 0xa5, 0x5c, 0xa1, 0x62,
	0xe7, 0x4c, 0xb1, 0x78, 0xc3, 0x1b, 0xd5, 0x84, 0xa9, 0x8f, 0xc9, 0x91, 0xb3, 0x90, 0x26, 0x47,
	0x20, 0xf8, 0x9c, 0xfb, 0xcc, 0xbb, 0xe0, 0x1e, 0x66, 0x16, 0xd6, 0x06, 0xde, 0xf9, 0xbd, 0x04,
	0x76, 0xfa, 0xdc, 0xac, 0x52, 0xfc, 0x00, 0xad, 0xdb, 0xfc, 0x13, 0x64, 0xd2, 0x34, 0xcf, 0x1e,
	0x6e, 0x7f, 0xa0, 0x68, 0xd1, 0x9c, 0xbc, 0x81, 0x56, 0x98, 0xaf, 0x3f, 0xb9, 0xdb, 0xf4, 0x89,
	0xbb, 0xe7, 0x74, 0xb4, 0xe8, 0xa6, 0x7b, 0xf8, 0x09, 0x85, 0xd4, 0x11, 0x4a, 0x66, 0x92, 0x52,
	0x51, 0xf7, 0xd0, 0x15, 0xc8, 0x14, 0xce, 0x78, 0x72, 0x03, 0x65, 0x9a, 0x41, 0xc8, 0x14, 0x0e,
	0xb8, 0x2f, 0x15, 0xf3, 0x5d, 0x9c, 0x2a, 0xa6, 0x50, 0x3a, 0x15, 0x33, 0xcc, 0xcf, 0x0a, 0x87,
	0x48, 0x73, 0xf7, 0x46, 0x39, 0xeb, 0x78, 0xac, 0x0b, 0x21, 0x8e, 0x7f, 0x86, 0x07, 0x5b, 0xcc,
	0xb2, 0x9c, 0xae, 0xc4, 0x9c, 0x3e, 0xcd, 0x72, 0xfa, 0xe0, 0xec, 0xa8, 0x90, 0xd4, 0x38, 0x67,
	0x49, 0x1d, 0x40, 0x3d, 0x0d, 0xac, 0x57, 0xe6, 0x6d, 0xa1, 0xb8, 0xe4, 0x02, 0x1e, 0xdd, 0x53,
	0x3b, 0xdd, 0x70, 0xd0, 0xed, 0x49, 0x6b, 0x1f, 0xdd, 0x98, 0x2a, 0x2a, 0x34, 0x83, 0x74, 0xde,
	0x02, 0xf4, 0xa5, 0xe4, 0x73, 0x7f, 0x89, 0xbe, 0x22, 0xcf, 0xa0, 0x9e, 0xea, 0x92, 0x54, 0xe9,
	0x73, 0x9f, 0x56, 0x45, 0x57, 0x06, 0xfa, 0x5f, 0xd6, 0xaf, 0x81, 0x58, 0xa0, 0x48, 0x02, 0x37,
	0xe8, 0x4a, 0x3e, 0x1d, 0xc3, 0xd1, 0xb6, 0xbf, 0x50, 0xc4, 0x86, 0xbd, 0xfe, 0xec, 0x72, 0xd8,
	0x9f, 0xce, 0x7e, 0x19, 0x5f, 0x0f, 0x86, 0xf6, 0x0e, 0x69, 0x41, 0xb3, 0x3f, 0xbb, 0x1a, 0xa7,
	0x80, 0x45, 0x8e, 0xc0, 0x1e, 0x5e, 0x5c, 0x0c, 0x07, 0xb3, 0xd1, 0xbb, 0xe1, 0xe5, 0xfb, 0x18,
	0xdd, 0x3d, 0x7d, 0x0e, 0x76, 0xf1, 0x81, 0xd0, 0xaf, 0xfa, 0xf4, 0x4d, 0x9f, 0x0e, 0xcf, 0xed,
	0x1d, 0xb2, 0x07, 0xf5, 0x8b, 0xfe, 0xe8, 0x72, 0xfc, 0x6e, 0x48, 0x6d, 0xeb, 0xf4, 0x04, 0xf6,
	0x73, 0x2d, 0x26, 0x4d, 0xa8, 0xd1, 0xb7, 0xd7, 0xd7, 0xa3, 0xeb, 0xd7, 0xf6, 0x8e, 0x16, 0xa6,
	0xb3, 0xf1, 0x64, 0x32, 0x3c, 0xb7, 0xad, 0x57, 0x2f, 0xe0, 0x69, 0x20, 0xe6, 0x3d, 0x16, 0x32,
	0xf7, 0x23, 0xf6, 0xc2, 0xc8, 0x93, 0x4c, 0xf4, 0xd2, 0x36, 0xca, 0xf8, 0xf4, 0xaf, 0xea, 0x69,
	0xb4, 0x0f, 0x55, 0x03, 0x7c, 0xfb, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfe, 0xa8, 0xa5, 0x5f,
	0xa3, 0x0b, 0x00, 0x00,
}
