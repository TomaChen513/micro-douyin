package logic

import (
	"context"

	"miniTikTok/rpc/interaction/contact"
	"miniTikTok/rpc/interaction/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakeFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeFriendsLogic {
	return &MakeFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeFriendsLogic) MakeFriends(in *contact.MakeFriendsRequest) (*contact.Empty, error) {
	// todo: add your logic here and delete this line

	return &contact.Empty{}, nil
}
