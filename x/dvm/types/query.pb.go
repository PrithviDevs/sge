// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sge/dvm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfb9990f7b27fb76, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfb9990f7b27fb76, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryPubKeysRequest struct {
}

func (m *QueryPubKeysRequest) Reset()         { *m = QueryPubKeysRequest{} }
func (m *QueryPubKeysRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPubKeysRequest) ProtoMessage()    {}
func (*QueryPubKeysRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfb9990f7b27fb76, []int{2}
}
func (m *QueryPubKeysRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPubKeysRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPubKeysRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPubKeysRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPubKeysRequest.Merge(m, src)
}
func (m *QueryPubKeysRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryPubKeysRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPubKeysRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPubKeysRequest proto.InternalMessageInfo

type QueryPubKeysResponse struct {
	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (m *QueryPubKeysResponse) Reset()         { *m = QueryPubKeysResponse{} }
func (m *QueryPubKeysResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPubKeysResponse) ProtoMessage()    {}
func (*QueryPubKeysResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfb9990f7b27fb76, []int{3}
}
func (m *QueryPubKeysResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryPubKeysResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryPubKeysResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryPubKeysResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPubKeysResponse.Merge(m, src)
}
func (m *QueryPubKeysResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryPubKeysResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPubKeysResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPubKeysResponse proto.InternalMessageInfo

func (m *QueryPubKeysResponse) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "sgenetwork.sge.dvm.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "sgenetwork.sge.dvm.QueryParamsResponse")
	proto.RegisterType((*QueryPubKeysRequest)(nil), "sgenetwork.sge.dvm.QueryPubKeysRequest")
	proto.RegisterType((*QueryPubKeysResponse)(nil), "sgenetwork.sge.dvm.QueryPubKeysResponse")
}

func init() { proto.RegisterFile("sge/dvm/query.proto", fileDescriptor_cfb9990f7b27fb76) }

var fileDescriptor_cfb9990f7b27fb76 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4f, 0xe2, 0x40,
	0x14, 0xc7, 0x5b, 0x96, 0x65, 0xb3, 0xb3, 0xb7, 0xa1, 0x9b, 0x90, 0xee, 0x6e, 0x97, 0x54, 0x83,
	0x84, 0xc4, 0x4e, 0xc0, 0x8b, 0x47, 0xc3, 0xd5, 0x83, 0xca, 0xd1, 0x0b, 0x99, 0xc2, 0xcb, 0xd8,
	0x40, 0x3b, 0xa5, 0x33, 0x45, 0x7b, 0xf1, 0xe0, 0xc9, 0xa3, 0xd1, 0x7f, 0x8a, 0x23, 0x89, 0x17,
	0x4f, 0xc6, 0x80, 0x7f, 0x88, 0x61, 0x3a, 0x98, 0x34, 0x10, 0xb9, 0x4d, 0xde, 0x7c, 0xbe, 0x3f,
	0xde, 0x64, 0x50, 0x55, 0x30, 0x20, 0xc3, 0x69, 0x48, 0x26, 0x29, 0x24, 0x99, 0x17, 0x27, 0x5c,
	0x72, 0x8c, 0x05, 0x83, 0x08, 0xe4, 0x35, 0x4f, 0x46, 0x9e, 0x60, 0xe0, 0x0d, 0xa7, 0xa1, 0x6d,
	0x31, 0xce, 0xb8, 0xba, 0x26, 0xab, 0x53, 0x4e, 0xda, 0x7f, 0x19, 0xe7, 0x6c, 0x0c, 0x84, 0xc6,
	0x01, 0xa1, 0x51, 0xc4, 0x25, 0x95, 0x01, 0x8f, 0x84, 0xbe, 0x6d, 0x0d, 0xb8, 0x08, 0xb9, 0x20,
	0x3e, 0x15, 0x90, 0x07, 0x90, 0x69, 0xdb, 0x07, 0x49, 0xdb, 0x24, 0xa6, 0x2c, 0x88, 0x14, 0xac,
	0x59, 0x6b, 0x5d, 0x24, 0xa6, 0x09, 0x0d, 0xb5, 0x83, 0x6b, 0x21, 0x7c, 0xb1, 0xd2, 0x9d, 0xab,
	0x61, 0x0f, 0x26, 0x29, 0x08, 0xe9, 0x9e, 0xa1, 0x6a, 0x61, 0x2a, 0x62, 0x1e, 0x09, 0xc0, 0xc7,
	0xa8, 0x92, 0x8b, 0x6b, 0x66, 0xdd, 0x6c, 0xfe, 0xea, 0xd8, 0xde, 0xe6, 0x1e, 0x5e, 0xae, 0xe9,
	0x96, 0x67, 0xaf, 0xff, 0x8d, 0x9e, 0xe6, 0xdd, 0xdf, 0x6b, 0xc3, 0xd4, 0x3f, 0x85, 0xec, 0x33,
	0xa7, 0x85, 0xac, 0xe2, 0x58, 0x07, 0x61, 0x54, 0x1e, 0x07, 0x42, 0xd6, 0xcc, 0xfa, 0xb7, 0xe6,
	0xcf, 0x9e, 0x3a, 0x77, 0x1e, 0x4b, 0xe8, 0xbb, 0x82, 0xf1, 0x2d, 0xaa, 0xe4, 0x21, 0xb8, 0xb1,
	0xad, 0xc0, 0xe6, 0x3e, 0xf6, 0xc1, 0x4e, 0x2e, 0x0f, 0x76, 0xf7, 0xee, 0x9e, 0xdf, 0x9f, 0x4a,
	0xff, 0xf0, 0x1f, 0x22, 0x18, 0x1c, 0x6a, 0x05, 0x29, 0xbe, 0x1c, 0xbe, 0x37, 0xd1, 0x0f, 0xdd,
	0x18, 0x7f, 0xe1, 0x5c, 0x58, 0xd5, 0x6e, 0xee, 0x06, 0x75, 0x87, 0x96, 0xea, 0xb0, 0x8f, 0xdd,
	0xed, 0x1d, 0x52, 0xbf, 0x3f, 0x82, 0x4c, 0xf4, 0x57, 0x8f, 0xd2, 0x3d, 0x99, 0x2d, 0x1c, 0x73,
	0xbe, 0x70, 0xcc, 0xb7, 0x85, 0x63, 0x3e, 0x2c, 0x1d, 0x63, 0xbe, 0x74, 0x8c, 0x97, 0xa5, 0x63,
	0x5c, 0x36, 0x58, 0x20, 0xaf, 0x52, 0xdf, 0x1b, 0xf0, 0x70, 0xc3, 0xe7, 0x46, 0x39, 0xc9, 0x2c,
	0x06, 0xe1, 0x57, 0xd4, 0x3f, 0x38, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x74, 0xa2, 0xa9, 0x14,
	0xa8, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of PubKeys items.
	PubKeys(ctx context.Context, in *QueryPubKeysRequest, opts ...grpc.CallOption) (*QueryPubKeysResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.dvm.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) PubKeys(ctx context.Context, in *QueryPubKeysRequest, opts ...grpc.CallOption) (*QueryPubKeysResponse, error) {
	out := new(QueryPubKeysResponse)
	err := c.cc.Invoke(ctx, "/sgenetwork.sge.dvm.Query/PubKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of PubKeys items.
	PubKeys(context.Context, *QueryPubKeysRequest) (*QueryPubKeysResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) PubKeys(ctx context.Context, req *QueryPubKeysRequest) (*QueryPubKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PubKeys not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.dvm.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_PubKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPubKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).PubKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sgenetwork.sge.dvm.Query/PubKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).PubKeys(ctx, req.(*QueryPubKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sgenetwork.sge.dvm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "PubKeys",
			Handler:    _Query_PubKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sge/dvm/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryPubKeysRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPubKeysRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPubKeysRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryPubKeysResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryPubKeysResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryPubKeysResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.List[iNdEx])
			copy(dAtA[i:], m.List[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.List[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryPubKeysRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryPubKeysResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, s := range m.List {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPubKeysRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPubKeysRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPubKeysRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryPubKeysResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryPubKeysResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryPubKeysResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
