package logic

import (
	"context"

	"miniTikTok/rpc/video/internal/svc"
	"miniTikTok/rpc/video/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListLogic) GetVideoList(in *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
	// todo: add your logic here and delete this line

	return &video.GetVideoListResponse{}, nil
}
