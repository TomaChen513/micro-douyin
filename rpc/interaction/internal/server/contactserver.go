// Code generated by goctl. DO NOT EDIT.
// Source: interaction.proto

package server

import (
	"context"

	"miniTikTok/rpc/interaction/contact"
	"miniTikTok/rpc/interaction/internal/logic"
	"miniTikTok/rpc/interaction/internal/svc"
)

type ContactServer struct {
	svcCtx *svc.ServiceContext
	contact.UnimplementedContactServer
}

func NewContactServer(svcCtx *svc.ServiceContext) *ContactServer {
	return &ContactServer{
		svcCtx: svcCtx,
	}
}

func (s *ContactServer) CreateMessage(ctx context.Context, in *contact.CreateMessageRequest) (*contact.Empty, error) {
	l := logic.NewCreateMessageLogic(ctx, s.svcCtx)
	return l.CreateMessage(in)
}

func (s *ContactServer) GetLatestMessage(ctx context.Context, in *contact.GetLatestMessageRequest) (*contact.GetLatestMessageResponse, error) {
	l := logic.NewGetLatestMessageLogic(ctx, s.svcCtx)
	return l.GetLatestMessage(in)
}

func (s *ContactServer) GetMessageList(ctx context.Context, in *contact.GetMessageListRequest) (*contact.GetMessageListResponse, error) {
	l := logic.NewGetMessageListLogic(ctx, s.svcCtx)
	return l.GetMessageList(in)
}

func (s *ContactServer) MakeFriends(ctx context.Context, in *contact.MakeFriendsRequest) (*contact.Empty, error) {
	l := logic.NewMakeFriendsLogic(ctx, s.svcCtx)
	return l.MakeFriends(in)
}

func (s *ContactServer) LoseFriends(ctx context.Context, in *contact.LoseFriendsRequest) (*contact.Empty, error) {
	l := logic.NewLoseFriendsLogic(ctx, s.svcCtx)
	return l.LoseFriends(in)
}

func (s *ContactServer) GetFriendsList(ctx context.Context, in *contact.GetFriendsListRequest) (*contact.GetFriendsListResponse, error) {
	l := logic.NewGetFriendsListLogic(ctx, s.svcCtx)
	return l.GetFriendsList(in)
}
