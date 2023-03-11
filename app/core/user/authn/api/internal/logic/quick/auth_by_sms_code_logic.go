package quick

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthBySmsCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthBySmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthBySmsCodeLogic {
	return &AuthBySmsCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthBySmsCodeLogic) AuthBySmsCode(req *types.QuickAuthSmsReq) (resp *types.AuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
