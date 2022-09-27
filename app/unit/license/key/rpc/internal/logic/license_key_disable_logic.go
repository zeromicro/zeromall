package logic

import (
	"context"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyDisableLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyDisableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyDisableLogic {
	return &LicenseKeyDisableLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码封禁:
func (l *LicenseKeyDisableLogic) LicenseKeyDisable(in *pb.Request) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
