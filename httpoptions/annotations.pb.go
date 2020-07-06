// Code generated by protoc-gen-go. DO NOT EDIT.
// source: httpoptions/annotations.proto

package ease_api

import (
	fmt "fmt"
	data "github.com/binchencoder/gateway-proto/data"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

type ApiSourceType int32

const (
	ApiSourceType_EASE_GATEWAY ApiSourceType = 0
	ApiSourceType_OPEN_GATEWAY ApiSourceType = 1
)

var ApiSourceType_name = map[int32]string{
	0: "EASE_GATEWAY",
	1: "OPEN_GATEWAY",
}

var ApiSourceType_value = map[string]int32{
	"EASE_GATEWAY": 0,
	"OPEN_GATEWAY": 1,
}

func (x ApiSourceType) String() string {
	return proto.EnumName(ApiSourceType_name, int32(x))
}

func (ApiSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{0}
}

type AuthTokenType int32

const (
	AuthTokenType_EASE_AUTH_TOKEN   AuthTokenType = 0
	AuthTokenType_BASE_ACCESS_TOKEN AuthTokenType = 1
)

var AuthTokenType_name = map[int32]string{
	0: "EASE_AUTH_TOKEN",
	1: "BASE_ACCESS_TOKEN",
}

var AuthTokenType_value = map[string]int32{
	"EASE_AUTH_TOKEN":   0,
	"BASE_ACCESS_TOKEN": 1,
}

func (x AuthTokenType) String() string {
	return proto.EnumName(AuthTokenType_name, int32(x))
}

func (AuthTokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{1}
}

type SpecSourceType int32

const (
	SpecSourceType_UNSPECIFIED SpecSourceType = 0
	SpecSourceType_WEB         SpecSourceType = 1
)

var SpecSourceType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "WEB",
}

var SpecSourceType_value = map[string]int32{
	"UNSPECIFIED": 0,
	"WEB":         1,
}

func (x SpecSourceType) String() string {
	return proto.EnumName(SpecSourceType_name, int32(x))
}

func (SpecSourceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{2}
}

type LoadBalancer int32

const (
	LoadBalancer_ROUND_ROBIN LoadBalancer = 0
	LoadBalancer_CONSISTENT  LoadBalancer = 1
)

var LoadBalancer_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "CONSISTENT",
}

var LoadBalancer_value = map[string]int32{
	"ROUND_ROBIN": 0,
	"CONSISTENT":  1,
}

func (x LoadBalancer) String() string {
	return proto.EnumName(LoadBalancer_name, int32(x))
}

func (LoadBalancer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{3}
}

type OperatorType int32

const (
	OperatorType_OPERATOR_TYPE_UNKNOWN OperatorType = 0
	OperatorType_GT                    OperatorType = 1
	OperatorType_LT                    OperatorType = 2
	OperatorType_EQ                    OperatorType = 3
	OperatorType_MATCH                 OperatorType = 4
	OperatorType_NON_NIL               OperatorType = 5
	OperatorType_LEN_GT                OperatorType = 6
	OperatorType_LEN_LT                OperatorType = 7
	OperatorType_LEN_EQ                OperatorType = 8
)

var OperatorType_name = map[int32]string{
	0: "OPERATOR_TYPE_UNKNOWN",
	1: "GT",
	2: "LT",
	3: "EQ",
	4: "MATCH",
	5: "NON_NIL",
	6: "LEN_GT",
	7: "LEN_LT",
	8: "LEN_EQ",
}

var OperatorType_value = map[string]int32{
	"OPERATOR_TYPE_UNKNOWN": 0,
	"GT":                    1,
	"LT":                    2,
	"EQ":                    3,
	"MATCH":                 4,
	"NON_NIL":               5,
	"LEN_GT":                6,
	"LEN_LT":                7,
	"LEN_EQ":                8,
}

func (x OperatorType) String() string {
	return proto.EnumName(OperatorType_name, int32(x))
}

func (OperatorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{4}
}

type FunctionType int32

const (
	FunctionType_FUNCTION_TYPE_UNKNOWN FunctionType = 0
	FunctionType_TRIM                  FunctionType = 1
)

var FunctionType_name = map[int32]string{
	0: "FUNCTION_TYPE_UNKNOWN",
	1: "TRIM",
}

var FunctionType_value = map[string]int32{
	"FUNCTION_TYPE_UNKNOWN": 0,
	"TRIM":                  1,
}

func (x FunctionType) String() string {
	return proto.EnumName(FunctionType_name, int32(x))
}

func (FunctionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{5}
}

type ValueType int32

const (
	ValueType_VALUE_TYPE_UNKNOWN ValueType = 0
	ValueType_NUMBER             ValueType = 1
	ValueType_STRING             ValueType = 2
	ValueType_OBJ                ValueType = 3
)

var ValueType_name = map[int32]string{
	0: "VALUE_TYPE_UNKNOWN",
	1: "NUMBER",
	2: "STRING",
	3: "OBJ",
}

var ValueType_value = map[string]int32{
	"VALUE_TYPE_UNKNOWN": 0,
	"NUMBER":             1,
	"STRING":             2,
	"OBJ":                3,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{6}
}

type ApiMethod struct {
	LoginNotRequired     bool           `protobuf:"varint,1,opt,name=login_not_required,json=loginNotRequired,proto3" json:"login_not_required,omitempty"`
	ClientSignRequired   bool           `protobuf:"varint,2,opt,name=client_sign_required,json=clientSignRequired,proto3" json:"client_sign_required,omitempty"`
	HashKey              string         `protobuf:"bytes,3,opt,name=hash_key,json=hashKey,proto3" json:"hash_key,omitempty"`
	IsThirdParty         bool           `protobuf:"varint,4,opt,name=is_third_party,json=isThirdParty,proto3" json:"is_third_party,omitempty"`
	Timeout              string         `protobuf:"bytes,5,opt,name=timeout,proto3" json:"timeout,omitempty"`
	ApiSource            ApiSourceType  `protobuf:"varint,6,opt,name=api_source,json=apiSource,proto3,enum=ease.api.ApiSourceType" json:"api_source,omitempty"`
	TokenType            AuthTokenType  `protobuf:"varint,7,opt,name=token_type,json=tokenType,proto3,enum=ease.api.AuthTokenType" json:"token_type,omitempty"`
	SpecSourceType       SpecSourceType `protobuf:"varint,8,opt,name=spec_source_type,json=specSourceType,proto3,enum=ease.api.SpecSourceType" json:"spec_source_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ApiMethod) Reset()         { *m = ApiMethod{} }
func (m *ApiMethod) String() string { return proto.CompactTextString(m) }
func (*ApiMethod) ProtoMessage()    {}
func (*ApiMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{0}
}

func (m *ApiMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiMethod.Unmarshal(m, b)
}
func (m *ApiMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiMethod.Marshal(b, m, deterministic)
}
func (m *ApiMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiMethod.Merge(m, src)
}
func (m *ApiMethod) XXX_Size() int {
	return xxx_messageInfo_ApiMethod.Size(m)
}
func (m *ApiMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ApiMethod proto.InternalMessageInfo

func (m *ApiMethod) GetLoginNotRequired() bool {
	if m != nil {
		return m.LoginNotRequired
	}
	return false
}

func (m *ApiMethod) GetClientSignRequired() bool {
	if m != nil {
		return m.ClientSignRequired
	}
	return false
}

func (m *ApiMethod) GetHashKey() string {
	if m != nil {
		return m.HashKey
	}
	return ""
}

func (m *ApiMethod) GetIsThirdParty() bool {
	if m != nil {
		return m.IsThirdParty
	}
	return false
}

func (m *ApiMethod) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

func (m *ApiMethod) GetApiSource() ApiSourceType {
	if m != nil {
		return m.ApiSource
	}
	return ApiSourceType_EASE_GATEWAY
}

func (m *ApiMethod) GetTokenType() AuthTokenType {
	if m != nil {
		return m.TokenType
	}
	return AuthTokenType_EASE_AUTH_TOKEN
}

func (m *ApiMethod) GetSpecSourceType() SpecSourceType {
	if m != nil {
		return m.SpecSourceType
	}
	return SpecSourceType_UNSPECIFIED
}

type ServiceSpec struct {
	ServiceId            data.ServiceId `protobuf:"varint,1,opt,name=service_id,json=serviceId,proto3,enum=data.ServiceId" json:"service_id,omitempty"`
	PortName             string         `protobuf:"bytes,2,opt,name=port_name,json=portName,proto3" json:"port_name,omitempty"`
	Namespace            string         `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	GenController        bool           `protobuf:"varint,4,opt,name=gen_controller,json=genController,proto3" json:"gen_controller,omitempty"`
	Balancer             LoadBalancer   `protobuf:"varint,5,opt,name=balancer,proto3,enum=ease.api.LoadBalancer" json:"balancer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ServiceSpec) Reset()         { *m = ServiceSpec{} }
func (m *ServiceSpec) String() string { return proto.CompactTextString(m) }
func (*ServiceSpec) ProtoMessage()    {}
func (*ServiceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{1}
}

func (m *ServiceSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceSpec.Unmarshal(m, b)
}
func (m *ServiceSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceSpec.Marshal(b, m, deterministic)
}
func (m *ServiceSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceSpec.Merge(m, src)
}
func (m *ServiceSpec) XXX_Size() int {
	return xxx_messageInfo_ServiceSpec.Size(m)
}
func (m *ServiceSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceSpec proto.InternalMessageInfo

func (m *ServiceSpec) GetServiceId() data.ServiceId {
	if m != nil {
		return m.ServiceId
	}
	return data.ServiceId_SERVICE_NONE
}

func (m *ServiceSpec) GetPortName() string {
	if m != nil {
		return m.PortName
	}
	return ""
}

func (m *ServiceSpec) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceSpec) GetGenController() bool {
	if m != nil {
		return m.GenController
	}
	return false
}

func (m *ServiceSpec) GetBalancer() LoadBalancer {
	if m != nil {
		return m.Balancer
	}
	return LoadBalancer_ROUND_ROBIN
}

type ValidationRule struct {
	Operator             OperatorType `protobuf:"varint,1,opt,name=operator,proto3,enum=ease.api.OperatorType" json:"operator,omitempty"`
	Type                 ValueType    `protobuf:"varint,2,opt,name=type,proto3,enum=ease.api.ValueType" json:"type,omitempty"`
	Value                string       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Function             FunctionType `protobuf:"varint,4,opt,name=function,proto3,enum=ease.api.FunctionType" json:"function,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ValidationRule) Reset()         { *m = ValidationRule{} }
func (m *ValidationRule) String() string { return proto.CompactTextString(m) }
func (*ValidationRule) ProtoMessage()    {}
func (*ValidationRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{2}
}

func (m *ValidationRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationRule.Unmarshal(m, b)
}
func (m *ValidationRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationRule.Marshal(b, m, deterministic)
}
func (m *ValidationRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationRule.Merge(m, src)
}
func (m *ValidationRule) XXX_Size() int {
	return xxx_messageInfo_ValidationRule.Size(m)
}
func (m *ValidationRule) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationRule.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationRule proto.InternalMessageInfo

func (m *ValidationRule) GetOperator() OperatorType {
	if m != nil {
		return m.Operator
	}
	return OperatorType_OPERATOR_TYPE_UNKNOWN
}

func (m *ValidationRule) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_VALUE_TYPE_UNKNOWN
}

func (m *ValidationRule) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ValidationRule) GetFunction() FunctionType {
	if m != nil {
		return m.Function
	}
	return FunctionType_FUNCTION_TYPE_UNKNOWN
}

type ValidationRules struct {
	Rules                []*ValidationRule `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ValidationRules) Reset()         { *m = ValidationRules{} }
func (m *ValidationRules) String() string { return proto.CompactTextString(m) }
func (*ValidationRules) ProtoMessage()    {}
func (*ValidationRules) Descriptor() ([]byte, []int) {
	return fileDescriptor_1be5d6cda9fc6cfe, []int{3}
}

func (m *ValidationRules) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidationRules.Unmarshal(m, b)
}
func (m *ValidationRules) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidationRules.Marshal(b, m, deterministic)
}
func (m *ValidationRules) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidationRules.Merge(m, src)
}
func (m *ValidationRules) XXX_Size() int {
	return xxx_messageInfo_ValidationRules.Size(m)
}
func (m *ValidationRules) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidationRules.DiscardUnknown(m)
}

var xxx_messageInfo_ValidationRules proto.InternalMessageInfo

func (m *ValidationRules) GetRules() []*ValidationRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

var E_Http = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*HttpRule)(nil),
	Field:         108345,
	Name:          "ease.api.http",
	Tag:           "bytes,108345,opt,name=http",
	Filename:      "httpoptions/annotations.proto",
}

var E_Method = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*ApiMethod)(nil),
	Field:         108361,
	Name:          "ease.api.method",
	Tag:           "bytes,108361,opt,name=method",
	Filename:      "httpoptions/annotations.proto",
}

var E_ServiceSpec = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*ServiceSpec)(nil),
	Field:         108349,
	Name:          "ease.api.service_spec",
	Tag:           "bytes,108349,opt,name=service_spec",
	Filename:      "httpoptions/annotations.proto",
}

var E_Rules = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*ValidationRules)(nil),
	Field:         108102,
	Name:          "ease.api.rules",
	Tag:           "bytes,108102,opt,name=rules",
	Filename:      "httpoptions/annotations.proto",
}

func init() {
	proto.RegisterEnum("ease.api.ApiSourceType", ApiSourceType_name, ApiSourceType_value)
	proto.RegisterEnum("ease.api.AuthTokenType", AuthTokenType_name, AuthTokenType_value)
	proto.RegisterEnum("ease.api.SpecSourceType", SpecSourceType_name, SpecSourceType_value)
	proto.RegisterEnum("ease.api.LoadBalancer", LoadBalancer_name, LoadBalancer_value)
	proto.RegisterEnum("ease.api.OperatorType", OperatorType_name, OperatorType_value)
	proto.RegisterEnum("ease.api.FunctionType", FunctionType_name, FunctionType_value)
	proto.RegisterEnum("ease.api.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*ApiMethod)(nil), "ease.api.ApiMethod")
	proto.RegisterType((*ServiceSpec)(nil), "ease.api.ServiceSpec")
	proto.RegisterType((*ValidationRule)(nil), "ease.api.ValidationRule")
	proto.RegisterType((*ValidationRules)(nil), "ease.api.ValidationRules")
	proto.RegisterExtension(E_Http)
	proto.RegisterExtension(E_Method)
	proto.RegisterExtension(E_ServiceSpec)
	proto.RegisterExtension(E_Rules)
}

func init() { proto.RegisterFile("httpoptions/annotations.proto", fileDescriptor_1be5d6cda9fc6cfe) }

var fileDescriptor_1be5d6cda9fc6cfe = []byte{
	// 966 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0x4d, 0x6f, 0xe3, 0x54,
	0x17, 0xae, 0xd3, 0x34, 0x1f, 0x27, 0x69, 0x7a, 0xdf, 0xdb, 0x69, 0x5f, 0x77, 0x60, 0x20, 0xaa,
	0x40, 0x94, 0x08, 0x39, 0x28, 0x95, 0x58, 0x94, 0x95, 0x93, 0xba, 0xad, 0x69, 0x62, 0x67, 0x6c,
	0xa7, 0xd5, 0x88, 0x85, 0xe5, 0x3a, 0x77, 0x12, 0x6b, 0x1c, 0x5f, 0x63, 0xdf, 0x8c, 0x94, 0x05,
	0x7f, 0x86, 0x3f, 0x81, 0x58, 0xb0, 0x45, 0x42, 0x42, 0x6c, 0xf8, 0x37, 0xac, 0xd0, 0xbd, 0xb6,
	0xf3, 0x31, 0x15, 0x62, 0x13, 0x9f, 0xf3, 0x9c, 0xe7, 0x3c, 0xc7, 0xf7, 0xdc, 0xe3, 0x13, 0x78,
	0x35, 0x67, 0x2c, 0xa6, 0x31, 0x0b, 0x68, 0x94, 0x76, 0xbd, 0x28, 0xa2, 0xcc, 0x13, 0xb6, 0x12,
	0x27, 0x94, 0x51, 0x5c, 0x23, 0x5e, 0x4a, 0x14, 0x2f, 0x0e, 0x5e, 0xb6, 0x67, 0x94, 0xce, 0x42,
	0xd2, 0x15, 0xf8, 0xd3, 0xf2, 0x6d, 0x77, 0x4a, 0x52, 0x3f, 0x09, 0x62, 0x46, 0x93, 0x8c, 0xfb,
	0xf2, 0x74, 0x5b, 0x8a, 0xdb, 0x39, 0x7e, 0x34, 0xf5, 0x98, 0xd7, 0xe5, 0x3f, 0x19, 0x70, 0xfe,
	0x77, 0x09, 0xea, 0x6a, 0x1c, 0x8c, 0x08, 0x9b, 0xd3, 0x29, 0xfe, 0x0a, 0x70, 0x48, 0x67, 0x41,
	0xe4, 0x46, 0x94, 0xb9, 0x09, 0xf9, 0x61, 0x19, 0x24, 0x64, 0x2a, 0x4b, 0x6d, 0xe9, 0xa2, 0x66,
	0x21, 0x11, 0x31, 0x28, 0xb3, 0x72, 0x1c, 0x7f, 0x0d, 0x2f, 0xfc, 0x30, 0x20, 0x11, 0x73, 0xd3,
	0x60, 0x16, 0x6d, 0xf8, 0x25, 0xc1, 0xc7, 0x59, 0xcc, 0x0e, 0x66, 0xd1, 0x3a, 0xe3, 0x0c, 0x6a,
	0x73, 0x2f, 0x9d, 0xbb, 0xef, 0xc8, 0x4a, 0xde, 0x6f, 0x4b, 0x17, 0x75, 0xab, 0xca, 0xfd, 0x7b,
	0xb2, 0xc2, 0x9f, 0x41, 0x2b, 0x48, 0x5d, 0x36, 0x0f, 0x92, 0xa9, 0x1b, 0x7b, 0x09, 0x5b, 0xc9,
	0x65, 0x21, 0xd3, 0x0c, 0x52, 0x87, 0x83, 0x63, 0x8e, 0x61, 0x19, 0xaa, 0x2c, 0x58, 0x10, 0xba,
	0x64, 0xf2, 0x41, 0x96, 0x9f, 0xbb, 0xf8, 0x1b, 0x00, 0x2f, 0x0e, 0xdc, 0x94, 0x2e, 0x13, 0x9f,
	0xc8, 0x95, 0xb6, 0x74, 0xd1, 0xea, 0xfd, 0x5f, 0x29, 0x5a, 0xa6, 0xa8, 0x71, 0x60, 0x8b, 0x90,
	0xb3, 0x8a, 0x89, 0x55, 0xf7, 0x0a, 0x97, 0xe7, 0x31, 0xfa, 0x8e, 0x44, 0x2e, 0x5b, 0xc5, 0x44,
	0xae, 0x3e, 0xcb, 0x5b, 0xb2, 0xb9, 0xc3, 0xe3, 0x59, 0x1e, 0x2b, 0x4c, 0xdc, 0x07, 0x94, 0xc6,
	0xc4, 0xcf, 0x0b, 0x66, 0xd9, 0x35, 0x91, 0x2d, 0x6f, 0xb2, 0xed, 0x98, 0xf8, 0x5b, 0x65, 0x5b,
	0xe9, 0x8e, 0x7f, 0xfe, 0x97, 0x04, 0x0d, 0x9b, 0x24, 0xef, 0x03, 0x9f, 0x70, 0x26, 0x56, 0x00,
	0xd2, 0xcc, 0x75, 0x83, 0xac, 0xed, 0xad, 0xde, 0x91, 0x22, 0x6e, 0x2b, 0xa7, 0xe9, 0x53, 0xab,
	0x9e, 0x16, 0x26, 0xfe, 0x08, 0xea, 0x31, 0x4d, 0x98, 0x1b, 0x79, 0x0b, 0x22, 0xba, 0x5e, 0xb7,
	0x6a, 0x1c, 0x30, 0xbc, 0x05, 0xc1, 0x1f, 0x43, 0x9d, 0xe3, 0x69, 0xec, 0xf9, 0x24, 0x6f, 0xf6,
	0x06, 0xc0, 0x9f, 0x43, 0x6b, 0x46, 0x22, 0xd7, 0xa7, 0x11, 0x4b, 0x68, 0x18, 0x92, 0x24, 0x6f,
	0xf7, 0xe1, 0x8c, 0x44, 0x83, 0x35, 0x88, 0x7b, 0x50, 0x7b, 0xf2, 0x42, 0x2f, 0xf2, 0x49, 0x22,
	0x1a, 0xde, 0xea, 0x9d, 0x6e, 0x4e, 0x37, 0xa4, 0xde, 0xb4, 0x9f, 0x47, 0xad, 0x35, 0xef, 0xfc,
	0x67, 0x09, 0x5a, 0x0f, 0x5e, 0x18, 0x4c, 0xc5, 0xf4, 0x5a, 0xcb, 0x90, 0x70, 0x19, 0x1a, 0x93,
	0xc4, 0x63, 0x34, 0xc9, 0x8f, 0xb5, 0x25, 0x63, 0xe6, 0x11, 0xd1, 0xa2, 0x35, 0x0f, 0x7f, 0x01,
	0x65, 0xd1, 0xd4, 0x92, 0xe0, 0x1f, 0x6f, 0xf8, 0x0f, 0x5e, 0xb8, 0xcc, 0xfa, 0x29, 0x08, 0xf8,
	0x05, 0x1c, 0xbc, 0xe7, 0x50, 0x7e, 0xc8, 0xcc, 0xe1, 0x25, 0xdf, 0x2e, 0x23, 0x9f, 0xbf, 0x82,
	0x38, 0xda, 0x4e, 0xc9, 0x9b, 0x3c, 0x92, 0x95, 0x2c, 0x78, 0xe7, 0x2a, 0x1c, 0xed, 0xbe, 0x78,
	0x8a, 0x15, 0x38, 0x48, 0xb8, 0x21, 0x4b, 0xed, 0xfd, 0x8b, 0xc6, 0xf6, 0xdd, 0xee, 0x32, 0xad,
	0x8c, 0xd6, 0xb9, 0x84, 0xc3, 0x9d, 0x51, 0xc3, 0x08, 0x9a, 0x9a, 0x6a, 0x6b, 0xee, 0xad, 0xea,
	0x68, 0x8f, 0xea, 0x1b, 0xb4, 0xc7, 0x11, 0x73, 0xac, 0x19, 0x6b, 0x44, 0xea, 0x7c, 0x0b, 0x87,
	0x3b, 0x73, 0x86, 0x8f, 0xe1, 0x48, 0x24, 0xa9, 0x13, 0xe7, 0xce, 0x75, 0xcc, 0x7b, 0xcd, 0x40,
	0x7b, 0xf8, 0x04, 0xfe, 0xd7, 0x17, 0xe0, 0x60, 0xa0, 0xd9, 0x76, 0x0e, 0x4b, 0x9d, 0x0e, 0xb4,
	0x76, 0xc7, 0x0c, 0x1f, 0x41, 0x63, 0x62, 0xd8, 0x63, 0x6d, 0xa0, 0xdf, 0xe8, 0xda, 0x35, 0xda,
	0xc3, 0x55, 0xd8, 0x7f, 0xd4, 0xfa, 0x48, 0xea, 0x74, 0xa1, 0xb9, 0x7d, 0x69, 0x9c, 0x69, 0x99,
	0x13, 0xe3, 0xda, 0xb5, 0xcc, 0xbe, 0xce, 0x6b, 0xb4, 0x00, 0x06, 0xa6, 0x61, 0xeb, 0xb6, 0xa3,
	0x19, 0x0e, 0x92, 0x3a, 0x3f, 0x42, 0x73, 0xfb, 0x7a, 0xf0, 0x19, 0x9c, 0x98, 0x63, 0xcd, 0x52,
	0x1d, 0xd3, 0x72, 0x9d, 0x37, 0x63, 0xcd, 0x9d, 0x18, 0xf7, 0x86, 0xf9, 0xc8, 0x53, 0x2b, 0x50,
	0xba, 0x75, 0x90, 0xc4, 0x9f, 0x43, 0x07, 0x95, 0xf8, 0x53, 0x7b, 0x8d, 0xf6, 0x71, 0x1d, 0x0e,
	0x46, 0xaa, 0x33, 0xb8, 0x43, 0x65, 0xdc, 0x80, 0xaa, 0x61, 0x1a, 0xae, 0xa1, 0x0f, 0xd1, 0x01,
	0x06, 0xa8, 0x0c, 0x79, 0x17, 0x1c, 0x54, 0x29, 0xec, 0xa1, 0x83, 0xaa, 0x85, 0xad, 0xbd, 0x46,
	0xb5, 0xce, 0x25, 0x34, 0xb7, 0xaf, 0x8a, 0x97, 0xbf, 0x99, 0x18, 0x03, 0x47, 0x37, 0x8d, 0x0f,
	0xcb, 0xd7, 0xa0, 0xec, 0x58, 0xfa, 0x08, 0x49, 0x9d, 0x6b, 0xa8, 0xaf, 0x47, 0x04, 0x9f, 0x02,
	0x7e, 0x50, 0x87, 0x13, 0xed, 0x43, 0x3a, 0x40, 0xc5, 0x98, 0x8c, 0xfa, 0x9a, 0x85, 0x24, 0x6e,
	0xdb, 0x8e, 0xa5, 0x1b, 0xb7, 0xa8, 0xc4, 0x5b, 0x65, 0xf6, 0xbf, 0x43, 0xfb, 0x57, 0x77, 0x50,
	0xe6, 0x7b, 0x13, 0x7f, 0xa2, 0x64, 0xcb, 0x56, 0x29, 0x96, 0xad, 0x92, 0xed, 0x4a, 0x33, 0x5b,
	0xae, 0xf2, 0x2f, 0x7f, 0xf2, 0x5d, 0xd3, 0xe8, 0xe1, 0xcd, 0x64, 0xdc, 0x31, 0x16, 0x8b, 0x99,
	0x10, 0x0a, 0x57, 0x23, 0xa8, 0x2c, 0xb2, 0xf5, 0xfa, 0x5f, 0x5a, 0xbf, 0xe7, 0x5a, 0xc7, 0x3b,
	0x7b, 0x2b, 0xe3, 0x58, 0xb9, 0xc8, 0xd5, 0xf7, 0xd0, 0x2c, 0x96, 0x04, 0x5f, 0x27, 0xf8, 0xd3,
	0x67, 0xa2, 0xf9, 0xae, 0x28, 0x54, 0x7f, 0xcd, 0x55, 0x4f, 0xb6, 0xf6, 0xd2, 0x66, 0xe9, 0x58,
	0x8d, 0x74, 0xe3, 0x5c, 0x8d, 0xf3, 0x71, 0xc7, 0xaf, 0x9e, 0xa9, 0xde, 0x04, 0x24, 0x5c, 0xbf,
	0xe9, 0x6f, 0x7f, 0x64, 0x9a, 0x67, 0xff, 0xf6, 0x3d, 0xa4, 0xf9, 0x07, 0xd1, 0xff, 0x12, 0x9a,
	0x3e, 0x5d, 0xac, 0x69, 0x7d, 0xa4, 0x6e, 0xfe, 0xd8, 0xc6, 0x5c, 0x7b, 0x2c, 0xfd, 0x54, 0x2a,
	0x6b, 0xea, 0x58, 0x7f, 0xaa, 0x88, 0x5a, 0x97, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x31, 0x32,
	0xc5, 0x86, 0x08, 0x07, 0x00, 0x00,
}
