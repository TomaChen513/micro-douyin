package logic

import (
	"context"

	"miniTikTok/rpc/user/internal/svc"
	"miniTikTok/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByNameLogic {
	return &GetUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByNameLogic) GetUserByName(in *user.GetUserByNameRequest) (*user.GetUserReply, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserReply{}, nil
}
