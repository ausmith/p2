// Code generated by protoc-gen-go.
// source: pkg/grpc/labelstore/protos/label_store.proto
// DO NOT EDIT!

/*
Package label_store_protos is a generated protocol buffer package.

It is generated from these files:
	pkg/grpc/labelstore/protos/label_store.proto

It has these top-level messages:
	WatchMatchesRequest
	Labeled
	WatchMatchesResponse
*/
package label_store_protos

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

type LabelType int32

const (
	LabelType_unknown                LabelType = 0
	LabelType_pod                    LabelType = 1
	LabelType_node                   LabelType = 2
	LabelType_pod_clusters           LabelType = 3
	LabelType_replication_controller LabelType = 4
	LabelType_rolls                  LabelType = 5
)

var LabelType_name = map[int32]string{
	0: "unknown",
	1: "pod",
	2: "node",
	3: "pod_clusters",
	4: "replication_controller",
	5: "rolls",
}
var LabelType_value = map[string]int32{
	"unknown":                0,
	"pod":                    1,
	"node":                   2,
	"pod_clusters":           3,
	"replication_controller": 4,
	"rolls":                  5,
}

func (x LabelType) String() string {
	return proto.EnumName(LabelType_name, int32(x))
}
func (LabelType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type WatchMatchesRequest struct {
	Selector  string    `protobuf:"bytes,1,opt,name=selector" json:"selector,omitempty"`
	LabelType LabelType `protobuf:"varint,2,opt,name=label_type,json=labelType,enum=label_store_protos.LabelType" json:"label_type,omitempty"`
}

func (m *WatchMatchesRequest) Reset()                    { *m = WatchMatchesRequest{} }
func (m *WatchMatchesRequest) String() string            { return proto.CompactTextString(m) }
func (*WatchMatchesRequest) ProtoMessage()               {}
func (*WatchMatchesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *WatchMatchesRequest) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

func (m *WatchMatchesRequest) GetLabelType() LabelType {
	if m != nil {
		return m.LabelType
	}
	return LabelType_unknown
}

type Labeled struct {
	LabelType LabelType         `protobuf:"varint,1,opt,name=label_type,json=labelType,enum=label_store_protos.LabelType" json:"label_type,omitempty"`
	Id        string            `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Labels    map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Labeled) Reset()                    { *m = Labeled{} }
func (m *Labeled) String() string            { return proto.CompactTextString(m) }
func (*Labeled) ProtoMessage()               {}
func (*Labeled) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Labeled) GetLabelType() LabelType {
	if m != nil {
		return m.LabelType
	}
	return LabelType_unknown
}

func (m *Labeled) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Labeled) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type WatchMatchesResponse struct {
	Labeled []*Labeled `protobuf:"bytes,1,rep,name=labeled" json:"labeled,omitempty"`
}

func (m *WatchMatchesResponse) Reset()                    { *m = WatchMatchesResponse{} }
func (m *WatchMatchesResponse) String() string            { return proto.CompactTextString(m) }
func (*WatchMatchesResponse) ProtoMessage()               {}
func (*WatchMatchesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WatchMatchesResponse) GetLabeled() []*Labeled {
	if m != nil {
		return m.Labeled
	}
	return nil
}

func init() {
	proto.RegisterType((*WatchMatchesRequest)(nil), "label_store_protos.WatchMatchesRequest")
	proto.RegisterType((*Labeled)(nil), "label_store_protos.Labeled")
	proto.RegisterType((*WatchMatchesResponse)(nil), "label_store_protos.WatchMatchesResponse")
	proto.RegisterEnum("label_store_protos.LabelType", LabelType_name, LabelType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for P2LabelStore service

type P2LabelStoreClient interface {
	WatchMatches(ctx context.Context, in *WatchMatchesRequest, opts ...grpc.CallOption) (P2LabelStore_WatchMatchesClient, error)
}

type p2LabelStoreClient struct {
	cc *grpc.ClientConn
}

func NewP2LabelStoreClient(cc *grpc.ClientConn) P2LabelStoreClient {
	return &p2LabelStoreClient{cc}
}

func (c *p2LabelStoreClient) WatchMatches(ctx context.Context, in *WatchMatchesRequest, opts ...grpc.CallOption) (P2LabelStore_WatchMatchesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_P2LabelStore_serviceDesc.Streams[0], c.cc, "/label_store_protos.P2LabelStore/WatchMatches", opts...)
	if err != nil {
		return nil, err
	}
	x := &p2LabelStoreWatchMatchesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type P2LabelStore_WatchMatchesClient interface {
	Recv() (*WatchMatchesResponse, error)
	grpc.ClientStream
}

type p2LabelStoreWatchMatchesClient struct {
	grpc.ClientStream
}

func (x *p2LabelStoreWatchMatchesClient) Recv() (*WatchMatchesResponse, error) {
	m := new(WatchMatchesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for P2LabelStore service

type P2LabelStoreServer interface {
	WatchMatches(*WatchMatchesRequest, P2LabelStore_WatchMatchesServer) error
}

func RegisterP2LabelStoreServer(s *grpc.Server, srv P2LabelStoreServer) {
	s.RegisterService(&_P2LabelStore_serviceDesc, srv)
}

func _P2LabelStore_WatchMatches_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchMatchesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(P2LabelStoreServer).WatchMatches(m, &p2LabelStoreWatchMatchesServer{stream})
}

type P2LabelStore_WatchMatchesServer interface {
	Send(*WatchMatchesResponse) error
	grpc.ServerStream
}

type p2LabelStoreWatchMatchesServer struct {
	grpc.ServerStream
}

func (x *p2LabelStoreWatchMatchesServer) Send(m *WatchMatchesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _P2LabelStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "label_store_protos.P2LabelStore",
	HandlerType: (*P2LabelStoreServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchMatches",
			Handler:       _P2LabelStore_WatchMatches_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/grpc/labelstore/protos/label_store.proto",
}

func init() { proto.RegisterFile("pkg/grpc/labelstore/protos/label_store.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x52, 0x51, 0x8b, 0xd4, 0x30,
	0x10, 0xbe, 0xb4, 0xb7, 0xd7, 0xeb, 0xec, 0x72, 0x84, 0xf1, 0x90, 0x52, 0x11, 0x96, 0xbe, 0x58,
	0x44, 0x76, 0xa5, 0x22, 0xa8, 0x08, 0x3e, 0xf9, 0xe6, 0x81, 0x54, 0xc1, 0xc7, 0xd2, 0x6b, 0x86,
	0xb3, 0x6c, 0x48, 0x62, 0x92, 0x2a, 0xfd, 0x9d, 0xfe, 0x21, 0x69, 0xb3, 0x77, 0xac, 0xba, 0x2a,
	0xbe, 0xb4, 0x33, 0x5f, 0xbe, 0x7c, 0xdf, 0xcc, 0x47, 0xe0, 0x89, 0xd9, 0xdd, 0x6c, 0x6f, 0xac,
	0xe9, 0xb6, 0xb2, 0xbd, 0x26, 0xe9, 0xbc, 0xb6, 0xb4, 0x35, 0x56, 0x7b, 0xed, 0x02, 0xd2, 0xcc,
	0xd0, 0x66, 0x86, 0x10, 0x0f, 0xa0, 0x26, 0xb0, 0x0a, 0x0d, 0xf7, 0x3e, 0xb5, 0xbe, 0xfb, 0x7c,
	0x35, 0x7d, 0xc8, 0xd5, 0xf4, 0x65, 0x20, 0xe7, 0x31, 0x87, 0x73, 0x47, 0x92, 0x3a, 0xaf, 0x6d,
	0xc6, 0xd6, 0xac, 0x4c, 0xeb, 0xbb, 0x1e, 0x5f, 0x03, 0x04, 0x21, 0x3f, 0x1a, 0xca, 0xa2, 0x35,
	0x2b, 0x2f, 0xaa, 0x87, 0x9b, 0xdf, 0xb5, 0x37, 0xef, 0x26, 0xe8, 0xe3, 0x68, 0xa8, 0x4e, 0xe5,
	0x6d, 0x59, 0x7c, 0x67, 0x90, 0xcc, 0x07, 0x24, 0x7e, 0x51, 0x62, 0xff, 0xa7, 0x84, 0x17, 0x10,
	0xf5, 0x62, 0xf6, 0x4f, 0xeb, 0xa8, 0x17, 0xf8, 0x06, 0xce, 0x42, 0x0a, 0x59, 0xbc, 0x8e, 0xcb,
	0x65, 0xf5, 0xe8, 0x8f, 0x4a, 0x24, 0xc2, 0xdf, 0xbd, 0x55, 0xde, 0x8e, 0xf5, 0xfe, 0x5a, 0xfe,
	0x12, 0x96, 0x07, 0x30, 0x72, 0x88, 0x77, 0x34, 0xee, 0xd7, 0x9f, 0x4a, 0xbc, 0x84, 0xc5, 0xd7,
	0x56, 0x0e, 0xb4, 0x37, 0x0d, 0xcd, 0xab, 0xe8, 0x05, 0x2b, 0xae, 0xe0, 0xf2, 0xe7, 0x18, 0x9d,
	0xd1, 0xca, 0x11, 0x3e, 0x87, 0x44, 0x06, 0xc7, 0x8c, 0xcd, 0x43, 0x3d, 0xf8, 0xcb, 0x50, 0xf5,
	0x2d, 0xf7, 0xb1, 0x80, 0xf4, 0x6e, 0x65, 0x5c, 0x42, 0x32, 0xa8, 0x9d, 0xd2, 0xdf, 0x14, 0x3f,
	0xc1, 0x04, 0x62, 0xa3, 0x05, 0x67, 0x78, 0x0e, 0xa7, 0x4a, 0x0b, 0xe2, 0x11, 0x72, 0x58, 0x19,
	0x2d, 0x9a, 0x4e, 0x0e, 0xce, 0x93, 0x75, 0x3c, 0xc6, 0x1c, 0xee, 0x5b, 0x32, 0xb2, 0xef, 0x5a,
	0xdf, 0x6b, 0xd5, 0x74, 0x5a, 0x79, 0xab, 0xa5, 0x24, 0xcb, 0x4f, 0x31, 0x85, 0xc5, 0x54, 0x3b,
	0xbe, 0xa8, 0x06, 0x58, 0xbd, 0xaf, 0x66, 0x9f, 0x0f, 0xd3, 0x34, 0x48, 0xb0, 0x3a, 0x5c, 0x02,
	0x8f, 0x06, 0x78, 0xe4, 0xb5, 0xe4, 0xe5, 0xbf, 0x89, 0x21, 0x8f, 0xe2, 0xe4, 0x29, 0xbb, 0x3e,
	0x9b, 0x09, 0xcf, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0xad, 0x5e, 0xbd, 0xbd, 0x02, 0x00,
	0x00,
}
