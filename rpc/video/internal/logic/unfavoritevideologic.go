package logic

import (
	"context"

	"miniTikTok/rpc/video/internal/svc"
	"miniTikTok/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteVideoLogic {
	return &UnFavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnFavoriteVideoLogic) UnFavoriteVideo(in *video.UnFavoriteVideoRequest) (*video.Empty, error) {
	// todo: add your logic here and delete this line

	return &video.Empty{}, nil
}
