package quick

import (
	"context"

	"user/authn/api/internal/svc"
	"user/authn/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthByOauthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthByOauthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthByOauthLogic {
	return &AuthByOauthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthByOauthLogic) AuthByOauth(req *types.QuickAuthOAuthReq) (resp *types.AuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
