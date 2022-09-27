package logic

import (
	"context"

	"license/key/api/internal/svc"
	"license/key/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLicenseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseListLogic {
	return &LicenseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LicenseListLogic) LicenseList(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
