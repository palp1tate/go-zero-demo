package logic

import (
	"context"
	"github.com/palp1tate/go-zero-demo/restful/user/internal/svc"
	"github.com/palp1tate/go-zero-demo/restful/user/internal/types"
	"github.com/palp1tate/go-zero-demo/service/user/pb"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &pb.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	resp = &types.LoginResponse{
		Base: types.Base{Code: 200, Msg: "登录成功"},
		Data: types.UserInfo{UserID: res.UserInfo.UserID, Mobile: res.UserInfo.Mobile, Name: res.UserInfo.Name},
	}
	return
}
