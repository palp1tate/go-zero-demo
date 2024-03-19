package logic

import (
	"context"

	"github.com/palp1tate/go-zero-demo/service/user/internal/svc"
	"github.com/palp1tate/go-zero-demo/service/user/pb"

	"github.com/palp1tate/go-crypto-guard/pbkdf2"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := l.svcCtx.UserModel.FindUserByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, status.Error(codes.NotFound, "该手机号未注册")
	}
	if ok, _ := pwd.VerifySHA512(in.Password, user.Password); !ok {
		return nil, status.Errorf(codes.Unauthenticated, "密码错误")
	}
	out := &pb.LoginResponse{UserInfo: &pb.UserInfo{UserID: int64(user.ID), Mobile: user.Mobile, Name: user.Name}}
	return out, nil
}
