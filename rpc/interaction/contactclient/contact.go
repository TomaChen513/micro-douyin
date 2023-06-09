// Code generated by goctl. DO NOT EDIT.
// Source: interaction.proto

package contactclient

import (
	"context"

	"miniTikTok/rpc/interaction/contact"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateMessageRequest     = contact.CreateMessageRequest
	Empty                    = contact.Empty
	GetFriendsListRequest    = contact.GetFriendsListRequest
	GetFriendsListResponse   = contact.GetFriendsListResponse
	GetLatestMessageRequest  = contact.GetLatestMessageRequest
	GetLatestMessageResponse = contact.GetLatestMessageResponse
	GetMessageListRequest    = contact.GetMessageListRequest
	GetMessageListResponse   = contact.GetMessageListResponse
	LoseFriendsRequest       = contact.LoseFriendsRequest
	MakeFriendsRequest       = contact.MakeFriendsRequest
	Message                  = contact.Message

	Contact interface {
		CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error)
		GetLatestMessage(ctx context.Context, in *GetLatestMessageRequest, opts ...grpc.CallOption) (*GetLatestMessageResponse, error)
		GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error)
		MakeFriends(ctx context.Context, in *MakeFriendsRequest, opts ...grpc.CallOption) (*Empty, error)
		LoseFriends(ctx context.Context, in *LoseFriendsRequest, opts ...grpc.CallOption) (*Empty, error)
		GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error)
	}

	defaultContact struct {
		cli zrpc.Client
	}
)

func NewContact(cli zrpc.Client) Contact {
	return &defaultContact{
		cli: cli,
	}
}

func (m *defaultContact) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := contact.NewContactClient(m.cli.Conn())
	return client.CreateMessage(ctx, in, opts...)
}

func (m *defaultContact) GetLatestMessage(ctx context.Context, in *GetLatestMessageRequest, opts ...grpc.CallOption) (*GetLatestMessageResponse, error) {
	client := contact.NewContactClient(m.cli.Conn())
	return client.GetLatestMessage(ctx, in, opts...)
}

func (m *defaultContact) GetMessageList(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error) {
	client := contact.NewContactClient(m.cli.Conn())
	return client.GetMessageList(ctx, in, opts...)
}

func (m *defaultContact) MakeFriends(ctx context.Context, in *MakeFriendsRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := contact.NewContactClient(m.cli.Conn())
	return client.MakeFriends(ctx, in, opts...)
}

func (m *defaultContact) LoseFriends(ctx context.Context, in *LoseFriendsRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := contact.NewContactClient(m.cli.Conn())
	return client.LoseFriends(ctx, in, opts...)
}

func (m *defaultContact) GetFriendsList(ctx context.Context, in *GetFriendsListRequest, opts ...grpc.CallOption) (*GetFriendsListResponse, error) {
	client := contact.NewContactClient(m.cli.Conn())
	return client.GetFriendsList(ctx, in, opts...)
}
