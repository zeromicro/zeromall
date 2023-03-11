package security

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MfaRemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMfaRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MfaRemoveLogic {
	return &MfaRemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MfaRemoveLogic) MfaRemove(req *types.MfaRemoveReq) (resp *types.MfaRemoveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
