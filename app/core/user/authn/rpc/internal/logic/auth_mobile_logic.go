package logic

import (
	"context"

	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthMobileLogic {
	return &AuthMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AuthMobile 移动端: 鉴权 + token
func (l *AuthMobileLogic) AuthMobile(in *pb.AuthMobileReq) (*pb.AuthMobileResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AuthMobileResp{}, nil
}
