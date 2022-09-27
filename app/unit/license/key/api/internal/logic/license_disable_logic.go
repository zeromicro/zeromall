package logic

import (
	"context"

	"license/key/api/internal/svc"
	"license/key/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseDisableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLicenseDisableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseDisableLogic {
	return &LicenseDisableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LicenseDisableLogic) LicenseDisable(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
