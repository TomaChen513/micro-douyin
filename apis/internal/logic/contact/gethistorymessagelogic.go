package contact

import (
	"context"

	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHistoryMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHistoryMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHistoryMessageLogic {
	return &GetHistoryMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHistoryMessageLogic) GetHistoryMessage(req *types.GetHistoryMessageRequest) (resp *types.GetHistoryMessageReply, err error) {
	// todo: add your logic here and delete this line

	return
}
