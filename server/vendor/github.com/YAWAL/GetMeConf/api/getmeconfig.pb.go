// Code generated by protoc-gen-go. DO NOT EDIT.
// source: getmeconfig.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	getmeconfig.proto

It has these top-level messages:
	ConfigInfo
	Config
*/
package api

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

type ConfigInfo struct {
	ConfigId   string `protobuf:"bytes,1,opt,name=ConfigId" json:"ConfigId,omitempty"`
	ConfigPath string `protobuf:"bytes,2,opt,name=ConfigPath" json:"ConfigPath,omitempty"`
}

func (m *ConfigInfo) Reset()                    { *m = ConfigInfo{} }
func (m *ConfigInfo) String() string            { return proto.CompactTextString(m) }
func (*ConfigInfo) ProtoMessage()               {}
func (*ConfigInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConfigInfo) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func (m *ConfigInfo) GetConfigPath() string {
	if m != nil {
		return m.ConfigPath
	}
	return ""
}

type Config struct {
	Config []byte `protobuf:"bytes,1,opt,name=Config,proto3" json:"Config,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetConfig() []byte {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigInfo)(nil), "api.ConfigInfo")
	proto.RegisterType((*Config)(nil), "api.Config")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ConfigService service

type ConfigServiceClient interface {
	SearchConfig(ctx context.Context, in *ConfigInfo, opts ...grpc.CallOption) (*ConfigInfo, error)
	GetConfig(ctx context.Context, in *ConfigInfo, opts ...grpc.CallOption) (*Config, error)
}

type configServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigServiceClient(cc *grpc.ClientConn) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) SearchConfig(ctx context.Context, in *ConfigInfo, opts ...grpc.CallOption) (*ConfigInfo, error) {
	out := new(ConfigInfo)
	err := grpc.Invoke(ctx, "/api.ConfigService/SearchConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) GetConfig(ctx context.Context, in *ConfigInfo, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := grpc.Invoke(ctx, "/api.ConfigService/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigService service

type ConfigServiceServer interface {
	SearchConfig(context.Context, *ConfigInfo) (*ConfigInfo, error)
	GetConfig(context.Context, *ConfigInfo) (*Config, error)
}

func RegisterConfigServiceServer(s *grpc.Server, srv ConfigServiceServer) {
	s.RegisterService(&_ConfigService_serviceDesc, srv)
}

func _ConfigService_SearchConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).SearchConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ConfigService/SearchConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).SearchConfig(ctx, req.(*ConfigInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ConfigService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).GetConfig(ctx, req.(*ConfigInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchConfig",
			Handler:    _ConfigService_SearchConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _ConfigService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "getmeconfig.proto",
}

func init() { proto.RegisterFile("getmeconfig.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0x4f, 0x2d, 0xc9,
	0x4d, 0x4d, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e,
	0x2c, 0xc8, 0x54, 0xf2, 0xe0, 0xe2, 0x72, 0x06, 0x0b, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b, 0x49, 0x71,
	0x71, 0x40, 0x79, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x1c, 0x4c,
	0x65, 0x40, 0x62, 0x49, 0x86, 0x04, 0x13, 0x58, 0x16, 0x49, 0x44, 0x49, 0x81, 0x8b, 0x0d, 0xc2,
	0x13, 0x12, 0x83, 0xb1, 0xc0, 0x66, 0xf0, 0x04, 0x41, 0x79, 0x46, 0x39, 0x5c, 0xbc, 0x10, 0x56,
	0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x01, 0x17, 0x4f, 0x70, 0x6a, 0x62, 0x51, 0x72,
	0x06, 0x54, 0x23, 0xbf, 0x5e, 0x62, 0x41, 0xa6, 0x1e, 0xc2, 0x3d, 0x52, 0xe8, 0x02, 0x42, 0x9a,
	0x5c, 0x9c, 0xee, 0xa9, 0x25, 0xb8, 0x94, 0x73, 0x23, 0x09, 0x24, 0xb1, 0x81, 0x7d, 0x69, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x78, 0x81, 0xda, 0xfa, 0x00, 0x00, 0x00,
}