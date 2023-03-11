package security

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MfaAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMfaAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MfaAddLogic {
	return &MfaAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MfaAddLogic) MfaAdd(req *types.MfaAddReq) (resp *types.MfaAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
