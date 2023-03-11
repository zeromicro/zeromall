package logic

import (
	"context"

	"notification/sms/rpc/internal/svc"
	"notification/sms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckMobileStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckMobileStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckMobileStatusLogic {
	return &CheckMobileStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CheckMobileStatus 检查手机号码状态:
func (l *CheckMobileStatusLogic) CheckMobileStatus(in *pb.CheckMobileStatusReq) (*pb.CheckMobileStatusResp, error) {
	// todo: add your logic here and delete this line

	return &pb.CheckMobileStatusResp{}, nil
}
