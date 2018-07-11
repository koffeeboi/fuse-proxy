// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/youtubeproxy/youtube.proto

package youtube

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Status tells whether or not the music file exists
type YtMusicReply_Status int32

const (
	YtMusicReply_EXIST     YtMusicReply_Status = 0
	YtMusicReply_NOT_EXIST YtMusicReply_Status = 1
)

var YtMusicReply_Status_name = map[int32]string{
	0: "EXIST",
	1: "NOT_EXIST",
}
var YtMusicReply_Status_value = map[string]int32{
	"EXIST":     0,
	"NOT_EXIST": 1,
}

func (x YtMusicReply_Status) String() string {
	return proto.EnumName(YtMusicReply_Status_name, int32(x))
}
func (YtMusicReply_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_youtube_c11f11c41a29fd09, []int{1, 0}
}

type YtMusicRequest struct {
	// Url is the url of the youtube music video
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YtMusicRequest) Reset()         { *m = YtMusicRequest{} }
func (m *YtMusicRequest) String() string { return proto.CompactTextString(m) }
func (*YtMusicRequest) ProtoMessage()    {}
func (*YtMusicRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_youtube_c11f11c41a29fd09, []int{0}
}
func (m *YtMusicRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YtMusicRequest.Unmarshal(m, b)
}
func (m *YtMusicRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YtMusicRequest.Marshal(b, m, deterministic)
}
func (dst *YtMusicRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YtMusicRequest.Merge(dst, src)
}
func (m *YtMusicRequest) XXX_Size() int {
	return xxx_messageInfo_YtMusicRequest.Size(m)
}
func (m *YtMusicRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_YtMusicRequest.DiscardUnknown(m)
}

var xxx_messageInfo_YtMusicRequest proto.InternalMessageInfo

func (m *YtMusicRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type YtMusicReply struct {
	// Name is the name of the music file
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Status               YtMusicReply_Status `protobuf:"varint,2,opt,name=status,proto3,enum=youtube.YtMusicReply_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *YtMusicReply) Reset()         { *m = YtMusicReply{} }
func (m *YtMusicReply) String() string { return proto.CompactTextString(m) }
func (*YtMusicReply) ProtoMessage()    {}
func (*YtMusicReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_youtube_c11f11c41a29fd09, []int{1}
}
func (m *YtMusicReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YtMusicReply.Unmarshal(m, b)
}
func (m *YtMusicReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YtMusicReply.Marshal(b, m, deterministic)
}
func (dst *YtMusicReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YtMusicReply.Merge(dst, src)
}
func (m *YtMusicReply) XXX_Size() int {
	return xxx_messageInfo_YtMusicReply.Size(m)
}
func (m *YtMusicReply) XXX_DiscardUnknown() {
	xxx_messageInfo_YtMusicReply.DiscardUnknown(m)
}

var xxx_messageInfo_YtMusicReply proto.InternalMessageInfo

func (m *YtMusicReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *YtMusicReply) GetStatus() YtMusicReply_Status {
	if m != nil {
		return m.Status
	}
	return YtMusicReply_EXIST
}

func init() {
	proto.RegisterType((*YtMusicRequest)(nil), "youtube.YtMusicRequest")
	proto.RegisterType((*YtMusicReply)(nil), "youtube.YtMusicReply")
	proto.RegisterEnum("youtube.YtMusicReply_Status", YtMusicReply_Status_name, YtMusicReply_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// YoutubeServiceClient is the client API for YoutubeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type YoutubeServiceClient interface {
	FindYoutubeMusic(ctx context.Context, in *YtMusicRequest, opts ...grpc.CallOption) (*YtMusicReply, error)
}

type youtubeServiceClient struct {
	cc *grpc.ClientConn
}

func NewYoutubeServiceClient(cc *grpc.ClientConn) YoutubeServiceClient {
	return &youtubeServiceClient{cc}
}

func (c *youtubeServiceClient) FindYoutubeMusic(ctx context.Context, in *YtMusicRequest, opts ...grpc.CallOption) (*YtMusicReply, error) {
	out := new(YtMusicReply)
	err := c.cc.Invoke(ctx, "/youtube.YoutubeService/FindYoutubeMusic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YoutubeServiceServer is the server API for YoutubeService service.
type YoutubeServiceServer interface {
	FindYoutubeMusic(context.Context, *YtMusicRequest) (*YtMusicReply, error)
}

func RegisterYoutubeServiceServer(s *grpc.Server, srv YoutubeServiceServer) {
	s.RegisterService(&_YoutubeService_serviceDesc, srv)
}

func _YoutubeService_FindYoutubeMusic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(YtMusicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YoutubeServiceServer).FindYoutubeMusic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/youtube.YoutubeService/FindYoutubeMusic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YoutubeServiceServer).FindYoutubeMusic(ctx, req.(*YtMusicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _YoutubeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "youtube.YoutubeService",
	HandlerType: (*YoutubeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindYoutubeMusic",
			Handler:    _YoutubeService_FindYoutubeMusic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/youtubeproxy/youtube.proto",
}

func init() {
	proto.RegisterFile("proto/youtubeproxy/youtube.proto", fileDescriptor_youtube_c11f11c41a29fd09)
}

var fileDescriptor_youtube_c11f11c41a29fd09 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xaf, 0xcc, 0x2f, 0x2d, 0x29, 0x4d, 0x4a, 0x2d, 0x28, 0xca, 0xaf, 0xa8, 0x84, 0x71,
	0xf4, 0xc0, 0x52, 0x42, 0xec, 0x50, 0xae, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x7e,
	0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31, 0x44,
	0x99, 0x92, 0x12, 0x17, 0x5f, 0x64, 0x89, 0x6f, 0x69, 0x71, 0x66, 0x72, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x90, 0x00, 0x17, 0x73, 0x69, 0x51, 0x8e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x88, 0xa9, 0x54, 0xc3, 0xc5, 0x03, 0x57, 0x53, 0x90, 0x53, 0x29, 0x24, 0xc4, 0xc5, 0x92,
	0x97, 0x98, 0x9b, 0x0a, 0x55, 0x02, 0x66, 0x0b, 0x99, 0x70, 0xb1, 0x15, 0x97, 0x24, 0x96, 0x94,
	0x16, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0xf0, 0x19, 0xc9, 0xe8, 0xc1, 0x9c, 0x83, 0xac, 0x55, 0x2f,
	0x18, 0xac, 0x26, 0x08, 0xaa, 0x56, 0x49, 0x89, 0x8b, 0x0d, 0x22, 0x22, 0xc4, 0xc9, 0xc5, 0xea,
	0x1a, 0xe1, 0x19, 0x1c, 0x22, 0xc0, 0x20, 0xc4, 0xcb, 0xc5, 0xe9, 0xe7, 0x1f, 0x12, 0x0f, 0xe1,
	0x32, 0x1a, 0x15, 0x73, 0xf1, 0x45, 0x42, 0x8c, 0x0a, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15,
	0x4a, 0xe4, 0x12, 0x70, 0xcb, 0xcc, 0x4b, 0x81, 0x8a, 0x82, 0x4d, 0x17, 0x12, 0xc7, 0xb4, 0x0f,
	0xec, 0x1d, 0x29, 0x51, 0xac, 0x0e, 0x51, 0x92, 0x6d, 0xba, 0xfc, 0x64, 0x32, 0x93, 0xb8, 0x92,
	0x10, 0x2c, 0xd8, 0xf4, 0x73, 0x41, 0x92, 0xfa, 0xe9, 0xa9, 0x25, 0x56, 0x8c, 0x5a, 0x49, 0x6c,
	0xe0, 0xd0, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x47, 0x86, 0x3c, 0xe6, 0x68, 0x01, 0x00,
	0x00,
}