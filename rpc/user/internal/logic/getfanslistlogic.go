package logic

import (
	"context"

	"miniTikTok/rpc/user/internal/svc"
	"miniTikTok/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFansListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFansListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansListLogic {
	return &GetFansListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFansListLogic) GetFansList(in *user.GetFansListRequest) (*user.GetFansListReply, error) {
	// todo: add your logic here and delete this line

	return &user.GetFansListReply{}, nil
}
