package logic

import (
	"context"

	"user/authn/rpc/internal/svc"
	"user/authn/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthWebLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthWebLogic {
	return &AuthWebLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AuthWeb web端: 鉴权 + cookie
func (l *AuthWebLogic) AuthWeb(in *pb.AuthWebReq) (*pb.AuthWebResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AuthWebResp{}, nil
}
