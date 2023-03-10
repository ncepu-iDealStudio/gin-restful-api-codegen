// 这个就是protobuf的中间文件

// 指定的当前proto语法的版本，有2和3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0
// source: TokenService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 定义request model
type VerifyTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"` // 1代表顺序
}

func (x *VerifyTokenRequest) Reset() {
	*x = VerifyTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenRequest) ProtoMessage() {}

func (x *VerifyTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyTokenRequest) Descriptor() ([]byte, []int) {
	return file_TokenService_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// 定义response model
type VerifyTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *VerifyTokenResponse) Reset() {
	*x = VerifyTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTokenResponse) ProtoMessage() {}

func (x *VerifyTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyTokenResponse) Descriptor() ([]byte, []int) {
	return file_TokenService_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyTokenResponse) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

// 定义response model
type FreeTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FreeTokenResponse) Reset() {
	*x = FreeTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TokenService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeTokenResponse) ProtoMessage() {}

func (x *FreeTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_TokenService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeTokenResponse.ProtoReflect.Descriptor instead.
func (*FreeTokenResponse) Descriptor() ([]byte, []int) {
	return file_TokenService_proto_rawDescGZIP(), []int{2}
}

var File_TokenService_proto protoreflect.FileDescriptor

var file_TokenService_proto_rawDesc = []byte{
	0x0a, 0x12, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x12, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2d, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22,
	0x13, 0x0a, 0x11, 0x46, 0x72, 0x65, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x7e, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x09, 0x46, 0x72, 0x65, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TokenService_proto_rawDescOnce sync.Once
	file_TokenService_proto_rawDescData = file_TokenService_proto_rawDesc
)

func file_TokenService_proto_rawDescGZIP() []byte {
	file_TokenService_proto_rawDescOnce.Do(func() {
		file_TokenService_proto_rawDescData = protoimpl.X.CompressGZIP(file_TokenService_proto_rawDescData)
	})
	return file_TokenService_proto_rawDescData
}

var file_TokenService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_TokenService_proto_goTypes = []interface{}{
	(*VerifyTokenRequest)(nil),  // 0: VerifyTokenRequest
	(*VerifyTokenResponse)(nil), // 1: VerifyTokenResponse
	(*FreeTokenResponse)(nil),   // 2: FreeTokenResponse
}
var file_TokenService_proto_depIdxs = []int32{
	0, // 0: TokenService.VerifyToken:input_type -> VerifyTokenRequest
	0, // 1: TokenService.FreeToken:input_type -> VerifyTokenRequest
	1, // 2: TokenService.VerifyToken:output_type -> VerifyTokenResponse
	2, // 3: TokenService.FreeToken:output_type -> FreeTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TokenService_proto_init() }
func file_TokenService_proto_init() {
	if File_TokenService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TokenService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenRequest); i {
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
		file_TokenService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTokenResponse); i {
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
		file_TokenService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeTokenResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TokenService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_TokenService_proto_goTypes,
		DependencyIndexes: file_TokenService_proto_depIdxs,
		MessageInfos:      file_TokenService_proto_msgTypes,
	}.Build()
	File_TokenService_proto = out.File
	file_TokenService_proto_rawDesc = nil
	file_TokenService_proto_goTypes = nil
	file_TokenService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenServiceClient interface {
	// 定义方法
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error)
	FreeToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*FreeTokenResponse, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*VerifyTokenResponse, error) {
	out := new(VerifyTokenResponse)
	err := c.cc.Invoke(ctx, "/TokenService/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) FreeToken(ctx context.Context, in *VerifyTokenRequest, opts ...grpc.CallOption) (*FreeTokenResponse, error) {
	out := new(FreeTokenResponse)
	err := c.cc.Invoke(ctx, "/TokenService/FreeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
type TokenServiceServer interface {
	// 定义方法
	VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error)
	FreeToken(context.Context, *VerifyTokenRequest) (*FreeTokenResponse, error)
}

// UnimplementedTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (*UnimplementedTokenServiceServer) VerifyToken(context.Context, *VerifyTokenRequest) (*VerifyTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}
func (*UnimplementedTokenServiceServer) FreeToken(context.Context, *VerifyTokenRequest) (*FreeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FreeToken not implemented")
}

func RegisterTokenServiceServer(s *grpc.Server, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TokenService/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).VerifyToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_FreeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).FreeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TokenService/FreeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).FreeToken(ctx, req.(*VerifyTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _TokenService_VerifyToken_Handler,
		},
		{
			MethodName: "FreeToken",
			Handler:    _TokenService_FreeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "TokenService.proto",
}
