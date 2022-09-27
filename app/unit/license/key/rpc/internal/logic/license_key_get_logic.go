package logic

import (
	"context"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyGetLogic {
	return &LicenseKeyGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码查询:
func (l *LicenseKeyGetLogic) LicenseKeyGet(in *pb.LicenseKeyGetReq) (*pb.LicenseKeyGetResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LicenseKeyGetResp{}, nil
}
