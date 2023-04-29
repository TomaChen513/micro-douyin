package video

import (
	"context"

	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentVideoLogic {
	return &CommentVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentVideoLogic) CommentVideo(req *types.CommentVideoRequest) (resp *types.CommentVideoReply, err error) {
	// todo: add your logic here and delete this line

	return
}
