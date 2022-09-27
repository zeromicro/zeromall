package logic

import (
	"context"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyGenerateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyGenerateLogic {
	return &LicenseKeyGenerateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码批量生成:
func (l *LicenseKeyGenerateLogic) LicenseKeyGenerate(in *pb.LicenseKeyGenerateReq) (*pb.LicenseKeyGenerateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LicenseKeyGenerateResp{}, nil
}
