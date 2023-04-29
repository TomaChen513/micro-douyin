package logic

import (
	"context"

	"miniTikTok/rpc/user/internal/svc"
	"miniTikTok/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *user.CreateUserRequest) (*user.CreatUserReply, error) {
	// todo: add your logic here and delete this line

	return &user.CreatUserReply{}, nil
}
