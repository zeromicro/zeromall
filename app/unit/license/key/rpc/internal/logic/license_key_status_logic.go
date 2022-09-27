package logic

import (
	"context"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyStatusLogic {
	return &LicenseKeyStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码状态: 有效期
func (l *LicenseKeyStatusLogic) LicenseKeyStatus(in *pb.Request) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
