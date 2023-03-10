package logic

import (
	"context"

	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Logout 用户退出:
func (l *LogoutLogic) Logout(in *pb.UserLogoutReq) (*pb.UserLogoutResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserLogoutResp{}, nil
}
