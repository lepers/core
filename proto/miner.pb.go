// Code generated by protoc-gen-go. DO NOT EDIT.
// source: miner.proto

package sonm

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

type NATType int32

const (
	NATType_NONE                   NATType = 0
	NATType_BLOCKED                NATType = 1
	NATType_FULL                   NATType = 2
	NATType_SYMMETRIC              NATType = 3
	NATType_RESTRICTED             NATType = 4
	NATType_PORT_RESTRICTED        NATType = 5
	NATType_SYMMETRIC_UDP_FIREWALL NATType = 6
	NATType_UNKNOWN                NATType = 7
)

var NATType_name = map[int32]string{
	0: "NONE",
	1: "BLOCKED",
	2: "FULL",
	3: "SYMMETRIC",
	4: "RESTRICTED",
	5: "PORT_RESTRICTED",
	6: "SYMMETRIC_UDP_FIREWALL",
	7: "UNKNOWN",
}
var NATType_value = map[string]int32{
	"NONE":                   0,
	"BLOCKED":                1,
	"FULL":                   2,
	"SYMMETRIC":              3,
	"RESTRICTED":             4,
	"PORT_RESTRICTED":        5,
	"SYMMETRIC_UDP_FIREWALL": 6,
	"UNKNOWN":                7,
}

func (x NATType) String() string {
	return proto.EnumName(NATType_name, int32(x))
}
func (NATType) EnumDescriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MinerInfoRequest struct {
}

func (m *MinerInfoRequest) Reset()                    { *m = MinerInfoRequest{} }
func (m *MinerInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerInfoRequest) ProtoMessage()               {}
func (*MinerInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

type MinerHandshakeRequest struct {
	Hub   string      `protobuf:"bytes,1,opt,name=hub" json:"hub,omitempty"`
	Tasks []*TaskInfo `protobuf:"bytes,2,rep,name=tasks" json:"tasks,omitempty"`
}

func (m *MinerHandshakeRequest) Reset()                    { *m = MinerHandshakeRequest{} }
func (m *MinerHandshakeRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeRequest) ProtoMessage()               {}
func (*MinerHandshakeRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *MinerHandshakeRequest) GetHub() string {
	if m != nil {
		return m.Hub
	}
	return ""
}

func (m *MinerHandshakeRequest) GetTasks() []*TaskInfo {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type MinerHandshakeReply struct {
	Miner        string        `protobuf:"bytes,1,opt,name=miner" json:"miner,omitempty"`
	Capabilities *Capabilities `protobuf:"bytes,2,opt,name=capabilities" json:"capabilities,omitempty"`
	NatType      NATType       `protobuf:"varint,3,opt,name=natType,enum=sonm.NATType" json:"natType,omitempty"`
}

func (m *MinerHandshakeReply) Reset()                    { *m = MinerHandshakeReply{} }
func (m *MinerHandshakeReply) String() string            { return proto.CompactTextString(m) }
func (*MinerHandshakeReply) ProtoMessage()               {}
func (*MinerHandshakeReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *MinerHandshakeReply) GetMiner() string {
	if m != nil {
		return m.Miner
	}
	return ""
}

func (m *MinerHandshakeReply) GetCapabilities() *Capabilities {
	if m != nil {
		return m.Capabilities
	}
	return nil
}

func (m *MinerHandshakeReply) GetNatType() NATType {
	if m != nil {
		return m.NatType
	}
	return NATType_NONE
}

type MinerStartRequest struct {
	Id            string                    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Registry      string                    `protobuf:"bytes,2,opt,name=registry" json:"registry,omitempty"`
	Image         string                    `protobuf:"bytes,3,opt,name=image" json:"image,omitempty"`
	Auth          string                    `protobuf:"bytes,4,opt,name=auth" json:"auth,omitempty"`
	RestartPolicy *ContainerRestartPolicy   `protobuf:"bytes,5,opt,name=restartPolicy" json:"restartPolicy,omitempty"`
	Usage         *TaskResourceRequirements `protobuf:"bytes,6,opt,name=usage" json:"usage,omitempty"`
	PublicKeyData string                    `protobuf:"bytes,7,opt,name=PublicKeyData" json:"PublicKeyData,omitempty"`
	CommitOnStop  bool                      `protobuf:"varint,8,opt,name=commitOnStop" json:"commitOnStop,omitempty"`
	Env           map[string]string         `protobuf:"bytes,9,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MinerStartRequest) Reset()                    { *m = MinerStartRequest{} }
func (m *MinerStartRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStartRequest) ProtoMessage()               {}
func (*MinerStartRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *MinerStartRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MinerStartRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

func (m *MinerStartRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *MinerStartRequest) GetAuth() string {
	if m != nil {
		return m.Auth
	}
	return ""
}

func (m *MinerStartRequest) GetRestartPolicy() *ContainerRestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return nil
}

func (m *MinerStartRequest) GetUsage() *TaskResourceRequirements {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *MinerStartRequest) GetPublicKeyData() string {
	if m != nil {
		return m.PublicKeyData
	}
	return ""
}

func (m *MinerStartRequest) GetCommitOnStop() bool {
	if m != nil {
		return m.CommitOnStop
	}
	return false
}

func (m *MinerStartRequest) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

type MinerStartReply struct {
	Container string                          `protobuf:"bytes,1,opt,name=container" json:"container,omitempty"`
	Ports     map[string]*MinerStartReplyPort `protobuf:"bytes,2,rep,name=Ports" json:"Ports,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *MinerStartReply) Reset()                    { *m = MinerStartReply{} }
func (m *MinerStartReply) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReply) ProtoMessage()               {}
func (*MinerStartReply) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *MinerStartReply) GetContainer() string {
	if m != nil {
		return m.Container
	}
	return ""
}

func (m *MinerStartReply) GetPorts() map[string]*MinerStartReplyPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

type MinerStartReplyPort struct {
	IP   string `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Port string `protobuf:"bytes,2,opt,name=port" json:"port,omitempty"`
}

func (m *MinerStartReplyPort) Reset()                    { *m = MinerStartReplyPort{} }
func (m *MinerStartReplyPort) String() string            { return proto.CompactTextString(m) }
func (*MinerStartReplyPort) ProtoMessage()               {}
func (*MinerStartReplyPort) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4, 0} }

func (m *MinerStartReplyPort) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *MinerStartReplyPort) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

type TaskInfo struct {
	Request *MinerStartRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Reply   *MinerStartReply   `protobuf:"bytes,2,opt,name=reply" json:"reply,omitempty"`
}

func (m *TaskInfo) Reset()                    { *m = TaskInfo{} }
func (m *TaskInfo) String() string            { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()               {}
func (*TaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *TaskInfo) GetRequest() *MinerStartRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *TaskInfo) GetReply() *MinerStartReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type MinerStatusMapRequest struct {
}

func (m *MinerStatusMapRequest) Reset()                    { *m = MinerStatusMapRequest{} }
func (m *MinerStatusMapRequest) String() string            { return proto.CompactTextString(m) }
func (*MinerStatusMapRequest) ProtoMessage()               {}
func (*MinerStatusMapRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func init() {
	proto.RegisterType((*MinerInfoRequest)(nil), "sonm.MinerInfoRequest")
	proto.RegisterType((*MinerHandshakeRequest)(nil), "sonm.MinerHandshakeRequest")
	proto.RegisterType((*MinerHandshakeReply)(nil), "sonm.MinerHandshakeReply")
	proto.RegisterType((*MinerStartRequest)(nil), "sonm.MinerStartRequest")
	proto.RegisterType((*MinerStartReply)(nil), "sonm.MinerStartReply")
	proto.RegisterType((*MinerStartReplyPort)(nil), "sonm.MinerStartReply.port")
	proto.RegisterType((*TaskInfo)(nil), "sonm.TaskInfo")
	proto.RegisterType((*MinerStatusMapRequest)(nil), "sonm.MinerStatusMapRequest")
	proto.RegisterEnum("sonm.NATType", NATType_name, NATType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Miner service

type MinerClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error)
	Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error)
	Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error)
	Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error)
	TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error)
	TaskDetails(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error)
	TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error)
	DiscoverHub(ctx context.Context, in *DiscoverHubRequest, opts ...grpc.CallOption) (*EmptyReply, error)
}

type minerClient struct {
	cc *grpc.ClientConn
}

func NewMinerClient(cc *grpc.ClientConn) MinerClient {
	return &minerClient{cc}
}

func (c *minerClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Info(ctx context.Context, in *MinerInfoRequest, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Info", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Handshake(ctx context.Context, in *MinerHandshakeRequest, opts ...grpc.CallOption) (*MinerHandshakeReply, error) {
	out := new(MinerHandshakeReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Handshake", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Start(ctx context.Context, in *MinerStartRequest, opts ...grpc.CallOption) (*MinerStartReply, error) {
	out := new(MinerStartReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) Stop(ctx context.Context, in *StopTaskRequest, opts ...grpc.CallOption) (*StopTaskReply, error) {
	out := new(StopTaskReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TasksStatus(ctx context.Context, opts ...grpc.CallOption) (Miner_TasksStatusClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[0], c.cc, "/sonm.Miner/TasksStatus", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTasksStatusClient{stream}
	return x, nil
}

type Miner_TasksStatusClient interface {
	Send(*MinerStatusMapRequest) error
	Recv() (*StatusMapReply, error)
	grpc.ClientStream
}

type minerTasksStatusClient struct {
	grpc.ClientStream
}

func (x *minerTasksStatusClient) Send(m *MinerStatusMapRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *minerTasksStatusClient) Recv() (*StatusMapReply, error) {
	m := new(StatusMapReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) TaskDetails(ctx context.Context, in *TaskStatusRequest, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/TaskDetails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerClient) TaskLogs(ctx context.Context, in *TaskLogsRequest, opts ...grpc.CallOption) (Miner_TaskLogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Miner_serviceDesc.Streams[1], c.cc, "/sonm.Miner/TaskLogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &minerTaskLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Miner_TaskLogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type minerTaskLogsClient struct {
	grpc.ClientStream
}

func (x *minerTaskLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *minerClient) DiscoverHub(ctx context.Context, in *DiscoverHubRequest, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := grpc.Invoke(ctx, "/sonm.Miner/DiscoverHub", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Miner service

type MinerServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	Info(context.Context, *MinerInfoRequest) (*InfoReply, error)
	Handshake(context.Context, *MinerHandshakeRequest) (*MinerHandshakeReply, error)
	Start(context.Context, *MinerStartRequest) (*MinerStartReply, error)
	Stop(context.Context, *StopTaskRequest) (*StopTaskReply, error)
	TasksStatus(Miner_TasksStatusServer) error
	TaskDetails(context.Context, *TaskStatusRequest) (*TaskStatusReply, error)
	TaskLogs(*TaskLogsRequest, Miner_TaskLogsServer) error
	DiscoverHub(context.Context, *DiscoverHubRequest) (*EmptyReply, error)
}

func RegisterMinerServer(s *grpc.Server, srv MinerServer) {
	s.RegisterService(&_Miner_serviceDesc, srv)
}

func _Miner_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Info(ctx, req.(*MinerInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerHandshakeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Handshake(ctx, req.(*MinerHandshakeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinerStartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Start(ctx, req.(*MinerStartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).Stop(ctx, req.(*StopTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TasksStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MinerServer).TasksStatus(&minerTasksStatusServer{stream})
}

type Miner_TasksStatusServer interface {
	Send(*StatusMapReply) error
	Recv() (*MinerStatusMapRequest, error)
	grpc.ServerStream
}

type minerTasksStatusServer struct {
	grpc.ServerStream
}

func (x *minerTasksStatusServer) Send(m *StatusMapReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *minerTasksStatusServer) Recv() (*MinerStatusMapRequest, error) {
	m := new(MinerStatusMapRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Miner_TaskDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).TaskDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/TaskDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).TaskDetails(ctx, req.(*TaskStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Miner_TaskLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MinerServer).TaskLogs(m, &minerTaskLogsServer{stream})
}

type Miner_TaskLogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type minerTaskLogsServer struct {
	grpc.ServerStream
}

func (x *minerTaskLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Miner_DiscoverHub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverHubRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerServer).DiscoverHub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Miner/DiscoverHub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerServer).DiscoverHub(ctx, req.(*DiscoverHubRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Miner_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Miner",
	HandlerType: (*MinerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Miner_Ping_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _Miner_Info_Handler,
		},
		{
			MethodName: "Handshake",
			Handler:    _Miner_Handshake_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Miner_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Miner_Stop_Handler,
		},
		{
			MethodName: "TaskDetails",
			Handler:    _Miner_TaskDetails_Handler,
		},
		{
			MethodName: "DiscoverHub",
			Handler:    _Miner_DiscoverHub_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TasksStatus",
			Handler:       _Miner_TasksStatus_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TaskLogs",
			Handler:       _Miner_TaskLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "miner.proto",
}

func init() { proto.RegisterFile("miner.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 846 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0x5f, 0x6f, 0xe2, 0x46,
	0x10, 0x8f, 0xc1, 0x04, 0x18, 0x2e, 0x89, 0x6f, 0x72, 0x7f, 0x5c, 0xf7, 0x54, 0x21, 0xeb, 0xa4,
	0xa2, 0x6b, 0x85, 0x72, 0xf4, 0x14, 0xb5, 0xd7, 0x87, 0xea, 0x2e, 0x38, 0x3a, 0x14, 0x02, 0x68,
	0x21, 0x3a, 0xf5, 0x29, 0x5a, 0xc8, 0x36, 0xd9, 0x02, 0xb6, 0xeb, 0x5d, 0x47, 0xf2, 0x27, 0xe8,
	0x4b, 0xbf, 0x40, 0x5f, 0xfb, 0xed, 0xfa, 0x2d, 0xaa, 0xdd, 0xb5, 0x13, 0x43, 0xb9, 0x27, 0x76,
	0x7e, 0x33, 0xf3, 0x9b, 0xdf, 0xcc, 0xec, 0x1a, 0x68, 0xad, 0x79, 0xc8, 0x92, 0x6e, 0x9c, 0x44,
	0x32, 0x42, 0x5b, 0x44, 0xe1, 0xda, 0x3b, 0xe2, 0xa1, 0xfa, 0x0d, 0x39, 0x35, 0xb0, 0x87, 0x0b,
	0x1a, 0xd3, 0x39, 0x5f, 0x71, 0xc9, 0x99, 0x30, 0x98, 0x8f, 0xe0, 0x5c, 0xaa, 0xcc, 0x41, 0xf8,
	0x5b, 0x44, 0xd8, 0x1f, 0x29, 0x13, 0xd2, 0x1f, 0xc3, 0x73, 0x8d, 0x7d, 0xa2, 0xe1, 0x8d, 0xb8,
	0xa3, 0x4b, 0x96, 0x3b, 0xd0, 0x81, 0xea, 0x5d, 0x3a, 0x77, 0xad, 0xb6, 0xd5, 0x69, 0x12, 0x75,
	0xc4, 0xd7, 0x50, 0x93, 0x54, 0x2c, 0x85, 0x5b, 0x69, 0x57, 0x3b, 0xad, 0xde, 0x61, 0x57, 0x55,
	0xec, 0xce, 0xa8, 0x58, 0x6a, 0x42, 0xe3, 0xf4, 0xff, 0xb2, 0xe0, 0x78, 0x9b, 0x31, 0x5e, 0x65,
	0xf8, 0x0c, 0x6a, 0x5a, 0x76, 0xce, 0x68, 0x0c, 0x3c, 0x85, 0x27, 0x65, 0xa1, 0x6e, 0xa5, 0x6d,
	0x75, 0x5a, 0x3d, 0x34, 0xd4, 0x67, 0x25, 0x0f, 0xd9, 0x88, 0xc3, 0x6f, 0xa1, 0x1e, 0x52, 0x39,
	0xcb, 0x62, 0xe6, 0x56, 0xdb, 0x56, 0xe7, 0xb0, 0x77, 0x60, 0x52, 0x46, 0x1f, 0x66, 0x0a, 0x24,
	0x85, 0xd7, 0xff, 0xbb, 0x0a, 0x4f, 0xb5, 0x9c, 0xa9, 0xa4, 0x89, 0x2c, 0x9a, 0x3b, 0x84, 0x0a,
	0xbf, 0xc9, 0x95, 0x54, 0xf8, 0x0d, 0x7a, 0xd0, 0x48, 0xd8, 0x2d, 0x17, 0x32, 0xc9, 0xb4, 0x84,
	0x26, 0x79, 0xb0, 0x95, 0x70, 0xbe, 0xa6, 0xb7, 0xa6, 0x50, 0x93, 0x18, 0x03, 0x11, 0x6c, 0x9a,
	0xca, 0x3b, 0xd7, 0xd6, 0xa0, 0x3e, 0xe3, 0x47, 0x38, 0x48, 0x98, 0x50, 0x75, 0x26, 0xd1, 0x8a,
	0x2f, 0x32, 0xb7, 0xa6, 0xbb, 0x79, 0x95, 0x77, 0x13, 0x85, 0x92, 0x2a, 0x25, 0xa4, 0x1c, 0x43,
	0x36, 0x53, 0xf0, 0x1d, 0xd4, 0x52, 0xa1, 0xaa, 0xed, 0xeb, 0xdc, 0x6f, 0x1e, 0x87, 0x4c, 0x98,
	0x88, 0xd2, 0x64, 0xa1, 0x17, 0xc4, 0x13, 0xb6, 0x66, 0xa1, 0x14, 0xc4, 0x04, 0xe3, 0x6b, 0x38,
	0x98, 0xa4, 0xf3, 0x15, 0x5f, 0x5c, 0xb0, 0xac, 0x4f, 0x25, 0x75, 0xeb, 0x5a, 0xd6, 0x26, 0x88,
	0x3e, 0x3c, 0x59, 0x44, 0xeb, 0x35, 0x97, 0xe3, 0x70, 0x2a, 0xa3, 0xd8, 0x6d, 0xb4, 0xad, 0x4e,
	0x83, 0x6c, 0x60, 0xd8, 0x83, 0x2a, 0x0b, 0xef, 0xdd, 0xa6, 0x5e, 0x71, 0xdb, 0x54, 0xff, 0xdf,
	0xfc, 0xba, 0x41, 0x78, 0x1f, 0x84, 0x32, 0xc9, 0x88, 0x0a, 0xf6, 0x4e, 0xa1, 0x51, 0x00, 0xea,
	0xda, 0x2c, 0x59, 0x56, 0x5c, 0x9b, 0x25, 0xd3, 0xf3, 0xbb, 0xa7, 0xab, 0x94, 0xe5, 0x83, 0x35,
	0xc6, 0xfb, 0xca, 0x8f, 0x96, 0xff, 0xaf, 0x05, 0x47, 0x65, 0x6e, 0x75, 0x4d, 0x5e, 0x41, 0x73,
	0x51, 0x0c, 0x2a, 0x67, 0x79, 0x04, 0xf0, 0x14, 0x6a, 0x93, 0x28, 0x91, 0xc5, 0x15, 0xdc, 0xa1,
	0x2f, 0x5e, 0x65, 0x5d, 0x1d, 0x62, 0xf4, 0x99, 0x70, 0xef, 0x0d, 0xd8, 0x71, 0x94, 0xe8, 0xbd,
	0x0f, 0x26, 0xc5, 0xde, 0x07, 0x13, 0xb5, 0x45, 0x85, 0xe7, 0xd2, 0xf4, 0xd9, 0x9b, 0x01, 0x3c,
	0x12, 0xec, 0xe8, 0xe7, 0xa4, 0xdc, 0x4f, 0xab, 0xe7, 0xed, 0xd6, 0xa0, 0xa8, 0xca, 0xbd, 0xfe,
	0x0e, 0x8d, 0xe2, 0xa5, 0xe0, 0x5b, 0xa8, 0x27, 0x66, 0x90, 0x9a, 0xb7, 0xd5, 0x7b, 0xf9, 0x85,
	0x39, 0x93, 0x22, 0x0e, 0xbf, 0x83, 0x5a, 0xa2, 0x78, 0xf3, 0xa2, 0xcf, 0x77, 0x16, 0x25, 0x26,
	0xc6, 0x7f, 0x99, 0xbf, 0xe9, 0xa9, 0xa4, 0x32, 0x15, 0x97, 0x34, 0xce, 0xe9, 0xde, 0xfc, 0x69,
	0x41, 0x3d, 0x7f, 0x21, 0xd8, 0x00, 0x7b, 0x34, 0x1e, 0x05, 0xce, 0x1e, 0xb6, 0xa0, 0xfe, 0x71,
	0x38, 0x3e, 0xbb, 0x08, 0xfa, 0x8e, 0xa5, 0xe0, 0xf3, 0xab, 0xe1, 0xd0, 0xa9, 0xe0, 0x01, 0x34,
	0xa7, 0xbf, 0x5e, 0x5e, 0x06, 0x33, 0x32, 0x38, 0x73, 0xaa, 0x78, 0x08, 0x40, 0x82, 0xa9, 0x32,
	0x66, 0x41, 0xdf, 0xb1, 0xf1, 0x18, 0x8e, 0x26, 0x63, 0x32, 0xbb, 0x2e, 0x81, 0x35, 0xf4, 0xe0,
	0xc5, 0x43, 0xce, 0xf5, 0x55, 0x7f, 0x72, 0x7d, 0x3e, 0x20, 0xc1, 0xe7, 0x0f, 0xc3, 0xa1, 0xb3,
	0xaf, 0xca, 0x5c, 0x8d, 0x2e, 0x46, 0xe3, 0xcf, 0x23, 0xa7, 0xde, 0xfb, 0xc7, 0x86, 0x9a, 0xd6,
	0x88, 0xdf, 0x83, 0x3d, 0xe1, 0xe1, 0x2d, 0x3e, 0x35, 0x2d, 0xa9, 0x73, 0x2e, 0xd7, 0x3b, 0x2a,
	0x43, 0xaa, 0xb1, 0x3d, 0x7c, 0x0b, 0xb6, 0x1e, 0xe1, 0x8b, 0xd2, 0x00, 0x4a, 0x9f, 0xb3, 0x22,
	0xc5, 0x40, 0x26, 0x25, 0x80, 0xe6, 0xc3, 0xa7, 0x08, 0xbf, 0x2e, 0xe5, 0x6d, 0x7f, 0xf2, 0xbc,
	0xaf, 0x76, 0x3b, 0x0d, 0xcd, 0x4f, 0x50, 0xd3, 0x93, 0xc6, 0x2f, 0x2d, 0xcb, 0xdb, 0xbd, 0x14,
	0x7f, 0x0f, 0xdf, 0x81, 0xad, 0xdf, 0x56, 0x1e, 0xa0, 0xce, 0xe6, 0x41, 0x9b, 0xbc, 0xe3, 0x6d,
	0xd8, 0x64, 0x9d, 0x43, 0x4b, 0x99, 0xc2, 0x6c, 0x71, 0x43, 0xf9, 0xf6, 0x62, 0xbd, 0x67, 0x05,
	0xc5, 0x03, 0xae, 0x39, 0x3a, 0xd6, 0x89, 0x85, 0xbf, 0x18, 0x9e, 0x3e, 0x93, 0x94, 0xaf, 0x44,
	0x21, 0x5f, 0x41, 0x26, 0x7c, 0x4b, 0x7e, 0xd9, 0x61, 0x84, 0xbc, 0x37, 0x57, 0x77, 0x18, 0xdd,
	0x0a, 0x2c, 0x05, 0x29, 0x7b, 0xab, 0x85, 0x02, 0x3e, 0xbb, 0x4b, 0xc3, 0xa5, 0xbf, 0x77, 0x62,
	0xe1, 0xcf, 0xd0, 0xea, 0x73, 0xb1, 0x88, 0xee, 0x59, 0xf2, 0x29, 0x9d, 0xa3, 0x6b, 0xe2, 0x4a,
	0x50, 0xc1, 0xe0, 0x18, 0x4f, 0xb0, 0x8e, 0x65, 0x96, 0x17, 0x9e, 0xef, 0xeb, 0xbf, 0xad, 0x1f,
	0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x57, 0xa6, 0x17, 0xf0, 0x06, 0x00, 0x00,
}
