package logic

import (
	"context"

	"miniTikTok/rpc/interaction/contact"
	"miniTikTok/rpc/interaction/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLatestMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestMessageLogic {
	return &GetLatestMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLatestMessageLogic) GetLatestMessage(in *contact.GetLatestMessageRequest) (*contact.GetLatestMessageResponse, error) {
	// todo: add your logic here and delete this line

	return &contact.GetLatestMessageResponse{}, nil
}
