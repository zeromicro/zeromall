package logic

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuickBySmsCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuickBySmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuickBySmsCodeLogic {
	return &QuickBySmsCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuickBySmsCodeLogic) QuickBySmsCode(req *types.UserRegisterReq) (resp *types.UserLoginResp, err error) {
	// todo: add your logic here and delete this line

	return
}
