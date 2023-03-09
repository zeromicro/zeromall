package logic

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuickByEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuickByEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuickByEmailCodeLogic {
	return &QuickByEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuickByEmailCodeLogic) QuickByEmailCode(req *types.UserRegisterReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
