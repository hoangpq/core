// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/mesg-foundation/core/interface/grpc/core/api/service.proto

package api

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

// This is the definition of a MESG Service.
type Service struct {
	ID                   string        `protobuf:"bytes,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string        `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tasks                []*Task       `protobuf:"bytes,5,rep,name=tasks,proto3" json:"tasks,omitempty"`
	Events               []*Event      `protobuf:"bytes,6,rep,name=events,proto3" json:"events,omitempty"`
	Dependencies         []*Dependency `protobuf:"bytes,7,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Configuration        *Dependency   `protobuf:"bytes,8,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Repository           string        `protobuf:"bytes,9,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d9a9e98865b3866f, []int{0}
}
func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (dst *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(dst, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Service) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Service) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Service) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *Service) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Service) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Service) GetConfiguration() *Dependency {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func (m *Service) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

// Events are emitted by the service whenever the service wants.
// TODO(ilgooz) remove key, serviceName fields when Event type crafted manually.
type Event struct {
	Key                  string       `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Data                 []*Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d9a9e98865b3866f, []int{1}
}
func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (dst *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(dst, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Event) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Event) GetData() []*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// A task is a function that requires inputs and returns output.
// TODO(ilgooz) remove key, serviceName fields when Task type crafted manually.
type Task struct {
	Key                  string       `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Inputs               []*Parameter `protobuf:"bytes,6,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs              []*Output    `protobuf:"bytes,7,rep,name=outputs,proto3" json:"outputs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d9a9e98865b3866f, []int{2}
}
func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (dst *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(dst, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Task) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Task) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Task) GetInputs() []*Parameter {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *Task) GetOutputs() []*Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// A output is the data a task must return.
// TODO(ilgooz) remove key, taskKey, serviceName fields when Output type crafted manually.
type Output struct {
	Key                  string       `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Data                 []*Parameter `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}
func (*Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d9a9e98865b3866f, []int{3}
}
func (m *Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Output.Unmarshal(m, b)
}
func (m *Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Output.Marshal(b, m, deterministic)
}
func (dst *Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Output.Merge(dst, src)
}
func (m *Output) XXX_Size() int {
	return xxx_messageInfo_Output.Size(m)
}
func (m *Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Output proto.InternalMessageInfo

func (m *Output) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Output) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Output) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Output) GetData() []*Parameter {
	if m != nil {
		return m.Data
	}
	return nil
}

// A parameter is the definition of a specific value.
type Parameter struct {
	Key                  string   `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Optional             bool     `protobuf:"varint,4,opt,name=optional,proto3" json:"optional,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Parameter) Reset()         { *m = Parameter{} }
func (m *Parameter) String() string { return proto.CompactTextString(m) }
func (*Parameter) ProtoMessage()    {}
func (*Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d9a9e98865b3866f, []int{4}
}
func (m *Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameter.Unmarshal(m, b)
}
func (m *Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameter.Marshal(b, m, deterministic)
}
func (dst *Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameter.Merge(dst, src)
}
func (m *Parameter) XXX_Size() int {
	return xxx_messageInfo_Parameter.Size(m)
}
func (m *Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_Parameter proto.InternalMessageInfo

func (m *Parameter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Parameter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Parameter) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Parameter) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Parameter) GetOptional() bool {
	if m != nil {
		return m.Optional
	}
	return false
}

// A dependency is a configuration of an other Docker container that runs separately from the service.
type Dependency struct {
	Key                  string   `protobuf:"bytes,8,opt,name=key,proto3" json:"key,omitempty"`
	Image                string   `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Volumes              []string `protobuf:"bytes,2,rep,name=volumes,proto3" json:"volumes,omitempty"`
	Volumesfrom          []string `protobuf:"bytes,3,rep,name=volumesfrom,proto3" json:"volumesfrom,omitempty"`
	Ports                []string `protobuf:"bytes,4,rep,name=ports,proto3" json:"ports,omitempty"`
	Command              string   `protobuf:"bytes,5,opt,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}
func (*Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_d9a9e98865b3866f, []int{5}
}
func (m *Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dependency.Unmarshal(m, b)
}
func (m *Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dependency.Marshal(b, m, deterministic)
}
func (dst *Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dependency.Merge(dst, src)
}
func (m *Dependency) XXX_Size() int {
	return xxx_messageInfo_Dependency.Size(m)
}
func (m *Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Dependency proto.InternalMessageInfo

func (m *Dependency) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Dependency) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Dependency) GetVolumes() []string {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func (m *Dependency) GetVolumesfrom() []string {
	if m != nil {
		return m.Volumesfrom
	}
	return nil
}

func (m *Dependency) GetPorts() []string {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *Dependency) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*Service)(nil), "api.Service")
	proto.RegisterType((*Event)(nil), "api.Event")
	proto.RegisterType((*Task)(nil), "api.Task")
	proto.RegisterType((*Output)(nil), "api.Output")
	proto.RegisterType((*Parameter)(nil), "api.Parameter")
	proto.RegisterType((*Dependency)(nil), "api.Dependency")
}

func init() {
	proto.RegisterFile("github.com/mesg-foundation/core/interface/grpc/core/api/service.proto", fileDescriptor_service_d9a9e98865b3866f)
}

var fileDescriptor_service_d9a9e98865b3866f = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x93, 0xcb, 0x8e, 0xd3, 0x4c,
	0x10, 0x85, 0xe5, 0x6b, 0xe2, 0xca, 0xff, 0x0f, 0xa8, 0xc5, 0xa2, 0xc5, 0x02, 0x2c, 0x4b, 0xa0,
	0x6c, 0x88, 0x25, 0x46, 0xbc, 0x41, 0x66, 0x31, 0x2b, 0x90, 0xe1, 0x05, 0x7a, 0xec, 0x8a, 0x69,
	0x65, 0xfa, 0xa2, 0xee, 0x76, 0xa4, 0xac, 0x79, 0x0e, 0xa4, 0x79, 0x54, 0xe4, 0x72, 0x1c, 0x32,
	0x68, 0x56, 0x80, 0xd8, 0x55, 0x7d, 0xe7, 0x74, 0x57, 0xf5, 0x91, 0x0d, 0x37, 0xbd, 0x0c, 0x5f,
	0x87, 0xbb, 0x4d, 0x6b, 0x54, 0xad, 0xd0, 0xf7, 0xef, 0x76, 0x66, 0xd0, 0x9d, 0x08, 0xd2, 0xe8,
	0xba, 0x35, 0x0e, 0x6b, 0xa9, 0x03, 0xba, 0x9d, 0x68, 0xb1, 0xee, 0x9d, 0x6d, 0x27, 0x26, 0xac,
	0xac, 0x3d, 0xba, 0x83, 0x6c, 0x71, 0x63, 0x9d, 0x09, 0x86, 0x25, 0xc2, 0xca, 0xea, 0x21, 0x86,
	0xc5, 0xe7, 0x09, 0xb3, 0x2b, 0x88, 0x6f, 0xb7, 0x1c, 0xca, 0x68, 0x5d, 0x34, 0xf1, 0xed, 0x96,
	0x31, 0x48, 0xb5, 0x50, 0xc8, 0x23, 0x22, 0x54, 0xb3, 0x12, 0x56, 0x1d, 0xfa, 0xd6, 0x49, 0x3b,
	0x8e, 0xe3, 0x31, 0x49, 0x97, 0x88, 0xbd, 0x86, 0x2c, 0x08, 0xbf, 0xf7, 0x3c, 0x2b, 0x93, 0xf5,
	0xea, 0x7d, 0xb1, 0x11, 0x56, 0x6e, 0xbe, 0x08, 0xbf, 0x6f, 0x26, 0xce, 0x2a, 0xc8, 0xf1, 0x80,
	0x3a, 0x78, 0x9e, 0x93, 0x03, 0xc8, 0x71, 0x33, 0xa2, 0xe6, 0xa4, 0xb0, 0x6b, 0xf8, 0xaf, 0x43,
	0x8b, 0xba, 0x43, 0xdd, 0x4a, 0xf4, 0x7c, 0x41, 0xce, 0x67, 0xe4, 0xdc, 0xce, 0xc2, 0xb1, 0x79,
	0x64, 0x62, 0x1f, 0xe0, 0xff, 0xd6, 0xe8, 0x9d, 0xec, 0x07, 0x47, 0x61, 0xf0, 0x65, 0x19, 0x3d,
	0x75, 0xea, 0xb1, 0x8b, 0xbd, 0x02, 0x70, 0x68, 0x8d, 0x97, 0xc1, 0xb8, 0x23, 0x2f, 0xe8, 0x45,
	0x17, 0xa4, 0xf2, 0x90, 0xd1, 0x72, 0xec, 0x39, 0x24, 0x7b, 0x3c, 0xf2, 0x94, 0x1c, 0x63, 0xf9,
	0x9b, 0x09, 0x55, 0x90, 0x76, 0x22, 0x08, 0x9e, 0xd0, 0xa3, 0xae, 0x68, 0xbd, 0x4f, 0xc2, 0x09,
	0x85, 0x01, 0x5d, 0x43, 0x5a, 0xf5, 0x3d, 0x82, 0x74, 0x0c, 0x6d, 0x1e, 0xba, 0xfc, 0xd3, 0xa1,
	0x6f, 0x21, 0x97, 0xda, 0x0e, 0xe7, 0xd4, 0x7f, 0x1d, 0x7b, 0x52, 0xd9, 0x1b, 0x58, 0x98, 0x21,
	0x90, 0x71, 0x0a, 0x7d, 0x45, 0xc6, 0x8f, 0xc4, 0x9a, 0x59, 0xab, 0x02, 0xe4, 0x13, 0xfa, 0xa7,
	0xa9, 0x7c, 0x8b, 0xa0, 0x38, 0xb3, 0xbf, 0x16, 0x0d, 0x83, 0x34, 0x1c, 0x2d, 0xf2, 0x64, 0x3a,
	0x35, 0xd6, 0xec, 0x25, 0x2c, 0x0d, 0xa9, 0xe2, 0x9e, 0x9e, 0xb6, 0x6c, 0xce, 0x7d, 0xf5, 0x10,
	0x01, 0xfc, 0xfc, 0x9c, 0x9e, 0x58, 0xe3, 0x05, 0x64, 0x52, 0x89, 0x7e, 0xde, 0x63, 0x6a, 0x18,
	0x87, 0xc5, 0xc1, 0xdc, 0x0f, 0x0a, 0x3d, 0x8f, 0xcb, 0x64, 0x5d, 0x34, 0x73, 0x3b, 0xae, 0x78,
	0x2a, 0x77, 0xce, 0x28, 0x4a, 0xa0, 0x68, 0x2e, 0xd1, 0x78, 0xa3, 0x35, 0x2e, 0x78, 0x9e, 0x92,
	0x36, 0x35, 0xe3, 0x8d, 0xad, 0x51, 0x4a, 0xe8, 0x8e, 0x67, 0x34, 0x69, 0x6e, 0xef, 0x72, 0xfa,
	0xc5, 0xaf, 0x7f, 0x04, 0x00, 0x00, 0xff, 0xff, 0x38, 0x51, 0xe4, 0x72, 0x2b, 0x04, 0x00, 0x00,
}
