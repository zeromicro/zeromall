package register

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAccountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckAccountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAccountLogic {
	return &CheckAccountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckAccountLogic) CheckAccount(req *types.UserRegisterCheckReq) (resp *types.UserRegisterCheckResp, err error) {
	// todo: add your logic here and delete this line

	return
}
