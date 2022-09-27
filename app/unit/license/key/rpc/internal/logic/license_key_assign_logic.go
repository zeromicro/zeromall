package logic

import (
	"context"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyAssignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyAssignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyAssignLogic {
	return &LicenseKeyAssignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码分配:
func (l *LicenseKeyAssignLogic) LicenseKeyAssign(in *pb.LicenseKeyAssignReq) (*pb.LicenseKeyAssignResp, error) {
	// todo: add your logic here and delete this line

	return &pb.LicenseKeyAssignResp{}, nil
}
