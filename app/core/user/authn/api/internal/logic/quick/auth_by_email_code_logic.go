package quick

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthByEmailCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthByEmailCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByEmailCodeLogic {
	return &AuthByEmailCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthByEmailCodeLogic) AuthByEmailCode(req *types.QuickAuthEmailReq) (resp *types.AuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
