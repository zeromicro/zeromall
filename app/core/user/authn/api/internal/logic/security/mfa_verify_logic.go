package security

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MfaVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMfaVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MfaVerifyLogic {
	return &MfaVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MfaVerifyLogic) MfaVerify(req *types.MfaVerifyReq) (resp *types.MfaVerifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
