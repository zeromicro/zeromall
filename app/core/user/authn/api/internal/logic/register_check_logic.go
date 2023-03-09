package logic

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterCheckLogic {
	return &RegisterCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterCheckLogic) RegisterCheck(req *types.UserRegisterCheckReq) (resp *types.UserRegisterCheckResp, err error) {
	// todo: add your logic here and delete this line

	return
}
