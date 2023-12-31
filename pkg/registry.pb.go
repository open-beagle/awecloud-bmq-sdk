// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: pkg/registry.proto

package pkg

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	// 获取服务列表
	Action_GetServices Action = 0
	// 退出当前程序
	Action_Exits Action = 1
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "GetServices",
		1: "Exits",
	}
	Action_value = map[string]int32{
		"GetServices": 0,
		"Exits":       1,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_registry_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_pkg_registry_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{0}
}

type Kind int32

const (
	// 程序在K8S中运行
	Kind_K8S Kind = 0
	// 程序在容器中运行
	Kind_Containerd Kind = 1
	// 程序在命令行中运行
	Kind_Terminal Kind = 2
	// 程序是桌面客户端
	Kind_GUI Kind = 3
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "K8S",
		1: "Containerd",
		2: "Terminal",
		3: "GUI",
	}
	Kind_value = map[string]int32{
		"K8S":        0,
		"Containerd": 1,
		"Terminal":   2,
		"GUI":        3,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_registry_proto_enumTypes[1].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_pkg_registry_proto_enumTypes[1]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{1}
}

// LoginRequest 集群信息
type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID, 客户端集群ID，aliyun
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Secret, 客户端集群的密钥, Bp#q3sM6uxK4
	Secret string `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *LoginRequest) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

// LoginResponse 集群信息
type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path, 集群的URL的访问路径,
	// 其默认值， /awecloud/cluster/access
	Path string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	// Secret, 集群的密钥, Bp#q3sM6uxK4
	Secret string `protobuf:"bytes,2,opt,name=Secret,proto3" json:"Secret,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LoginResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

// ListenRequest 集群信息
type ListenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID, 客户端集群ID，aliyun
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Kind , 服务类型 ， 默认，K8S
	// 其他可选值，Containerd，Terminal，GUI
	Kind Kind `protobuf:"varint,2,opt,name=Kind,proto3,enum=Kind" json:"Kind,omitempty"`
	// OS , 操作系统 ， 如 Linux ， Windows
	OS string `protobuf:"bytes,3,opt,name=OS,proto3" json:"OS,omitempty"`
	// Arch , CPU架构 ， 如 amd64 ， arm64
	Arch string `protobuf:"bytes,4,opt,name=Arch,proto3" json:"Arch,omitempty"`
	// Kernel , 内核版本 ， 如 5.4
	Kernel string `protobuf:"bytes,5,opt,name=Kernel,proto3" json:"Kernel,omitempty"`
	// Labels , 其他标记
	Labels map[string]string `protobuf:"bytes,6,rep,name=Labels,proto3" json:"Labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ListenRequest) Reset() {
	*x = ListenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenRequest) ProtoMessage() {}

func (x *ListenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenRequest.ProtoReflect.Descriptor instead.
func (*ListenRequest) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{2}
}

func (x *ListenRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ListenRequest) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_K8S
}

func (x *ListenRequest) GetOS() string {
	if x != nil {
		return x.OS
	}
	return ""
}

func (x *ListenRequest) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

func (x *ListenRequest) GetKernel() string {
	if x != nil {
		return x.Kernel
	}
	return ""
}

func (x *ListenRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// ListenResponse ， 服务信息
type ListenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Action , 服务动作 ，
	Action Action `protobuf:"varint,1,opt,name=Action,proto3,enum=Action" json:"Action,omitempty"`
	// Data , 服务数据 ， 如Json
	Data string `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *ListenResponse) Reset() {
	*x = ListenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListenResponse) ProtoMessage() {}

func (x *ListenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListenResponse.ProtoReflect.Descriptor instead.
func (*ListenResponse) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{3}
}

func (x *ListenResponse) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_GetServices
}

func (x *ListenResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// GetServicesRequest 集群信息
type GetServicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID, 客户端集群ID，aliyun
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetServicesRequest) Reset() {
	*x = GetServicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesRequest) ProtoMessage() {}

func (x *GetServicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesRequest.ProtoReflect.Descriptor instead.
func (*GetServicesRequest) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{4}
}

func (x *GetServicesRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

// GetServicesResponse ， 服务信息
type GetServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AgentServices , 服务转发列表
	AgentServices []*ServiceConfig `protobuf:"bytes,1,rep,name=AgentServices,proto3" json:"AgentServices,omitempty"`
	// VisitorServices , 服务代理列表
	VisitorServices []*ServiceConfig `protobuf:"bytes,2,rep,name=VisitorServices,proto3" json:"VisitorServices,omitempty"`
}

func (x *GetServicesResponse) Reset() {
	*x = GetServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesResponse) ProtoMessage() {}

func (x *GetServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesResponse.ProtoReflect.Descriptor instead.
func (*GetServicesResponse) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{5}
}

func (x *GetServicesResponse) GetAgentServices() []*ServiceConfig {
	if x != nil {
		return x.AgentServices
	}
	return nil
}

func (x *GetServicesResponse) GetVisitorServices() []*ServiceConfig {
	if x != nil {
		return x.VisitorServices
	}
	return nil
}

// ServiceConfig ， 服务配置
type ServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 服务地址
	Server_Address string `protobuf:"bytes,1,opt,name=Server_Address,json=ServerAddress,proto3" json:"Server_Address,omitempty"`
	// 服务端口
	Server_Port string `protobuf:"bytes,2,opt,name=Server_Port,json=ServerPort,proto3" json:"Server_Port,omitempty"`
	// User , 租户 ，每个Secret服务都在租户下面进行管理与隔离
	User string `protobuf:"bytes,3,opt,name=User,proto3" json:"User,omitempty"`
	// Token , Server服务的密钥
	Token string `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	// Services , 服务集合
	Services map[string]*SecretService `protobuf:"bytes,5,rep,name=Services,proto3" json:"Services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ServiceConfig) Reset() {
	*x = ServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceConfig) ProtoMessage() {}

func (x *ServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceConfig.ProtoReflect.Descriptor instead.
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{6}
}

func (x *ServiceConfig) GetServer_Address() string {
	if x != nil {
		return x.Server_Address
	}
	return ""
}

func (x *ServiceConfig) GetServer_Port() string {
	if x != nil {
		return x.Server_Port
	}
	return ""
}

func (x *ServiceConfig) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ServiceConfig) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ServiceConfig) GetServices() map[string]*SecretService {
	if x != nil {
		return x.Services
	}
	return nil
}

// SecretService ， 服务
type SecretService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type , 服务类型 ， 默认，stcp
	Type string `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	// User , 服务名称 ， 如 k8s
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// Secret , 服务类型 ， 如 stcp
	Secret string `protobuf:"bytes,3,opt,name=Secret,proto3" json:"Secret,omitempty"`
	// Host , 服务主机 ， 如 kubernetes.default
	Host string `protobuf:"bytes,4,opt,name=Host,proto3" json:"Host,omitempty"`
	// Port , 服务端口 ， 如 6443
	Port int32 `protobuf:"varint,5,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *SecretService) Reset() {
	*x = SecretService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_registry_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecretService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecretService) ProtoMessage() {}

func (x *SecretService) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_registry_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecretService.ProtoReflect.Descriptor instead.
func (*SecretService) Descriptor() ([]byte, []int) {
	return file_pkg_registry_proto_rawDescGZIP(), []int{7}
}

func (x *SecretService) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SecretService) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecretService) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

func (x *SecretService) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *SecretService) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_pkg_registry_proto protoreflect.FileDescriptor

var file_pkg_registry_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x3b, 0x0a, 0x0d,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x4c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x04, 0x4b,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x4b, 0x69, 0x6e, 0x64,
	0x52, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x4f, 0x53, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x4f, 0x53, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x63, 0x68, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x72, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x4b, 0x65,
	0x72, 0x6e, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4b, 0x65, 0x72, 0x6e,
	0x65, 0x6c, 0x12, 0x32, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x45, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22, 0x85,
	0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0f,
	0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x56, 0x69, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x88, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x38, 0x0a, 0x08, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x1a, 0x4b, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x77, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x2a, 0x24, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x78, 0x69, 0x74, 0x73, 0x10, 0x01,
	0x2a, 0x36, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x4b, 0x38, 0x53, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x64, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x47, 0x55, 0x49, 0x10, 0x03, 0x32, 0x9f, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0d,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x2d, 0x0a, 0x06, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x13, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x62, 0x65,
	0x61, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x77, 0x65, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x62, 0x6d,
	0x71, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6b, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_registry_proto_rawDescOnce sync.Once
	file_pkg_registry_proto_rawDescData = file_pkg_registry_proto_rawDesc
)

func file_pkg_registry_proto_rawDescGZIP() []byte {
	file_pkg_registry_proto_rawDescOnce.Do(func() {
		file_pkg_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_registry_proto_rawDescData)
	})
	return file_pkg_registry_proto_rawDescData
}

var file_pkg_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_pkg_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pkg_registry_proto_goTypes = []interface{}{
	(Action)(0),                 // 0: Action
	(Kind)(0),                   // 1: Kind
	(*LoginRequest)(nil),        // 2: LoginRequest
	(*LoginResponse)(nil),       // 3: LoginResponse
	(*ListenRequest)(nil),       // 4: ListenRequest
	(*ListenResponse)(nil),      // 5: ListenResponse
	(*GetServicesRequest)(nil),  // 6: GetServicesRequest
	(*GetServicesResponse)(nil), // 7: GetServicesResponse
	(*ServiceConfig)(nil),       // 8: ServiceConfig
	(*SecretService)(nil),       // 9: SecretService
	nil,                         // 10: ListenRequest.LabelsEntry
	nil,                         // 11: ServiceConfig.ServicesEntry
}
var file_pkg_registry_proto_depIdxs = []int32{
	1,  // 0: ListenRequest.Kind:type_name -> Kind
	10, // 1: ListenRequest.Labels:type_name -> ListenRequest.LabelsEntry
	0,  // 2: ListenResponse.Action:type_name -> Action
	8,  // 3: GetServicesResponse.AgentServices:type_name -> ServiceConfig
	8,  // 4: GetServicesResponse.VisitorServices:type_name -> ServiceConfig
	11, // 5: ServiceConfig.Services:type_name -> ServiceConfig.ServicesEntry
	9,  // 6: ServiceConfig.ServicesEntry.value:type_name -> SecretService
	2,  // 7: Registry.Login:input_type -> LoginRequest
	4,  // 8: Registry.Listen:input_type -> ListenRequest
	6,  // 9: Registry.GetServices:input_type -> GetServicesRequest
	3,  // 10: Registry.Login:output_type -> LoginResponse
	5,  // 11: Registry.Listen:output_type -> ListenResponse
	7,  // 12: Registry.GetServices:output_type -> GetServicesResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_registry_proto_init() }
func file_pkg_registry_proto_init() {
	if File_pkg_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListenResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServiceConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_registry_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecretService); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_registry_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_registry_proto_goTypes,
		DependencyIndexes: file_pkg_registry_proto_depIdxs,
		EnumInfos:         file_pkg_registry_proto_enumTypes,
		MessageInfos:      file_pkg_registry_proto_msgTypes,
	}.Build()
	File_pkg_registry_proto = out.File
	file_pkg_registry_proto_rawDesc = nil
	file_pkg_registry_proto_goTypes = nil
	file_pkg_registry_proto_depIdxs = nil
}
