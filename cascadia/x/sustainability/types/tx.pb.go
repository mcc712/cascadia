// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cascadia/sustainability/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreatePenaltyAccount struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MultisigAddress string `protobuf:"bytes,3,opt,name=multisigAddress,proto3" json:"multisigAddress,omitempty"`
}

func (m *MsgCreatePenaltyAccount) Reset()         { *m = MsgCreatePenaltyAccount{} }
func (m *MsgCreatePenaltyAccount) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePenaltyAccount) ProtoMessage()    {}
func (*MsgCreatePenaltyAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_321d1b19a8f56691, []int{0}
}
func (m *MsgCreatePenaltyAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePenaltyAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePenaltyAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePenaltyAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePenaltyAccount.Merge(m, src)
}
func (m *MsgCreatePenaltyAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePenaltyAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePenaltyAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePenaltyAccount proto.InternalMessageInfo

func (m *MsgCreatePenaltyAccount) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePenaltyAccount) GetMultisigAddress() string {
	if m != nil {
		return m.MultisigAddress
	}
	return ""
}

type MsgCreatePenaltyAccountResponse struct {
}

func (m *MsgCreatePenaltyAccountResponse) Reset()         { *m = MsgCreatePenaltyAccountResponse{} }
func (m *MsgCreatePenaltyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePenaltyAccountResponse) ProtoMessage()    {}
func (*MsgCreatePenaltyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_321d1b19a8f56691, []int{1}
}
func (m *MsgCreatePenaltyAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePenaltyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePenaltyAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePenaltyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePenaltyAccountResponse.Merge(m, src)
}
func (m *MsgCreatePenaltyAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePenaltyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePenaltyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePenaltyAccountResponse proto.InternalMessageInfo

type MsgUpdatePenaltyAccount struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MultisigAddress string `protobuf:"bytes,3,opt,name=multisigAddress,proto3" json:"multisigAddress,omitempty"`
}

func (m *MsgUpdatePenaltyAccount) Reset()         { *m = MsgUpdatePenaltyAccount{} }
func (m *MsgUpdatePenaltyAccount) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePenaltyAccount) ProtoMessage()    {}
func (*MsgUpdatePenaltyAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_321d1b19a8f56691, []int{2}
}
func (m *MsgUpdatePenaltyAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePenaltyAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePenaltyAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePenaltyAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePenaltyAccount.Merge(m, src)
}
func (m *MsgUpdatePenaltyAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePenaltyAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePenaltyAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePenaltyAccount proto.InternalMessageInfo

func (m *MsgUpdatePenaltyAccount) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdatePenaltyAccount) GetMultisigAddress() string {
	if m != nil {
		return m.MultisigAddress
	}
	return ""
}

type MsgUpdatePenaltyAccountResponse struct {
}

func (m *MsgUpdatePenaltyAccountResponse) Reset()         { *m = MsgUpdatePenaltyAccountResponse{} }
func (m *MsgUpdatePenaltyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePenaltyAccountResponse) ProtoMessage()    {}
func (*MsgUpdatePenaltyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_321d1b19a8f56691, []int{3}
}
func (m *MsgUpdatePenaltyAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePenaltyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePenaltyAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePenaltyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePenaltyAccountResponse.Merge(m, src)
}
func (m *MsgUpdatePenaltyAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePenaltyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePenaltyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePenaltyAccountResponse proto.InternalMessageInfo

type MsgDeletePenaltyAccount struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgDeletePenaltyAccount) Reset()         { *m = MsgDeletePenaltyAccount{} }
func (m *MsgDeletePenaltyAccount) String() string { return proto.CompactTextString(m) }
func (*MsgDeletePenaltyAccount) ProtoMessage()    {}
func (*MsgDeletePenaltyAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_321d1b19a8f56691, []int{4}
}
func (m *MsgDeletePenaltyAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeletePenaltyAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeletePenaltyAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeletePenaltyAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeletePenaltyAccount.Merge(m, src)
}
func (m *MsgDeletePenaltyAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeletePenaltyAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeletePenaltyAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeletePenaltyAccount proto.InternalMessageInfo

func (m *MsgDeletePenaltyAccount) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgDeletePenaltyAccountResponse struct {
}

func (m *MsgDeletePenaltyAccountResponse) Reset()         { *m = MsgDeletePenaltyAccountResponse{} }
func (m *MsgDeletePenaltyAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeletePenaltyAccountResponse) ProtoMessage()    {}
func (*MsgDeletePenaltyAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_321d1b19a8f56691, []int{5}
}
func (m *MsgDeletePenaltyAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeletePenaltyAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeletePenaltyAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeletePenaltyAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeletePenaltyAccountResponse.Merge(m, src)
}
func (m *MsgDeletePenaltyAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeletePenaltyAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeletePenaltyAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeletePenaltyAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePenaltyAccount)(nil), "cascadia.sustainability.v1.MsgCreatePenaltyAccount")
	proto.RegisterType((*MsgCreatePenaltyAccountResponse)(nil), "cascadia.sustainability.v1.MsgCreatePenaltyAccountResponse")
	proto.RegisterType((*MsgUpdatePenaltyAccount)(nil), "cascadia.sustainability.v1.MsgUpdatePenaltyAccount")
	proto.RegisterType((*MsgUpdatePenaltyAccountResponse)(nil), "cascadia.sustainability.v1.MsgUpdatePenaltyAccountResponse")
	proto.RegisterType((*MsgDeletePenaltyAccount)(nil), "cascadia.sustainability.v1.MsgDeletePenaltyAccount")
	proto.RegisterType((*MsgDeletePenaltyAccountResponse)(nil), "cascadia.sustainability.v1.MsgDeletePenaltyAccountResponse")
}

func init() {
	proto.RegisterFile("cascadia/sustainability/v1/tx.proto", fileDescriptor_321d1b19a8f56691)
}

var fileDescriptor_321d1b19a8f56691 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x4e, 0x2c, 0x4e,
	0x4e, 0x4c, 0xc9, 0x4c, 0xd4, 0x2f, 0x2e, 0x2d, 0x2e, 0x49, 0xcc, 0xcc, 0x4b, 0x4c, 0xca, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x82, 0x29, 0xd2, 0x43, 0x55, 0xa4, 0x57, 0x66, 0x28, 0x65, 0x80, 0xc7, 0x80, 0x82, 0xd4,
	0xbc, 0xc4, 0x9c, 0x92, 0xca, 0xf8, 0xc4, 0xe4, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x88, 0x69, 0x4a,
	0xb1, 0x5c, 0xe2, 0xbe, 0xc5, 0xe9, 0xce, 0x45, 0xa9, 0x89, 0x25, 0xa9, 0x01, 0x10, 0x15, 0x8e,
	0x10, 0x05, 0x42, 0x12, 0x5c, 0xec, 0xc9, 0x20, 0xf1, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x18, 0x57, 0x48, 0x83, 0x8b, 0x3f, 0xb7, 0x34, 0xa7, 0x24, 0xb3, 0x38, 0x33, 0xdd,
	0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x19, 0xac, 0x02, 0x5d, 0x58, 0x49, 0x91, 0x4b,
	0x1e, 0x87, 0xf1, 0x41, 0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x50, 0x17, 0x84, 0x16, 0xa4,
	0xd0, 0xd2, 0x05, 0xd8, 0x8c, 0x87, 0xbb, 0xc0, 0x18, 0xec, 0x02, 0x97, 0xd4, 0x9c, 0x54, 0xe2,
	0x5d, 0x00, 0x35, 0x17, 0x9b, 0x26, 0x98, 0xb9, 0x46, 0xcb, 0x98, 0xb9, 0x98, 0x7d, 0x8b, 0xd3,
	0x85, 0x3a, 0x18, 0xb9, 0x44, 0xb0, 0x86, 0xb0, 0xb1, 0x1e, 0xee, 0xb8, 0xd4, 0xc3, 0x11, 0x6e,
	0x52, 0xd6, 0x64, 0x68, 0x82, 0x39, 0x09, 0xec, 0x14, 0xac, 0x41, 0x4d, 0xc8, 0x29, 0xd8, 0x34,
	0x11, 0x74, 0x0a, 0xbe, 0x50, 0x07, 0x3b, 0x05, 0x6b, 0x98, 0x13, 0x72, 0x0a, 0x36, 0x4d, 0x04,
	0x9d, 0x82, 0x2f, 0xa2, 0x9c, 0x2c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0x4a, 0x1e, 0x9e, 0xa1, 0x2a, 0xd0, 0xb3, 0x54, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38,
	0x1b, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x31, 0x22, 0xbd, 0x0f, 0xbb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreatePenaltyAccount(ctx context.Context, in *MsgCreatePenaltyAccount, opts ...grpc.CallOption) (*MsgCreatePenaltyAccountResponse, error)
	UpdatePenaltyAccount(ctx context.Context, in *MsgUpdatePenaltyAccount, opts ...grpc.CallOption) (*MsgUpdatePenaltyAccountResponse, error)
	DeletePenaltyAccount(ctx context.Context, in *MsgDeletePenaltyAccount, opts ...grpc.CallOption) (*MsgDeletePenaltyAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePenaltyAccount(ctx context.Context, in *MsgCreatePenaltyAccount, opts ...grpc.CallOption) (*MsgCreatePenaltyAccountResponse, error) {
	out := new(MsgCreatePenaltyAccountResponse)
	err := c.cc.Invoke(ctx, "/cascadia.sustainability.v1.Msg/CreatePenaltyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePenaltyAccount(ctx context.Context, in *MsgUpdatePenaltyAccount, opts ...grpc.CallOption) (*MsgUpdatePenaltyAccountResponse, error) {
	out := new(MsgUpdatePenaltyAccountResponse)
	err := c.cc.Invoke(ctx, "/cascadia.sustainability.v1.Msg/UpdatePenaltyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePenaltyAccount(ctx context.Context, in *MsgDeletePenaltyAccount, opts ...grpc.CallOption) (*MsgDeletePenaltyAccountResponse, error) {
	out := new(MsgDeletePenaltyAccountResponse)
	err := c.cc.Invoke(ctx, "/cascadia.sustainability.v1.Msg/DeletePenaltyAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreatePenaltyAccount(context.Context, *MsgCreatePenaltyAccount) (*MsgCreatePenaltyAccountResponse, error)
	UpdatePenaltyAccount(context.Context, *MsgUpdatePenaltyAccount) (*MsgUpdatePenaltyAccountResponse, error)
	DeletePenaltyAccount(context.Context, *MsgDeletePenaltyAccount) (*MsgDeletePenaltyAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePenaltyAccount(ctx context.Context, req *MsgCreatePenaltyAccount) (*MsgCreatePenaltyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePenaltyAccount not implemented")
}
func (*UnimplementedMsgServer) UpdatePenaltyAccount(ctx context.Context, req *MsgUpdatePenaltyAccount) (*MsgUpdatePenaltyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePenaltyAccount not implemented")
}
func (*UnimplementedMsgServer) DeletePenaltyAccount(ctx context.Context, req *MsgDeletePenaltyAccount) (*MsgDeletePenaltyAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePenaltyAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePenaltyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePenaltyAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePenaltyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cascadia.sustainability.v1.Msg/CreatePenaltyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePenaltyAccount(ctx, req.(*MsgCreatePenaltyAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePenaltyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePenaltyAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePenaltyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cascadia.sustainability.v1.Msg/UpdatePenaltyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePenaltyAccount(ctx, req.(*MsgUpdatePenaltyAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePenaltyAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePenaltyAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePenaltyAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cascadia.sustainability.v1.Msg/DeletePenaltyAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePenaltyAccount(ctx, req.(*MsgDeletePenaltyAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cascadia.sustainability.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePenaltyAccount",
			Handler:    _Msg_CreatePenaltyAccount_Handler,
		},
		{
			MethodName: "UpdatePenaltyAccount",
			Handler:    _Msg_UpdatePenaltyAccount_Handler,
		},
		{
			MethodName: "DeletePenaltyAccount",
			Handler:    _Msg_DeletePenaltyAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cascadia/sustainability/v1/tx.proto",
}

func (m *MsgCreatePenaltyAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePenaltyAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePenaltyAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MultisigAddress) > 0 {
		i -= len(m.MultisigAddress)
		copy(dAtA[i:], m.MultisigAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MultisigAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePenaltyAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePenaltyAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePenaltyAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePenaltyAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePenaltyAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePenaltyAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MultisigAddress) > 0 {
		i -= len(m.MultisigAddress)
		copy(dAtA[i:], m.MultisigAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MultisigAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePenaltyAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePenaltyAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePenaltyAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeletePenaltyAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeletePenaltyAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeletePenaltyAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeletePenaltyAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeletePenaltyAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeletePenaltyAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreatePenaltyAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MultisigAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreatePenaltyAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdatePenaltyAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MultisigAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdatePenaltyAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeletePenaltyAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeletePenaltyAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePenaltyAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePenaltyAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePenaltyAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreatePenaltyAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreatePenaltyAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePenaltyAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdatePenaltyAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdatePenaltyAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePenaltyAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultisigAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultisigAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpdatePenaltyAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpdatePenaltyAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePenaltyAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeletePenaltyAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeletePenaltyAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeletePenaltyAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgDeletePenaltyAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgDeletePenaltyAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeletePenaltyAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
