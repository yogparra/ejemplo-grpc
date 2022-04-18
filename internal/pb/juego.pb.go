// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: juego.proto

package pb

import (
	context "context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Juego struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tipo   string `protobuf:"bytes,2,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Nombre string `protobuf:"bytes,3,opt,name=nombre,proto3" json:"nombre,omitempty"`
}

func (x *Juego) Reset() {
	*x = Juego{}
	if protoimpl.UnsafeEnabled {
		mi := &file_juego_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Juego) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Juego) ProtoMessage() {}

func (x *Juego) ProtoReflect() protoreflect.Message {
	mi := &file_juego_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Juego.ProtoReflect.Descriptor instead.
func (*Juego) Descriptor() ([]byte, []int) {
	return file_juego_proto_rawDescGZIP(), []int{0}
}

func (x *Juego) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Juego) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Juego) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

type GetJuegoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetJuegoRequest) Reset() {
	*x = GetJuegoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_juego_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJuegoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJuegoRequest) ProtoMessage() {}

func (x *GetJuegoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_juego_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJuegoRequest.ProtoReflect.Descriptor instead.
func (*GetJuegoRequest) Descriptor() ([]byte, []int) {
	return file_juego_proto_rawDescGZIP(), []int{1}
}

func (x *GetJuegoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetJuegoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Juego *Juego `protobuf:"bytes,1,opt,name=juego,proto3" json:"juego,omitempty"`
}

func (x *GetJuegoResponse) Reset() {
	*x = GetJuegoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_juego_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJuegoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJuegoResponse) ProtoMessage() {}

func (x *GetJuegoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_juego_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJuegoResponse.ProtoReflect.Descriptor instead.
func (*GetJuegoResponse) Descriptor() ([]byte, []int) {
	return file_juego_proto_rawDescGZIP(), []int{2}
}

func (x *GetJuegoResponse) GetJuego() *Juego {
	if x != nil {
		return x.Juego
	}
	return nil
}

var File_juego_proto protoreflect.FileDescriptor

var file_juego_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6a, 0x75, 0x65, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x05, 0x4a,
	0x75, 0x65, 0x67, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x30, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x05, 0x6a, 0x75, 0x65, 0x67, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x52, 0x05,
	0x6a, 0x75, 0x65, 0x67, 0x6f, 0x32, 0x5d, 0x0a, 0x0c, 0x4a, 0x75, 0x65, 0x67, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x65, 0x67,
	0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x65, 0x67, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x75, 0x65,
	0x67, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x63, 0x69,
	0x61, 0x72, 0x69, 0x6f, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_juego_proto_rawDescOnce sync.Once
	file_juego_proto_rawDescData = file_juego_proto_rawDesc
)

func file_juego_proto_rawDescGZIP() []byte {
	file_juego_proto_rawDescOnce.Do(func() {
		file_juego_proto_rawDescData = protoimpl.X.CompressGZIP(file_juego_proto_rawDescData)
	})
	return file_juego_proto_rawDescData
}

var file_juego_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_juego_proto_goTypes = []interface{}{
	(*Juego)(nil),            // 0: Juego
	(*GetJuegoRequest)(nil),  // 1: GetJuegoRequest
	(*GetJuegoResponse)(nil), // 2: GetJuegoResponse
}
var file_juego_proto_depIdxs = []int32{
	0, // 0: GetJuegoResponse.juego:type_name -> Juego
	1, // 1: JuegoService.GetJuegoById:input_type -> GetJuegoRequest
	2, // 2: JuegoService.GetJuegoById:output_type -> GetJuegoResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_juego_proto_init() }
func file_juego_proto_init() {
	if File_juego_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_juego_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Juego); i {
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
		file_juego_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJuegoRequest); i {
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
		file_juego_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJuegoResponse); i {
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
			RawDescriptor: file_juego_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_juego_proto_goTypes,
		DependencyIndexes: file_juego_proto_depIdxs,
		MessageInfos:      file_juego_proto_msgTypes,
	}.Build()
	File_juego_proto = out.File
	file_juego_proto_rawDesc = nil
	file_juego_proto_goTypes = nil
	file_juego_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JuegoServiceClient is the client API for JuegoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JuegoServiceClient interface {
	GetJuegoById(ctx context.Context, in *GetJuegoRequest, opts ...grpc.CallOption) (*GetJuegoResponse, error)
}

type juegoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJuegoServiceClient(cc grpc.ClientConnInterface) JuegoServiceClient {
	return &juegoServiceClient{cc}
}

func (c *juegoServiceClient) GetJuegoById(ctx context.Context, in *GetJuegoRequest, opts ...grpc.CallOption) (*GetJuegoResponse, error) {
	out := new(GetJuegoResponse)
	err := c.cc.Invoke(ctx, "/JuegoService/GetJuegoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JuegoServiceServer is the server API for JuegoService service.
type JuegoServiceServer interface {
	GetJuegoById(context.Context, *GetJuegoRequest) (*GetJuegoResponse, error)
}

// UnimplementedJuegoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJuegoServiceServer struct {
}

func (*UnimplementedJuegoServiceServer) GetJuegoById(context.Context, *GetJuegoRequest) (*GetJuegoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJuegoById not implemented")
}

func RegisterJuegoServiceServer(s *grpc.Server, srv JuegoServiceServer) {
	s.RegisterService(&_JuegoService_serviceDesc, srv)
}

func _JuegoService_GetJuegoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJuegoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JuegoServiceServer).GetJuegoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/JuegoService/GetJuegoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JuegoServiceServer).GetJuegoById(ctx, req.(*GetJuegoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JuegoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "JuegoService",
	HandlerType: (*JuegoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJuegoById",
			Handler:    _JuegoService_GetJuegoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "juego.proto",
}
