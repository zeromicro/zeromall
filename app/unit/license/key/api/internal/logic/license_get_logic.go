package logic

import (
	"context"

	"license/key/api/internal/svc"
	"license/key/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLicenseGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseGetLogic {
	return &LicenseGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LicenseGetLogic) LicenseGet(req *types.LicenseGetReq) (resp *types.LicenseGetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
