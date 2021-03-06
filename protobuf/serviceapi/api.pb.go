// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/protobuf/serviceapi/api.proto

package serviceapi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request's data for the `EmitEvent` API.
//
// **Example:**
// ```json
// {
//   "token":     "__SERVICE_TOKEN_FROM_ENV__",
//   "eventKey":  "__EVENT_KEY__",
//   "eventData": "{\"foo\":\"hello\",\"bar\":false}"
// }
// ```
type EmitEventRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	EventKey             string   `protobuf:"bytes,2,opt,name=eventKey,proto3" json:"eventKey,omitempty"`
	EventData            string   `protobuf:"bytes,3,opt,name=eventData,proto3" json:"eventData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmitEventRequest) Reset()         { *m = EmitEventRequest{} }
func (m *EmitEventRequest) String() string { return proto.CompactTextString(m) }
func (*EmitEventRequest) ProtoMessage()    {}
func (*EmitEventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_722fa58223ac81c9, []int{0}
}
func (m *EmitEventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmitEventRequest.Unmarshal(m, b)
}
func (m *EmitEventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmitEventRequest.Marshal(b, m, deterministic)
}
func (dst *EmitEventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitEventRequest.Merge(dst, src)
}
func (m *EmitEventRequest) XXX_Size() int {
	return xxx_messageInfo_EmitEventRequest.Size(m)
}
func (m *EmitEventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitEventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmitEventRequest proto.InternalMessageInfo

func (m *EmitEventRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *EmitEventRequest) GetEventKey() string {
	if m != nil {
		return m.EventKey
	}
	return ""
}

func (m *EmitEventRequest) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

// Reply of `EmitEvent` API doesn't contain any data.
type EmitEventReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmitEventReply) Reset()         { *m = EmitEventReply{} }
func (m *EmitEventReply) String() string { return proto.CompactTextString(m) }
func (*EmitEventReply) ProtoMessage()    {}
func (*EmitEventReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_722fa58223ac81c9, []int{1}
}
func (m *EmitEventReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmitEventReply.Unmarshal(m, b)
}
func (m *EmitEventReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmitEventReply.Marshal(b, m, deterministic)
}
func (dst *EmitEventReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitEventReply.Merge(dst, src)
}
func (m *EmitEventReply) XXX_Size() int {
	return xxx_messageInfo_EmitEventReply.Size(m)
}
func (m *EmitEventReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitEventReply.DiscardUnknown(m)
}

var xxx_messageInfo_EmitEventReply proto.InternalMessageInfo

// The request's data for the `ListenTask` stream API.
//
// **Example:**
// ```json
// {
//   "token": "__SERVICE_TOKEN_FROM_ENV__"
// }
// ```
type ListenTaskRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListenTaskRequest) Reset()         { *m = ListenTaskRequest{} }
func (m *ListenTaskRequest) String() string { return proto.CompactTextString(m) }
func (*ListenTaskRequest) ProtoMessage()    {}
func (*ListenTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_722fa58223ac81c9, []int{2}
}
func (m *ListenTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListenTaskRequest.Unmarshal(m, b)
}
func (m *ListenTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListenTaskRequest.Marshal(b, m, deterministic)
}
func (dst *ListenTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListenTaskRequest.Merge(dst, src)
}
func (m *ListenTaskRequest) XXX_Size() int {
	return xxx_messageInfo_ListenTaskRequest.Size(m)
}
func (m *ListenTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListenTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListenTaskRequest proto.InternalMessageInfo

func (m *ListenTaskRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// The data received from the stream of the `ListenTask` API.
// The data will be received over time as long as the stream is open.
// The `executionID` value must be kept and sent with the result when calling the [`SubmitResult` API](#submitresult).
//
// **Example:**
// ```json
// {
//   "executionID": "__EXECUTION_ID__",
//   "taskKey":     "__TASK_KEY__",
//   "inputData":   "{\"inputX\":\"Hello world!\",\"inputY\":true}"
// }
// ```
type TaskData struct {
	ExecutionID          string   `protobuf:"bytes,1,opt,name=executionID,proto3" json:"executionID,omitempty"`
	TaskKey              string   `protobuf:"bytes,2,opt,name=taskKey,proto3" json:"taskKey,omitempty"`
	InputData            string   `protobuf:"bytes,3,opt,name=inputData,proto3" json:"inputData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskData) Reset()         { *m = TaskData{} }
func (m *TaskData) String() string { return proto.CompactTextString(m) }
func (*TaskData) ProtoMessage()    {}
func (*TaskData) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_722fa58223ac81c9, []int{3}
}
func (m *TaskData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskData.Unmarshal(m, b)
}
func (m *TaskData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskData.Marshal(b, m, deterministic)
}
func (dst *TaskData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskData.Merge(dst, src)
}
func (m *TaskData) XXX_Size() int {
	return xxx_messageInfo_TaskData.Size(m)
}
func (m *TaskData) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskData.DiscardUnknown(m)
}

var xxx_messageInfo_TaskData proto.InternalMessageInfo

func (m *TaskData) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *TaskData) GetTaskKey() string {
	if m != nil {
		return m.TaskKey
	}
	return ""
}

func (m *TaskData) GetInputData() string {
	if m != nil {
		return m.InputData
	}
	return ""
}

// The request's data for the `SubmitResult` API.
// The data must contain the `executionID` of the executed task received from the stream of [`ListenTask` API](#listentask).
//
// **Example:**
// ```json
// {
//   "executionID": "__EXECUTION_ID__",
//   "outputKey":   "__OUTPUT_KEY__",
//   "outputData":  "{\"foo\":\"super result\",\"bar\":true}"
// }
// ```
type SubmitResultRequest struct {
	ExecutionID          string   `protobuf:"bytes,1,opt,name=executionID,proto3" json:"executionID,omitempty"`
	OutputKey            string   `protobuf:"bytes,2,opt,name=outputKey,proto3" json:"outputKey,omitempty"`
	OutputData           string   `protobuf:"bytes,3,opt,name=outputData,proto3" json:"outputData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitResultRequest) Reset()         { *m = SubmitResultRequest{} }
func (m *SubmitResultRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitResultRequest) ProtoMessage()    {}
func (*SubmitResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_722fa58223ac81c9, []int{4}
}
func (m *SubmitResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResultRequest.Unmarshal(m, b)
}
func (m *SubmitResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResultRequest.Marshal(b, m, deterministic)
}
func (dst *SubmitResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResultRequest.Merge(dst, src)
}
func (m *SubmitResultRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitResultRequest.Size(m)
}
func (m *SubmitResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResultRequest proto.InternalMessageInfo

func (m *SubmitResultRequest) GetExecutionID() string {
	if m != nil {
		return m.ExecutionID
	}
	return ""
}

func (m *SubmitResultRequest) GetOutputKey() string {
	if m != nil {
		return m.OutputKey
	}
	return ""
}

func (m *SubmitResultRequest) GetOutputData() string {
	if m != nil {
		return m.OutputData
	}
	return ""
}

// Reply of `SubmitResult` API doesn't contain any data.
type SubmitResultReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitResultReply) Reset()         { *m = SubmitResultReply{} }
func (m *SubmitResultReply) String() string { return proto.CompactTextString(m) }
func (*SubmitResultReply) ProtoMessage()    {}
func (*SubmitResultReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_722fa58223ac81c9, []int{5}
}
func (m *SubmitResultReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitResultReply.Unmarshal(m, b)
}
func (m *SubmitResultReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitResultReply.Marshal(b, m, deterministic)
}
func (dst *SubmitResultReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitResultReply.Merge(dst, src)
}
func (m *SubmitResultReply) XXX_Size() int {
	return xxx_messageInfo_SubmitResultReply.Size(m)
}
func (m *SubmitResultReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitResultReply.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitResultReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmitEventRequest)(nil), "api.EmitEventRequest")
	proto.RegisterType((*EmitEventReply)(nil), "api.EmitEventReply")
	proto.RegisterType((*ListenTaskRequest)(nil), "api.ListenTaskRequest")
	proto.RegisterType((*TaskData)(nil), "api.TaskData")
	proto.RegisterType((*SubmitResultRequest)(nil), "api.SubmitResultRequest")
	proto.RegisterType((*SubmitResultReply)(nil), "api.SubmitResultReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// Emit an event to Core.
	// The event and its data must be defined in the [service's definition file](../guide/service/service-file.md).
	EmitEvent(ctx context.Context, in *EmitEventRequest, opts ...grpc.CallOption) (*EmitEventReply, error)
	// Subscribe to the stream of tasks to execute.
	// Every task received must be executed and its result must be submitted using the `SubmitResult` API.
	ListenTask(ctx context.Context, in *ListenTaskRequest, opts ...grpc.CallOption) (Service_ListenTaskClient, error)
	// Submit the result of a task's execution to Core.
	// The result must be defined as a task's output in the [service's definition file](../guide/service/service-file.md).
	SubmitResult(ctx context.Context, in *SubmitResultRequest, opts ...grpc.CallOption) (*SubmitResultReply, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) EmitEvent(ctx context.Context, in *EmitEventRequest, opts ...grpc.CallOption) (*EmitEventReply, error) {
	out := new(EmitEventReply)
	err := c.cc.Invoke(ctx, "/api.Service/EmitEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ListenTask(ctx context.Context, in *ListenTaskRequest, opts ...grpc.CallOption) (Service_ListenTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Service_serviceDesc.Streams[0], "/api.Service/ListenTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceListenTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_ListenTaskClient interface {
	Recv() (*TaskData, error)
	grpc.ClientStream
}

type serviceListenTaskClient struct {
	grpc.ClientStream
}

func (x *serviceListenTaskClient) Recv() (*TaskData, error) {
	m := new(TaskData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) SubmitResult(ctx context.Context, in *SubmitResultRequest, opts ...grpc.CallOption) (*SubmitResultReply, error) {
	out := new(SubmitResultReply)
	err := c.cc.Invoke(ctx, "/api.Service/SubmitResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// Emit an event to Core.
	// The event and its data must be defined in the [service's definition file](../guide/service/service-file.md).
	EmitEvent(context.Context, *EmitEventRequest) (*EmitEventReply, error)
	// Subscribe to the stream of tasks to execute.
	// Every task received must be executed and its result must be submitted using the `SubmitResult` API.
	ListenTask(*ListenTaskRequest, Service_ListenTaskServer) error
	// Submit the result of a task's execution to Core.
	// The result must be defined as a task's output in the [service's definition file](../guide/service/service-file.md).
	SubmitResult(context.Context, *SubmitResultRequest) (*SubmitResultReply, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_EmitEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).EmitEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/EmitEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).EmitEvent(ctx, req.(*EmitEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_ListenTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).ListenTask(m, &serviceListenTaskServer{stream})
}

type Service_ListenTaskServer interface {
	Send(*TaskData) error
	grpc.ServerStream
}

type serviceListenTaskServer struct {
	grpc.ServerStream
}

func (x *serviceListenTaskServer) Send(m *TaskData) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_SubmitResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).SubmitResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Service/SubmitResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).SubmitResult(ctx, req.(*SubmitResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmitEvent",
			Handler:    _Service_EmitEvent_Handler,
		},
		{
			MethodName: "SubmitResult",
			Handler:    _Service_SubmitResult_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenTask",
			Handler:       _Service_ListenTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/mesg-foundation/core/protobuf/serviceapi/api.proto",
}

func init() {
	proto.RegisterFile("github.com/mesg-foundation/core/protobuf/serviceapi/api.proto", fileDescriptor_api_722fa58223ac81c9)
}

var fileDescriptor_api_722fa58223ac81c9 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4f, 0xc2, 0x40,
	0x10, 0x85, 0xa9, 0x44, 0xa1, 0x23, 0x1a, 0x58, 0x94, 0x34, 0x0d, 0x31, 0xa4, 0x27, 0x3d, 0x48,
	0x8d, 0x1e, 0x8c, 0x07, 0x2f, 0x04, 0x0e, 0x46, 0x4f, 0xe0, 0xc9, 0xdb, 0xb6, 0x0c, 0xb8, 0x81,
	0x76, 0x57, 0x76, 0x97, 0xc8, 0xcf, 0xf3, 0x9f, 0x99, 0x5d, 0xa4, 0x2d, 0x48, 0xf4, 0xd6, 0x79,
	0xf3, 0xd2, 0x37, 0xfb, 0xcd, 0xc0, 0xe3, 0x94, 0xa9, 0x77, 0x1d, 0x75, 0x63, 0x9e, 0x84, 0x09,
	0xca, 0xe9, 0xf5, 0x84, 0xeb, 0x74, 0x4c, 0x15, 0xe3, 0x69, 0x18, 0xf3, 0x05, 0x86, 0x62, 0xc1,
	0x15, 0x8f, 0xf4, 0x24, 0x94, 0xb8, 0x58, 0xb2, 0x18, 0xa9, 0x60, 0x21, 0x15, 0xac, 0x6b, 0x75,
	0x52, 0xa6, 0x82, 0x05, 0x11, 0xd4, 0x07, 0x09, 0x53, 0x83, 0x25, 0xa6, 0x6a, 0x88, 0x1f, 0x1a,
	0xa5, 0x22, 0x67, 0x70, 0xa8, 0xf8, 0x0c, 0x53, 0xcf, 0xe9, 0x38, 0x97, 0xee, 0x70, 0x5d, 0x10,
	0x1f, 0xaa, 0x68, 0x5c, 0xcf, 0xb8, 0xf2, 0x0e, 0x6c, 0x23, 0xab, 0x49, 0x1b, 0x5c, 0xfb, 0xdd,
	0xa7, 0x8a, 0x7a, 0x65, 0xdb, 0xcc, 0x85, 0xa0, 0x0e, 0xa7, 0x85, 0x0c, 0x31, 0x5f, 0x05, 0x57,
	0xd0, 0x78, 0x61, 0x52, 0x61, 0xfa, 0x4a, 0xe5, 0xec, 0xcf, 0xd8, 0x60, 0x0c, 0x55, 0x63, 0x32,
	0x3f, 0x22, 0x1d, 0x38, 0xc6, 0x4f, 0x8c, 0xb5, 0x79, 0xe1, 0x53, 0xff, 0xc7, 0x57, 0x94, 0x88,
	0x07, 0x15, 0x45, 0xe5, 0x2c, 0x9f, 0x71, 0x53, 0x9a, 0x11, 0x59, 0x2a, 0xf4, 0xd6, 0x88, 0x99,
	0x10, 0x68, 0x68, 0x8e, 0x74, 0x94, 0x30, 0x35, 0x44, 0xa9, 0xe7, 0x19, 0x89, 0xff, 0x03, 0xdb,
	0xe0, 0x72, 0xad, 0x84, 0x2e, 0x60, 0xc9, 0x05, 0x72, 0x01, 0xb0, 0x2e, 0x0a, 0xa9, 0x05, 0x25,
	0x68, 0x42, 0x63, 0x3b, 0x56, 0xcc, 0x57, 0xb7, 0x5f, 0x0e, 0x54, 0x46, 0xeb, 0x85, 0x91, 0x07,
	0x70, 0x33, 0x74, 0xe4, 0xbc, 0x6b, 0x96, 0xb7, 0xbb, 0x2e, 0xbf, 0xb9, 0x2b, 0x1b, 0xc2, 0x25,
	0x72, 0x0f, 0x90, 0x33, 0x26, 0x2d, 0x6b, 0xfa, 0x05, 0xdd, 0x3f, 0xb1, 0xfa, 0x86, 0x70, 0x50,
	0xba, 0x71, 0x48, 0x0f, 0x6a, 0xc5, 0xa1, 0x88, 0x67, 0x2d, 0x7b, 0xf0, 0xf8, 0xad, 0x3d, 0x1d,
	0x1b, 0xde, 0xab, 0xbd, 0x41, 0x7e, 0x73, 0xd1, 0x91, 0x3d, 0xb8, 0xbb, 0xef, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb3, 0x49, 0xee, 0xd0, 0xb1, 0x02, 0x00, 0x00,
}
