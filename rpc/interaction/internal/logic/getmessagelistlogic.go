package logic

import (
	"context"

	"miniTikTok/rpc/interaction/contact"
	"miniTikTok/rpc/interaction/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageListLogic) GetMessageList(in *contact.GetMessageListRequest) (*contact.GetMessageListResponse, error) {
	// todo: add your logic here and delete this line

	return &contact.GetMessageListResponse{}, nil
}
