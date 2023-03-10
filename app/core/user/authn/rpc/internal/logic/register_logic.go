package logic

import (
	"context"

	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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

// Register 用户注册:
func (l *RegisterLogic) Register(in *pb.UserRegisterReq) (*pb.UserRegisterResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserRegisterResp{}, nil
}
