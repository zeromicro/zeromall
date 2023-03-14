package quick

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsCodeLogic {
	return &SendSmsCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsCodeLogic) SendSmsCode(req *types.SendSmsCodeReq) (resp *types.SendCodeResp, err error) {
	// todo: add your logic here and delete this line

	/*
		TODO X:
		    1. 参数待补充 图形校验码 captcha code 验证结果.
	*/

	return
}
