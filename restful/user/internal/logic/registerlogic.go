package logic

import (
	"context"
	"net/http"

	"github.com/palp1tate/go-zero-demo/restful/user/internal/svc"
	"github.com/palp1tate/go-zero-demo/restful/user/internal/types"
	"github.com/palp1tate/go-zero-demo/service/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.Register(l.ctx, &pb.RegisterRequest{
		Name:     req.Name,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return
	}
	resp = &types.RegisterResponse{
		Base: types.Base{Code: http.StatusOK, Msg: "注册成功"},
		Data: types.UserInfo{UserID: res.UserInfo.UserID, Mobile: res.UserInfo.Mobile, Name: res.UserInfo.Name},
	}
	return
}
