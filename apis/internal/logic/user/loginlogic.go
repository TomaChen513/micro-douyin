package user

import (
	"context"

	"miniTikTok/apis/internal/svc"
	"miniTikTok/apis/internal/types"
	"miniTikTok/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginReply, err error) {
	// todo: add your logic here and delete this line
	r,err:=l.svcCtx.UserService.CreateUser(l.ctx,&user.CreateUserRequest{
		Name: req.Username,
		Password: req.Password,
	})
	if err!=nil {
		return nil,err
	}
	return &types.LoginReply{UserId: r.Id},nil
}
