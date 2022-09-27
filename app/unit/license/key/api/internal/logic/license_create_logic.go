package logic

import (
	"context"

	"license/key/api/internal/svc"
	"license/key/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseCreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLicenseCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseCreateLogic {
	return &LicenseCreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LicenseCreateLogic) LicenseCreate(req *types.LicenseCreateReq) (resp *types.LicenseCreateResp, err error) {
	// todo: add your logic here and delete this line

	return
}
