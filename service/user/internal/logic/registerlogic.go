package logic

import (
	"context"

	"github.com/palp1tate/go-zero-demo/internal/model"
	"github.com/palp1tate/go-zero-demo/service/user/internal/svc"
	"github.com/palp1tate/go-zero-demo/service/user/pb"

	"github.com/palp1tate/go-crypto-guard/pbkdf2"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	_, err := l.svcCtx.UserModel.FindUserByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "该手机号已注册")
	}
	encodedPassword, err := pwd.GenSHA512(in.Password, 12, 16, 100)
	user := model.User{
		Mobile:   in.Mobile,
		Password: encodedPassword,
		Name:     in.Name,
	}
	if err = l.svcCtx.UserModel.CreateUser(l.ctx, &user); err != nil {
		return nil, status.Error(codes.Internal, "用户注册失败")
	}
	out := &pb.RegisterResponse{UserInfo: &pb.UserInfo{UserID: int64(user.ID), Mobile: user.Mobile, Name: user.Name}}
	return out, nil
}
