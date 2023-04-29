package user

import (
	"context"

	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowListLogic {
	return &FollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowListLogic) FollowList(req *types.FollowListRequest) (resp *types.FollowListReply, err error) {
	// todo: add your logic here and delete this line

	return
}
