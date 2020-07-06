// Code generated by protoc-gen-go. DO NOT EDIT.
// source: httpoptions/http.proto

package ease_api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Http struct {
	Rules                        []*HttpRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	FullyDecodeReservedExpansion bool        `protobuf:"varint,2,opt,name=fully_decode_reserved_expansion,json=fullyDecodeReservedExpansion,proto3" json:"fully_decode_reserved_expansion,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}    `json:"-"`
	XXX_unrecognized             []byte      `json:"-"`
	XXX_sizecache                int32       `json:"-"`
}

func (m *Http) Reset()         { *m = Http{} }
func (m *Http) String() string { return proto.CompactTextString(m) }
func (*Http) ProtoMessage()    {}
func (*Http) Descriptor() ([]byte, []int) {
	return fileDescriptor_855b53eccb99bc99, []int{0}
}

func (m *Http) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Http.Unmarshal(m, b)
}
func (m *Http) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Http.Marshal(b, m, deterministic)
}
func (m *Http) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Http.Merge(m, src)
}
func (m *Http) XXX_Size() int {
	return xxx_messageInfo_Http.Size(m)
}
func (m *Http) XXX_DiscardUnknown() {
	xxx_messageInfo_Http.DiscardUnknown(m)
}

var xxx_messageInfo_Http proto.InternalMessageInfo

func (m *Http) GetRules() []*HttpRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Http) GetFullyDecodeReservedExpansion() bool {
	if m != nil {
		return m.FullyDecodeReservedExpansion
	}
	return false
}

type HttpRule struct {
	Selector string `protobuf:"bytes,1,opt,name=selector,proto3" json:"selector,omitempty"`
	// Types that are valid to be assigned to Pattern:
	//	*HttpRule_Get
	//	*HttpRule_Put
	//	*HttpRule_Post
	//	*HttpRule_Delete
	//	*HttpRule_Patch
	//	*HttpRule_Custom
	Pattern              isHttpRule_Pattern `protobuf_oneof:"pattern"`
	Body                 string             `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	ResponseBody         string             `protobuf:"bytes,12,opt,name=response_body,json=responseBody,proto3" json:"response_body,omitempty"`
	AdditionalBindings   []*HttpRule        `protobuf:"bytes,11,rep,name=additional_bindings,json=additionalBindings,proto3" json:"additional_bindings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HttpRule) Reset()         { *m = HttpRule{} }
func (m *HttpRule) String() string { return proto.CompactTextString(m) }
func (*HttpRule) ProtoMessage()    {}
func (*HttpRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_855b53eccb99bc99, []int{1}
}

func (m *HttpRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRule.Unmarshal(m, b)
}
func (m *HttpRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRule.Marshal(b, m, deterministic)
}
func (m *HttpRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRule.Merge(m, src)
}
func (m *HttpRule) XXX_Size() int {
	return xxx_messageInfo_HttpRule.Size(m)
}
func (m *HttpRule) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRule.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRule proto.InternalMessageInfo

func (m *HttpRule) GetSelector() string {
	if m != nil {
		return m.Selector
	}
	return ""
}

type isHttpRule_Pattern interface {
	isHttpRule_Pattern()
}

type HttpRule_Get struct {
	Get string `protobuf:"bytes,2,opt,name=get,proto3,oneof"`
}

type HttpRule_Put struct {
	Put string `protobuf:"bytes,3,opt,name=put,proto3,oneof"`
}

type HttpRule_Post struct {
	Post string `protobuf:"bytes,4,opt,name=post,proto3,oneof"`
}

type HttpRule_Delete struct {
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3,oneof"`
}

type HttpRule_Patch struct {
	Patch string `protobuf:"bytes,6,opt,name=patch,proto3,oneof"`
}

type HttpRule_Custom struct {
	Custom *CustomHttpPattern `protobuf:"bytes,8,opt,name=custom,proto3,oneof"`
}

func (*HttpRule_Get) isHttpRule_Pattern() {}

func (*HttpRule_Put) isHttpRule_Pattern() {}

func (*HttpRule_Post) isHttpRule_Pattern() {}

func (*HttpRule_Delete) isHttpRule_Pattern() {}

func (*HttpRule_Patch) isHttpRule_Pattern() {}

func (*HttpRule_Custom) isHttpRule_Pattern() {}

func (m *HttpRule) GetPattern() isHttpRule_Pattern {
	if m != nil {
		return m.Pattern
	}
	return nil
}

func (m *HttpRule) GetGet() string {
	if x, ok := m.GetPattern().(*HttpRule_Get); ok {
		return x.Get
	}
	return ""
}

func (m *HttpRule) GetPut() string {
	if x, ok := m.GetPattern().(*HttpRule_Put); ok {
		return x.Put
	}
	return ""
}

func (m *HttpRule) GetPost() string {
	if x, ok := m.GetPattern().(*HttpRule_Post); ok {
		return x.Post
	}
	return ""
}

func (m *HttpRule) GetDelete() string {
	if x, ok := m.GetPattern().(*HttpRule_Delete); ok {
		return x.Delete
	}
	return ""
}

func (m *HttpRule) GetPatch() string {
	if x, ok := m.GetPattern().(*HttpRule_Patch); ok {
		return x.Patch
	}
	return ""
}

func (m *HttpRule) GetCustom() *CustomHttpPattern {
	if x, ok := m.GetPattern().(*HttpRule_Custom); ok {
		return x.Custom
	}
	return nil
}

func (m *HttpRule) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *HttpRule) GetResponseBody() string {
	if m != nil {
		return m.ResponseBody
	}
	return ""
}

func (m *HttpRule) GetAdditionalBindings() []*HttpRule {
	if m != nil {
		return m.AdditionalBindings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpRule_Get)(nil),
		(*HttpRule_Put)(nil),
		(*HttpRule_Post)(nil),
		(*HttpRule_Delete)(nil),
		(*HttpRule_Patch)(nil),
		(*HttpRule_Custom)(nil),
	}
}

type CustomHttpPattern struct {
	Kind                 string   `protobuf:"bytes,1,opt,name=kind,proto3" json:"kind,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomHttpPattern) Reset()         { *m = CustomHttpPattern{} }
func (m *CustomHttpPattern) String() string { return proto.CompactTextString(m) }
func (*CustomHttpPattern) ProtoMessage()    {}
func (*CustomHttpPattern) Descriptor() ([]byte, []int) {
	return fileDescriptor_855b53eccb99bc99, []int{2}
}

func (m *CustomHttpPattern) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomHttpPattern.Unmarshal(m, b)
}
func (m *CustomHttpPattern) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomHttpPattern.Marshal(b, m, deterministic)
}
func (m *CustomHttpPattern) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomHttpPattern.Merge(m, src)
}
func (m *CustomHttpPattern) XXX_Size() int {
	return xxx_messageInfo_CustomHttpPattern.Size(m)
}
func (m *CustomHttpPattern) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomHttpPattern.DiscardUnknown(m)
}

var xxx_messageInfo_CustomHttpPattern proto.InternalMessageInfo

func (m *CustomHttpPattern) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CustomHttpPattern) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*Http)(nil), "ease.api.Http")
	proto.RegisterType((*HttpRule)(nil), "ease.api.HttpRule")
	proto.RegisterType((*CustomHttpPattern)(nil), "ease.api.CustomHttpPattern")
}

func init() { proto.RegisterFile("httpoptions/http.proto", fileDescriptor_855b53eccb99bc99) }

var fileDescriptor_855b53eccb99bc99 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0x49, 0x9b, 0x66, 0x93, 0x69, 0x39, 0x60, 0xd0, 0xca, 0x02, 0x24, 0xa2, 0x70, 0xc9,
	0x29, 0x48, 0x8b, 0x38, 0x71, 0x22, 0x4b, 0xa5, 0x72, 0xab, 0xf2, 0x02, 0x91, 0x1b, 0x0f, 0x6d,
	0x44, 0x6a, 0x5b, 0xf1, 0x04, 0xe8, 0xeb, 0xf0, 0x50, 0x3c, 0x0f, 0xb2, 0x93, 0xb4, 0x07, 0xb4,
	0xb7, 0xf9, 0xff, 0xf9, 0x34, 0xfe, 0x35, 0x1e, 0xb8, 0x3f, 0x11, 0x19, 0x6d, 0xa8, 0xd5, 0xca,
	0x7e, 0x70, 0x75, 0x61, 0x7a, 0x4d, 0x9a, 0xc5, 0x28, 0x2c, 0x16, 0xc2, 0xb4, 0xd9, 0x2f, 0x08,
	0x77, 0x44, 0x86, 0xe5, 0xb0, 0xea, 0x87, 0x0e, 0x2d, 0x0f, 0xd2, 0x65, 0xbe, 0x7e, 0x60, 0xc5,
	0x4c, 0x14, 0xae, 0x5d, 0x0d, 0x1d, 0x56, 0x23, 0xc0, 0xb6, 0xf0, 0xee, 0xfb, 0xd0, 0x75, 0x97,
	0x5a, 0x62, 0xa3, 0x25, 0xd6, 0x3d, 0x5a, 0xec, 0x7f, 0xa2, 0xac, 0xf1, 0xb7, 0x11, 0xca, 0xb6,
	0x5a, 0xf1, 0x45, 0x1a, 0xe4, 0x71, 0xf5, 0xd6, 0x63, 0x5f, 0x3d, 0x55, 0x4d, 0xd0, 0x76, 0x66,
	0xb2, 0xbf, 0x0b, 0x88, 0xe7, 0xd1, 0xec, 0x35, 0xc4, 0x16, 0x3b, 0x6c, 0x48, 0xf7, 0x3c, 0x48,
	0x83, 0x3c, 0xa9, 0xae, 0x9a, 0x31, 0x58, 0x1e, 0x91, 0xfc, 0xcc, 0x64, 0xf7, 0xac, 0x72, 0xc2,
	0x79, 0x66, 0x20, 0xbe, 0x9c, 0x3d, 0x33, 0x10, 0x7b, 0x05, 0xa1, 0xd1, 0x96, 0x78, 0x38, 0x99,
	0x5e, 0x31, 0x0e, 0x91, 0xc4, 0x0e, 0x09, 0xf9, 0x6a, 0xf2, 0x27, 0xcd, 0xee, 0x61, 0x65, 0x04,
	0x35, 0x27, 0x1e, 0x4d, 0x8d, 0x51, 0xb2, 0x4f, 0x10, 0x35, 0x83, 0x25, 0x7d, 0xe6, 0x71, 0x1a,
	0xe4, 0xeb, 0x87, 0x37, 0xb7, 0x55, 0x3c, 0x7a, 0xdf, 0xa5, 0xde, 0x0b, 0x22, 0xec, 0x95, 0x1b,
	0x37, 0xc2, 0x8c, 0x41, 0x78, 0xd0, 0xf2, 0xc2, 0xef, 0x7c, 0x7c, 0x5f, 0xb3, 0xf7, 0xf0, 0xbc,
	0x47, 0x6b, 0xb4, 0xb2, 0x58, 0xfb, 0xe6, 0xc6, 0x37, 0x37, 0xb3, 0x59, 0x3a, 0xe8, 0x11, 0x5e,
	0x0a, 0x29, 0x5b, 0xf7, 0x47, 0xa2, 0xab, 0x0f, 0xad, 0x92, 0xad, 0x3a, 0x5a, 0xbe, 0x7e, 0xf2,
	0x1f, 0xd8, 0x0d, 0x2f, 0x27, 0xba, 0x4c, 0xe0, 0xce, 0x8c, 0x91, 0xb2, 0xcf, 0xf0, 0xe2, 0xbf,
	0x9c, 0x2e, 0xdd, 0x8f, 0x56, 0xc9, 0x69, 0xb9, 0xbe, 0x76, 0x9e, 0x11, 0x74, 0x1a, 0x37, 0x5b,
	0xf9, 0xba, 0xcc, 0x60, 0xd3, 0xe8, 0xf3, 0xf5, 0xd1, 0x32, 0xf1, 0x43, 0xdc, 0xcd, 0xec, 0x83,
	0x3f, 0x8b, 0x70, 0xfb, 0x65, 0xff, 0xed, 0x10, 0xf9, 0x1b, 0xfa, 0xf8, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x64, 0x5b, 0xb8, 0x5e, 0x5d, 0x02, 0x00, 0x00,
}
