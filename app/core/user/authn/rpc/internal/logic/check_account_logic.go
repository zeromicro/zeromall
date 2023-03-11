package logic

import (
	"context"

	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAccountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAccountLogic {
	return &CheckAccountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckAccount 帐号状态检测:
func (l *CheckAccountLogic) CheckAccount(in *pb.UserCheckAccountReq) (*pb.UserCheckAccountResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserCheckAccountResp{}, nil
}
