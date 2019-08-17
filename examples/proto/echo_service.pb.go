// Code generated by protoc-gen-go. DO NOT EDIT.
// source: examples/proto/echo_service.proto

package proto

import (
	context "context"
	fmt "fmt"
	_ "github.com/binchencoder/ease-gateway/gateway/options"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Embedded struct {
	// Types that are valid to be assigned to Mark:
	//	*Embedded_Progress
	//	*Embedded_Note
	Mark                 isEmbedded_Mark `protobuf_oneof:"mark"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Embedded) Reset()         { *m = Embedded{} }
func (m *Embedded) String() string { return proto.CompactTextString(m) }
func (*Embedded) ProtoMessage()    {}
func (*Embedded) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d3f52be8fcd98, []int{0}
}

func (m *Embedded) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Embedded.Unmarshal(m, b)
}
func (m *Embedded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Embedded.Marshal(b, m, deterministic)
}
func (m *Embedded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Embedded.Merge(m, src)
}
func (m *Embedded) XXX_Size() int {
	return xxx_messageInfo_Embedded.Size(m)
}
func (m *Embedded) XXX_DiscardUnknown() {
	xxx_messageInfo_Embedded.DiscardUnknown(m)
}

var xxx_messageInfo_Embedded proto.InternalMessageInfo

type isEmbedded_Mark interface {
	isEmbedded_Mark()
}

type Embedded_Progress struct {
	Progress int64 `protobuf:"varint,1,opt,name=progress,proto3,oneof"`
}

type Embedded_Note struct {
	Note string `protobuf:"bytes,2,opt,name=note,proto3,oneof"`
}

func (*Embedded_Progress) isEmbedded_Mark() {}

func (*Embedded_Note) isEmbedded_Mark() {}

func (m *Embedded) GetMark() isEmbedded_Mark {
	if m != nil {
		return m.Mark
	}
	return nil
}

func (m *Embedded) GetProgress() int64 {
	if x, ok := m.GetMark().(*Embedded_Progress); ok {
		return x.Progress
	}
	return 0
}

func (m *Embedded) GetNote() string {
	if x, ok := m.GetMark().(*Embedded_Note); ok {
		return x.Note
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Embedded) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Embedded_Progress)(nil),
		(*Embedded_Note)(nil),
	}
}

type SimpleMessage struct {
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num int64  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	// Types that are valid to be assigned to Code:
	//	*SimpleMessage_LineNum
	//	*SimpleMessage_Lang
	Code   isSimpleMessage_Code `protobuf_oneof:"code"`
	Status *Embedded            `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are valid to be assigned to Ext:
	//	*SimpleMessage_En
	//	*SimpleMessage_No
	Ext                  isSimpleMessage_Ext `protobuf_oneof:"ext"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SimpleMessage) Reset()         { *m = SimpleMessage{} }
func (m *SimpleMessage) String() string { return proto.CompactTextString(m) }
func (*SimpleMessage) ProtoMessage()    {}
func (*SimpleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b4d3f52be8fcd98, []int{1}
}

func (m *SimpleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleMessage.Unmarshal(m, b)
}
func (m *SimpleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleMessage.Marshal(b, m, deterministic)
}
func (m *SimpleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleMessage.Merge(m, src)
}
func (m *SimpleMessage) XXX_Size() int {
	return xxx_messageInfo_SimpleMessage.Size(m)
}
func (m *SimpleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleMessage proto.InternalMessageInfo

func (m *SimpleMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SimpleMessage) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type isSimpleMessage_Code interface {
	isSimpleMessage_Code()
}

type SimpleMessage_LineNum struct {
	LineNum int64 `protobuf:"varint,3,opt,name=line_num,json=lineNum,proto3,oneof"`
}

type SimpleMessage_Lang struct {
	Lang string `protobuf:"bytes,4,opt,name=lang,proto3,oneof"`
}

func (*SimpleMessage_LineNum) isSimpleMessage_Code() {}

func (*SimpleMessage_Lang) isSimpleMessage_Code() {}

func (m *SimpleMessage) GetCode() isSimpleMessage_Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *SimpleMessage) GetLineNum() int64 {
	if x, ok := m.GetCode().(*SimpleMessage_LineNum); ok {
		return x.LineNum
	}
	return 0
}

func (m *SimpleMessage) GetLang() string {
	if x, ok := m.GetCode().(*SimpleMessage_Lang); ok {
		return x.Lang
	}
	return ""
}

func (m *SimpleMessage) GetStatus() *Embedded {
	if m != nil {
		return m.Status
	}
	return nil
}

type isSimpleMessage_Ext interface {
	isSimpleMessage_Ext()
}

type SimpleMessage_En struct {
	En int64 `protobuf:"varint,6,opt,name=en,proto3,oneof"`
}

type SimpleMessage_No struct {
	No *Embedded `protobuf:"bytes,7,opt,name=no,proto3,oneof"`
}

func (*SimpleMessage_En) isSimpleMessage_Ext() {}

func (*SimpleMessage_No) isSimpleMessage_Ext() {}

func (m *SimpleMessage) GetExt() isSimpleMessage_Ext {
	if m != nil {
		return m.Ext
	}
	return nil
}

func (m *SimpleMessage) GetEn() int64 {
	if x, ok := m.GetExt().(*SimpleMessage_En); ok {
		return x.En
	}
	return 0
}

func (m *SimpleMessage) GetNo() *Embedded {
	if x, ok := m.GetExt().(*SimpleMessage_No); ok {
		return x.No
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SimpleMessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SimpleMessage_LineNum)(nil),
		(*SimpleMessage_Lang)(nil),
		(*SimpleMessage_En)(nil),
		(*SimpleMessage_No)(nil),
	}
}

func init() {
	proto.RegisterType((*Embedded)(nil), "ease.gateway.examples.proto.Embedded")
	proto.RegisterType((*SimpleMessage)(nil), "ease.gateway.examples.proto.SimpleMessage")
}

func init() { proto.RegisterFile("examples/proto/echo_service.proto", fileDescriptor_4b4d3f52be8fcd98) }

var fileDescriptor_4b4d3f52be8fcd98 = []byte{
	// 489 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xee, 0x6c, 0xb6, 0x49, 0x3a, 0x45, 0x29, 0x83, 0xd2, 0x35, 0xb1, 0x92, 0x06, 0x85, 0x90,
	0xc3, 0x0e, 0x89, 0x01, 0x41, 0xf0, 0x12, 0xac, 0xf4, 0xa2, 0x87, 0xed, 0x2d, 0x97, 0x30, 0xc9,
	0x3c, 0xb7, 0x8b, 0xbb, 0x33, 0xcb, 0xce, 0x6c, 0x6d, 0x58, 0x72, 0xd1, 0xbf, 0xa0, 0x7f, 0xc3,
	0x8b, 0xe0, 0x8f, 0x08, 0xfe, 0x12, 0xaf, 0x39, 0x4b, 0x65, 0x66, 0x37, 0x82, 0xb6, 0x04, 0x7a,
	0xe8, 0x69, 0xe6, 0x7d, 0xef, 0x7b, 0xef, 0x7b, 0xef, 0x9b, 0xc1, 0xc7, 0x70, 0xc9, 0x92, 0x34,
	0x06, 0x45, 0xd3, 0x4c, 0x6a, 0x49, 0x61, 0x7e, 0x2e, 0xa7, 0x0a, 0xb2, 0x8b, 0x68, 0x0e, 0xbe,
	0x85, 0x48, 0x1b, 0x98, 0x02, 0x3f, 0x64, 0x1a, 0x3e, 0xb2, 0x85, 0xbf, 0xe1, 0x97, 0xc9, 0xd6,
	0x71, 0x85, 0x53, 0x99, 0xea, 0x48, 0x0a, 0x45, 0x99, 0x10, 0x52, 0x33, 0x7b, 0x2f, 0x29, 0xdd,
	0x37, 0xb8, 0x79, 0x92, 0xcc, 0x80, 0x73, 0xe0, 0xe4, 0x31, 0x6e, 0xa6, 0x99, 0x0c, 0x33, 0x50,
	0xca, 0x43, 0x1d, 0xd4, 0xab, 0x9d, 0xee, 0x04, 0x7f, 0x11, 0xf2, 0x00, 0xbb, 0x42, 0x6a, 0xf0,
	0x9c, 0x0e, 0xea, 0xed, 0x9d, 0xee, 0x04, 0x36, 0x1a, 0xd7, 0xb1, 0x9b, 0xb0, 0xec, 0x43, 0xf7,
	0x37, 0xc2, 0xf7, 0xce, 0x22, 0x23, 0xfe, 0x16, 0x94, 0x62, 0x21, 0x90, 0xfb, 0xd8, 0x89, 0xb8,
	0xed, 0xb3, 0x17, 0x38, 0x11, 0x27, 0x07, 0xb8, 0x26, 0xf2, 0xc4, 0x96, 0xd7, 0x02, 0x73, 0x25,
	0x6d, 0xdc, 0x8c, 0x23, 0x01, 0x53, 0x03, 0xd7, 0x2a, 0xbd, 0x86, 0x41, 0xde, 0xe5, 0x89, 0x91,
	0x8b, 0x99, 0x08, 0x3d, 0x77, 0x23, 0x67, 0x22, 0xf2, 0x0a, 0xd7, 0x95, 0x66, 0x3a, 0x57, 0xde,
	0x6e, 0x07, 0xf5, 0xf6, 0x87, 0xcf, 0xfc, 0x2d, 0xfb, 0xfb, 0x9b, 0xcd, 0x82, 0xaa, 0x88, 0x1c,
	0x60, 0x07, 0x84, 0x57, 0xb7, 0x5a, 0x28, 0x70, 0x40, 0x90, 0x17, 0xd8, 0x11, 0xd2, 0x6b, 0xdc,
	0xa2, 0x99, 0x29, 0x14, 0xd2, 0x2c, 0x3e, 0x97, 0x1c, 0xc6, 0xbb, 0xb8, 0x06, 0x97, 0x7a, 0xf8,
	0xd3, 0xc5, 0xfb, 0x27, 0xf3, 0x73, 0x79, 0x56, 0xbe, 0x0e, 0xf9, 0xe2, 0x60, 0xd7, 0xc4, 0xa4,
	0xbf, 0xb5, 0xe9, 0x3f, 0x96, 0xb5, 0x6e, 0xc1, 0xed, 0x7e, 0x47, 0xab, 0xf5, 0xe8, 0x1b, 0xea,
	0x3e, 0xa4, 0x17, 0x03, 0x5a, 0x31, 0xed, 0x07, 0xa1, 0x45, 0xc4, 0x97, 0x93, 0x23, 0xd2, 0xbe,
	0x31, 0x41, 0x0b, 0x91, 0x27, 0xcb, 0xc9, 0x53, 0xd2, 0xdd, 0x92, 0xa6, 0x85, 0x31, 0x7b, 0x39,
	0x19, 0x10, 0xfa, 0x3f, 0x6b, 0x50, 0xd1, 0x36, 0x0f, 0xb7, 0xa4, 0x45, 0xe9, 0xad, 0x6f, 0xbe,
	0xc3, 0x8d, 0xba, 0x43, 0x5a, 0x08, 0x59, 0xa6, 0xc9, 0x27, 0x84, 0x9b, 0xc6, 0x96, 0xb1, 0xe4,
	0x8b, 0x3b, 0xb3, 0xe6, 0xc9, 0x6a, 0x3d, 0x6a, 0x5d, 0x37, 0x66, 0x3a, 0x93, 0x7c, 0xf1, 0x12,
	0xf5, 0xc9, 0x67, 0x84, 0xb1, 0x19, 0xe2, 0x35, 0xc4, 0xa0, 0xe1, 0xce, 0xc6, 0x38, 0x5a, 0xad,
	0x47, 0x8f, 0xfa, 0x87, 0xd7, 0xc6, 0xe0, 0x56, 0xb6, 0xd5, 0xfe, 0xb5, 0x1e, 0x1d, 0x36, 0x7f,
	0x7c, 0xbd, 0xba, 0x6a, 0x10, 0x37, 0xcc, 0xd2, 0x79, 0xab, 0xc1, 0xe1, 0x3d, 0xcb, 0x63, 0xdd,
	0x41, 0xe3, 0xc6, 0x64, 0xd7, 0xb6, 0x9c, 0xd5, 0xed, 0xf1, 0xfc, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x35, 0x06, 0x8b, 0xf6, 0x0b, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServiceClient interface {
	Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
	EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/ease.gateway.examples.proto.EchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoBody(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/ease.gateway.examples.proto.EchoService/EchoBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/ease.gateway.examples.proto.EchoService/EchoDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
type EchoServiceServer interface {
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
	EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ease.gateway.examples.proto.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ease.gateway.examples.proto.EchoService/EchoBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoBody(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_EchoDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ease.gateway.examples.proto.EchoService/EchoDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoDelete(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ease.gateway.examples.proto.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "EchoBody",
			Handler:    _EchoService_EchoBody_Handler,
		},
		{
			MethodName: "EchoDelete",
			Handler:    _EchoService_EchoDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/proto/echo_service.proto",
}
