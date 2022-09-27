package logic

import (
	"context"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyListLogic {
	return &LicenseKeyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码批量查询:
func (l *LicenseKeyListLogic) LicenseKeyList(in *pb.Request) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
