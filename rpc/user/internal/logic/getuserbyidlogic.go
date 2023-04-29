package logic

import (
	"context"

	"miniTikTok/rpc/user/internal/svc"
	"miniTikTok/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *user.GetUserByIdRequest) (*user.GetUserReply, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserReply{}, nil
}
