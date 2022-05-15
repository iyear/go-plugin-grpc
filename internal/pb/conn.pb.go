// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: conn.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UnbindReason int32

const (
	UnbindReason_Exit      UnbindReason = 0 // 插件退出
	UnbindReason_UnUsed    UnbindReason = 1 // 插件不再使用
	UnbindReason_Upgrade   UnbindReason = 2 // 插件升级
	UnbindReason_Downgrade UnbindReason = 3 // 插件降级
	UnbindReason_Panic     UnbindReason = 4 // 插件异常
)

// Enum value maps for UnbindReason.
var (
	UnbindReason_name = map[int32]string{
		0: "Exit",
		1: "UnUsed",
		2: "Upgrade",
		3: "Downgrade",
		4: "Panic",
	}
	UnbindReason_value = map[string]int32{
		"Exit":      0,
		"UnUsed":    1,
		"Upgrade":   2,
		"Downgrade": 3,
		"Panic":     4,
	}
)

func (x UnbindReason) Enum() *UnbindReason {
	p := new(UnbindReason)
	*p = x
	return p
}

func (x UnbindReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnbindReason) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[0].Descriptor()
}

func (UnbindReason) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[0]
}

func (x UnbindReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnbindReason.Descriptor instead.
func (UnbindReason) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

type CommunicateType int32

const (
	CommunicateType_Bind         CommunicateType = 0 // 绑定
	CommunicateType_Unbind       CommunicateType = 1 // 解绑
	CommunicateType_ExecRequest  CommunicateType = 2 // core: 执行函数请求
	CommunicateType_ExecResponse CommunicateType = 3 // core: 执行函数响应
	CommunicateType_Ping         CommunicateType = 4 // plugin: 健康检查消息
	CommunicateType_Log          CommunicateType = 5 // plugin: 日志消息
)

// Enum value maps for CommunicateType.
var (
	CommunicateType_name = map[int32]string{
		0: "Bind",
		1: "Unbind",
		2: "ExecRequest",
		3: "ExecResponse",
		4: "Ping",
		5: "Log",
	}
	CommunicateType_value = map[string]int32{
		"Bind":         0,
		"Unbind":       1,
		"ExecRequest":  2,
		"ExecResponse": 3,
		"Ping":         4,
		"Log":          5,
	}
)

func (x CommunicateType) Enum() *CommunicateType {
	p := new(CommunicateType)
	*p = x
	return p
}

func (x CommunicateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommunicateType) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[1].Descriptor()
}

func (CommunicateType) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[1]
}

func (x CommunicateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommunicateType.Descriptor instead.
func (CommunicateType) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{1}
}

type LogLevel int32

const (
	LogLevel_Debug LogLevel = 0 // 调试
	LogLevel_Info  LogLevel = 1 // 信息
	LogLevel_Warn  LogLevel = 2 // 警告
	LogLevel_Error LogLevel = 3 // 错误
)

// Enum value maps for LogLevel.
var (
	LogLevel_name = map[int32]string{
		0: "Debug",
		1: "Info",
		2: "Warn",
		3: "Error",
	}
	LogLevel_value = map[string]int32{
		"Debug": 0,
		"Info":  1,
		"Warn":  2,
		"Error": 3,
	}
)

func (x LogLevel) Enum() *LogLevel {
	p := new(LogLevel)
	*p = x
	return p
}

func (x LogLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LogLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[2].Descriptor()
}

func (LogLevel) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[2]
}

func (x LogLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LogLevel.Descriptor instead.
func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2}
}

type PluginStatus int32

const (
	PluginStatus_Connected    PluginStatus = 0 // 已连接
	PluginStatus_Disconnected PluginStatus = 1 // 未连接
)

// Enum value maps for PluginStatus.
var (
	PluginStatus_name = map[int32]string{
		0: "Connected",
		1: "Disconnected",
	}
	PluginStatus_value = map[string]int32{
		"Connected":    0,
		"Disconnected": 1,
	}
)

func (x PluginStatus) Enum() *PluginStatus {
	p := new(PluginStatus)
	*p = x
	return p
}

func (x PluginStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PluginStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[3].Descriptor()
}

func (PluginStatus) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[3]
}

func (x PluginStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PluginStatus.Descriptor instead.
func (PluginStatus) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{3}
}

type CoreStatus int32

const (
	CoreStatus_Launched CoreStatus = 0 // 已启动
	CoreStatus_Stopped  CoreStatus = 1 // 已停止
)

// Enum value maps for CoreStatus.
var (
	CoreStatus_name = map[int32]string{
		0: "Launched",
		1: "Stopped",
	}
	CoreStatus_value = map[string]int32{
		"Launched": 0,
		"Stopped":  1,
	}
)

func (x CoreStatus) Enum() *CoreStatus {
	p := new(CoreStatus)
	*p = x
	return p
}

func (x CoreStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CoreStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[4].Descriptor()
}

func (CoreStatus) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[4]
}

func (x CoreStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CoreStatus.Descriptor instead.
func (CoreStatus) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{4}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

type BindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`         // 绑定密钥
	Name      string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`           // 插件名称
	Version   string   `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`     // 插件版本号
	Functions []string `protobuf:"bytes,4,rep,name=Functions,proto3" json:"Functions,omitempty"` // 函数列表，用于校验是否实现core要求的interfaces
}

func (x *BindRequest) Reset() {
	*x = BindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRequest) ProtoMessage() {}

func (x *BindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRequest.ProtoReflect.Descriptor instead.
func (*BindRequest) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{1}
}

func (x *BindRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *BindRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BindRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *BindRequest) GetFunctions() []string {
	if x != nil {
		return x.Functions
	}
	return nil
}

type UnbindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason  UnbindReason `protobuf:"varint,1,opt,name=Reason,proto3,enum=pb.UnbindReason" json:"Reason,omitempty"` // 解绑原因
	Token   string       `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`                         // 绑定密钥
	Name    string       `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`                           // 插件名称
	Version string       `protobuf:"bytes,4,opt,name=Version,proto3" json:"Version,omitempty"`                     // 插件版本号
	Msg     *string      `protobuf:"bytes,5,opt,name=Msg,proto3,oneof" json:"Msg,omitempty"`                       // 原因描述(可选)
}

func (x *UnbindRequest) Reset() {
	*x = UnbindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnbindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnbindRequest) ProtoMessage() {}

func (x *UnbindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnbindRequest.ProtoReflect.Descriptor instead.
func (*UnbindRequest) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2}
}

func (x *UnbindRequest) GetReason() UnbindReason {
	if x != nil {
		return x.Reason
	}
	return UnbindReason_Exit
}

func (x *UnbindRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UnbindRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UnbindRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *UnbindRequest) GetMsg() string {
	if x != nil && x.Msg != nil {
		return *x.Msg
	}
	return ""
}

type LogInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    LogLevel `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.LogLevel" json:"Type,omitempty"`
	Message string   `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"` // 日志信息
}

func (x *LogInfo) Reset() {
	*x = LogInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogInfo) ProtoMessage() {}

func (x *LogInfo) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogInfo.ProtoReflect.Descriptor instead.
func (*LogInfo) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{3}
}

func (x *LogInfo) GetType() LogLevel {
	if x != nil {
		return x.Type
	}
	return LogLevel_Debug
}

func (x *LogInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CommunicateMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type CommunicateType `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.CommunicateType" json:"Type,omitempty"`
	Data []byte          `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"` // 数据
}

func (x *CommunicateMsg) Reset() {
	*x = CommunicateMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicateMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicateMsg) ProtoMessage() {}

func (x *CommunicateMsg) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicateMsg.ProtoReflect.Descriptor instead.
func (*CommunicateMsg) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{4}
}

func (x *CommunicateMsg) GetType() CommunicateType {
	if x != nil {
		return x.Type
	}
	return CommunicateType_Bind
}

func (x *CommunicateMsg) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// 消息类型为执行函数时，为以下特殊信息体
type CommunicateExecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`            // 执行ID
	FuncName string           `protobuf:"bytes,2,opt,name=FuncName,proto3" json:"FuncName,omitempty"` // 函数名
	Args     *structpb.Struct `protobuf:"bytes,3,opt,name=Args,proto3" json:"Args,omitempty"`         // 参数
}

func (x *CommunicateExecRequest) Reset() {
	*x = CommunicateExecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicateExecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicateExecRequest) ProtoMessage() {}

func (x *CommunicateExecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicateExecRequest.ProtoReflect.Descriptor instead.
func (*CommunicateExecRequest) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{5}
}

func (x *CommunicateExecRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CommunicateExecRequest) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *CommunicateExecRequest) GetArgs() *structpb.Struct {
	if x != nil {
		return x.Args
	}
	return nil
}

type CommunicateExecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     uint64           `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`        // 执行ID
	Result *structpb.Struct `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"` // 返回值
	Err    *string          `protobuf:"bytes,3,opt,name=Err,proto3,oneof" json:"Err,omitempty"` // 错误信息(可选)
}

func (x *CommunicateExecResponse) Reset() {
	*x = CommunicateExecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunicateExecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunicateExecResponse) ProtoMessage() {}

func (x *CommunicateExecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunicateExecResponse.ProtoReflect.Descriptor instead.
func (*CommunicateExecResponse) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{6}
}

func (x *CommunicateExecResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *CommunicateExecResponse) GetResult() *structpb.Struct {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CommunicateExecResponse) GetErr() string {
	if x != nil && x.Err != nil {
		return *x.Err
	}
	return ""
}

var File_conn_proto protoreflect.FileDescriptor

var file_conn_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x6f, 0x0a, 0x0b, 0x42, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x46, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0d, 0x55, 0x6e, 0x62,
	0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x4d, 0x73, 0x67, 0x22, 0x45, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x20, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4d,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67,
	0x12, 0x27, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x71, 0x0a,
	0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45, 0x78, 0x65, 0x63,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x41, 0x72, 0x67, 0x73,
	0x22, 0x79, 0x0a, 0x17, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2f, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x15, 0x0a, 0x03,
	0x45, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x45, 0x72, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x45, 0x72, 0x72, 0x2a, 0x4b, 0x0a, 0x0c, 0x55,
	0x6e, 0x62, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x45,
	0x78, 0x69, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x55, 0x73, 0x65, 0x64, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x44, 0x6f, 0x77, 0x6e, 0x67, 0x72, 0x61, 0x64, 0x65, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x10, 0x04, 0x2a, 0x5d, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x42,
	0x69, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x78, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x10, 0x04, 0x12, 0x07,
	0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x10, 0x05, 0x2a, 0x34, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x65, 0x62, 0x75, 0x67, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x61, 0x72, 0x6e,
	0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x2a, 0x2f, 0x0a,
	0x0c, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a,
	0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x10, 0x01, 0x2a, 0x27,
	0x0a, 0x0a, 0x43, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08,
	0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74,
	0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x01, 0x32, 0x66, 0x0a, 0x04, 0x43, 0x6f, 0x6e, 0x6e, 0x12,
	0x21, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x28, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x4d, 0x73, 0x67, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4d, 0x73, 0x67, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conn_proto_rawDescOnce sync.Once
	file_conn_proto_rawDescData = file_conn_proto_rawDesc
)

func file_conn_proto_rawDescGZIP() []byte {
	file_conn_proto_rawDescOnce.Do(func() {
		file_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_conn_proto_rawDescData)
	})
	return file_conn_proto_rawDescData
}

var file_conn_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_conn_proto_goTypes = []interface{}{
	(UnbindReason)(0),               // 0: pb.UnbindReason
	(CommunicateType)(0),            // 1: pb.CommunicateType
	(LogLevel)(0),                   // 2: pb.LogLevel
	(PluginStatus)(0),               // 3: pb.PluginStatus
	(CoreStatus)(0),                 // 4: pb.CoreStatus
	(*Empty)(nil),                   // 5: pb.Empty
	(*BindRequest)(nil),             // 6: pb.BindRequest
	(*UnbindRequest)(nil),           // 7: pb.UnbindRequest
	(*LogInfo)(nil),                 // 8: pb.LogInfo
	(*CommunicateMsg)(nil),          // 9: pb.CommunicateMsg
	(*CommunicateExecRequest)(nil),  // 10: pb.CommunicateExecRequest
	(*CommunicateExecResponse)(nil), // 11: pb.CommunicateExecResponse
	(*structpb.Struct)(nil),         // 12: google.protobuf.Struct
}
var file_conn_proto_depIdxs = []int32{
	0,  // 0: pb.UnbindRequest.Reason:type_name -> pb.UnbindReason
	2,  // 1: pb.LogInfo.Type:type_name -> pb.LogLevel
	1,  // 2: pb.CommunicateMsg.Type:type_name -> pb.CommunicateType
	12, // 3: pb.CommunicateExecRequest.Args:type_name -> google.protobuf.Struct
	12, // 4: pb.CommunicateExecResponse.Result:type_name -> google.protobuf.Struct
	8,  // 5: pb.Conn.Log:input_type -> pb.LogInfo
	9,  // 6: pb.Conn.Communicate:input_type -> pb.CommunicateMsg
	5,  // 7: pb.Conn.Log:output_type -> pb.Empty
	9,  // 8: pb.Conn.Communicate:output_type -> pb.CommunicateMsg
	7,  // [7:9] is the sub-list for method output_type
	5,  // [5:7] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_conn_proto_init() }
func file_conn_proto_init() {
	if File_conn_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_conn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindRequest); i {
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
		file_conn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnbindRequest); i {
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
		file_conn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogInfo); i {
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
		file_conn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicateMsg); i {
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
		file_conn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicateExecRequest); i {
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
		file_conn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunicateExecResponse); i {
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
	file_conn_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_conn_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_conn_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conn_proto_goTypes,
		DependencyIndexes: file_conn_proto_depIdxs,
		EnumInfos:         file_conn_proto_enumTypes,
		MessageInfos:      file_conn_proto_msgTypes,
	}.Build()
	File_conn_proto = out.File
	file_conn_proto_rawDesc = nil
	file_conn_proto_goTypes = nil
	file_conn_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConnClient is the client API for Conn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConnClient interface {
	Log(ctx context.Context, opts ...grpc.CallOption) (Conn_LogClient, error)
	Communicate(ctx context.Context, opts ...grpc.CallOption) (Conn_CommunicateClient, error)
}

type connClient struct {
	cc grpc.ClientConnInterface
}

func NewConnClient(cc grpc.ClientConnInterface) ConnClient {
	return &connClient{cc}
}

func (c *connClient) Log(ctx context.Context, opts ...grpc.CallOption) (Conn_LogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Conn_serviceDesc.Streams[0], "/pb.Conn/Log", opts...)
	if err != nil {
		return nil, err
	}
	x := &connLogClient{stream}
	return x, nil
}

type Conn_LogClient interface {
	Send(*LogInfo) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type connLogClient struct {
	grpc.ClientStream
}

func (x *connLogClient) Send(m *LogInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connLogClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *connClient) Communicate(ctx context.Context, opts ...grpc.CallOption) (Conn_CommunicateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Conn_serviceDesc.Streams[1], "/pb.Conn/Communicate", opts...)
	if err != nil {
		return nil, err
	}
	x := &connCommunicateClient{stream}
	return x, nil
}

type Conn_CommunicateClient interface {
	Send(*CommunicateMsg) error
	Recv() (*CommunicateMsg, error)
	grpc.ClientStream
}

type connCommunicateClient struct {
	grpc.ClientStream
}

func (x *connCommunicateClient) Send(m *CommunicateMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *connCommunicateClient) Recv() (*CommunicateMsg, error) {
	m := new(CommunicateMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConnServer is the server API for Conn service.
type ConnServer interface {
	Log(Conn_LogServer) error
	Communicate(Conn_CommunicateServer) error
}

// UnimplementedConnServer can be embedded to have forward compatible implementations.
type UnimplementedConnServer struct {
}

func (*UnimplementedConnServer) Log(Conn_LogServer) error {
	return status.Errorf(codes.Unimplemented, "method Log not implemented")
}
func (*UnimplementedConnServer) Communicate(Conn_CommunicateServer) error {
	return status.Errorf(codes.Unimplemented, "method Communicate not implemented")
}

func RegisterConnServer(s *grpc.Server, srv ConnServer) {
	s.RegisterService(&_Conn_serviceDesc, srv)
}

func _Conn_Log_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnServer).Log(&connLogServer{stream})
}

type Conn_LogServer interface {
	SendAndClose(*Empty) error
	Recv() (*LogInfo, error)
	grpc.ServerStream
}

type connLogServer struct {
	grpc.ServerStream
}

func (x *connLogServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connLogServer) Recv() (*LogInfo, error) {
	m := new(LogInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Conn_Communicate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ConnServer).Communicate(&connCommunicateServer{stream})
}

type Conn_CommunicateServer interface {
	Send(*CommunicateMsg) error
	Recv() (*CommunicateMsg, error)
	grpc.ServerStream
}

type connCommunicateServer struct {
	grpc.ServerStream
}

func (x *connCommunicateServer) Send(m *CommunicateMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *connCommunicateServer) Recv() (*CommunicateMsg, error) {
	m := new(CommunicateMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Conn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Conn",
	HandlerType: (*ConnServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Log",
			Handler:       _Conn_Log_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Communicate",
			Handler:       _Conn_Communicate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "conn.proto",
}
