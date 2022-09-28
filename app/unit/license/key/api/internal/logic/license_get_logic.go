package logic

import (
	"context"

	"license/key/api/internal/svc"
	"license/key/api/internal/types"
	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLicenseGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseGetLogic {
	return &LicenseGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LicenseGetLogic) LicenseGet(req *types.LicenseGetReq) (resp *types.LicenseGetResp, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.LicenseGetResp)

	ret, rErr := l.svcCtx.KeyRpc.LicenseKeyGet(l.ctx, &pb.LicenseKeyGetReq{
		PublicKey: "test input public key",
	})

	resp.PublicKey = ret.PublicKey
	resp.SecretKey = ret.SecretKey
	l.Logger.Infof("LicenseGet: req=%+v, ret=%+v, err=%v", req, ret, err)
	return resp, rErr
}
