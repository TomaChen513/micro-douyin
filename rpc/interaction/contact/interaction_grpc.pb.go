// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: interaction.proto

package contact

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Contact_CreateMessage_FullMethodName    = "/contact.Contact/CreateMessage"
	Contact_GetLatestMessage_FullMethodName = "/contact.Contact/GetLatestMessage"
	Contact_GetMessageList_FullMethodName   = "/contact.Contact/GetMessageList"
	Contact_MakeFriends_FullMethodName      = "/contact.Contact/MakeFriends"
	Contact_LoseFriends_FullMethodName      = "/contact.Contact/LoseFriends"
	Contact_GetFriendsList_FullMethodName   = "/contact.Contact/GetFriendsList"
)

// ContactClient is the client API for Contact service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContactClient interface {
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error)
	GetLatestMessage(ctx context.Context, in *GetLatestMessageRequest, opts ...grpc.CallOption) (*GetLatestMessageResponse, error)
	GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error)
	MakeFriends(ctx context.Context, in *MakeFriendsRequest, opts ...grpc.CallOption) (*Empty, error)
	LoseFriends(ctx context.Context, in *LoseFriendsRequest, opts ...grpc.CallOption) (*Empty, error)
	GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error)
}

type contactClient struct {
	cc grpc.ClientConnInterface
}

func NewContactClient(cc grpc.ClientConnInterface) ContactClient {
	return &contactClient{cc}
}

func (c *contactClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Contact_CreateMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactClient) GetLatestMessage(ctx context.Context, in *GetLatestMessageRequest, opts ...grpc.CallOption) (*GetLatestMessageResponse, error) {
	out := new(GetLatestMessageResponse)
	err := c.cc.Invoke(ctx, Contact_GetLatestMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactClient) GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error) {
	out := new(GetMessageListResponse)
	err := c.cc.Invoke(ctx, Contact_GetMessageList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactClient) MakeFriends(ctx context.Context, in *MakeFriendsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Contact_MakeFriends_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactClient) LoseFriends(ctx context.Context, in *LoseFriendsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Contact_LoseFriends_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactClient) GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error) {
	out := new(GetFriendsListResponse)
	err := c.cc.Invoke(ctx, Contact_GetFriendsList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactServer is the server API for Contact service.
// All implementations must embed UnimplementedContactServer
// for forward compatibility
type ContactServer interface {
	CreateMessage(context.Context, *CreateMessageRequest) (*Empty, error)
	GetLatestMessage(context.Context, *GetLatestMessageRequest) (*GetLatestMessageResponse, error)
	GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListResponse, error)
	MakeFriends(context.Context, *MakeFriendsRequest) (*Empty, error)
	LoseFriends(context.Context, *LoseFriendsRequest) (*Empty, error)
	GetFriendsList(context.Context, *GetFriendsListRequest) (*GetFriendsListResponse, error)
	mustEmbedUnimplementedContactServer()
}

// UnimplementedContactServer must be embedded to have forward compatible implementations.
type UnimplementedContactServer struct {
}

func (UnimplementedContactServer) CreateMessage(context.Context, *CreateMessageRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMessage not implemented")
}
func (UnimplementedContactServer) GetLatestMessage(context.Context, *GetLatestMessageRequest) (*GetLatestMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestMessage not implemented")
}
func (UnimplementedContactServer) GetMessageList(context.Context, *GetMessageListRequest) (*GetMessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageList not implemented")
}
func (UnimplementedContactServer) MakeFriends(context.Context, *MakeFriendsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeFriends not implemented")
}
func (UnimplementedContactServer) LoseFriends(context.Context, *LoseFriendsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoseFriends not implemented")
}
func (UnimplementedContactServer) GetFriendsList(context.Context, *GetFriendsListRequest) (*GetFriendsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendsList not implemented")
}
func (UnimplementedContactServer) mustEmbedUnimplementedContactServer() {}

// UnsafeContactServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContactServer will
// result in compilation errors.
type UnsafeContactServer interface {
	mustEmbedUnimplementedContactServer()
}

func RegisterContactServer(s grpc.ServiceRegistrar, srv ContactServer) {
	s.RegisterService(&Contact_ServiceDesc, srv)
}

func _Contact_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_CreateMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contact_GetLatestMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).GetLatestMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_GetLatestMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).GetLatestMessage(ctx, req.(*GetLatestMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contact_GetMessageList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).GetMessageList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_GetMessageList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).GetMessageList(ctx, req.(*GetMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contact_MakeFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).MakeFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_MakeFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).MakeFriends(ctx, req.(*MakeFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contact_LoseFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoseFriendsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).LoseFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_LoseFriends_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).LoseFriends(ctx, req.(*LoseFriendsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contact_GetFriendsList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFriendsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactServer).GetFriendsList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Contact_GetFriendsList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactServer).GetFriendsList(ctx, req.(*GetFriendsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Contact_ServiceDesc is the grpc.ServiceDesc for Contact service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Contact_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "contact.Contact",
	HandlerType: (*ContactServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _Contact_CreateMessage_Handler,
		},
		{
			MethodName: "GetLatestMessage",
			Handler:    _Contact_GetLatestMessage_Handler,
		},
		{
			MethodName: "GetMessageList",
			Handler:    _Contact_GetMessageList_Handler,
		},
		{
			MethodName: "MakeFriends",
			Handler:    _Contact_MakeFriends_Handler,
		},
		{
			MethodName: "LoseFriends",
			Handler:    _Contact_LoseFriends_Handler,
		},
		{
			MethodName: "GetFriendsList",
			Handler:    _Contact_GetFriendsList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "interaction.proto",
}
