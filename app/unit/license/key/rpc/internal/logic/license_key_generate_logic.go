package logic

import (
	"context"
	"license/key/proto/model"
	"math/rand"

	"license/key/rpc/internal/svc"
	"license/key/rpc/pb"

	"github.com/better-go/pkg/random"

	"github.com/zeromicro/go-zero/core/logx"
)

type LicenseKeyGenerateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLicenseKeyGenerateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LicenseKeyGenerateLogic {
	return &LicenseKeyGenerateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  注册码批量生成:
func (l *LicenseKeyGenerateLogic) LicenseKeyGenerate(in *pb.LicenseKeyGenerateReq) (*pb.LicenseKeyGenerateResp, error) {
	// todo: add your logic here and delete this line

	req := &model.LicenseKeyVolume{
		Id:        int64(rand.Uint64()),
		Status:    0,
		Desc:      "License",
		PublicKey: "PublicKey-" + random.Gen36BitUUID4(),
		SecretKey: "SecretKey-" + random.Gen20BitUUID(),
	}
	_, err := l.svcCtx.ModelKeyVolume.Insert(l.ctx, req)
	if err != nil {
		return nil, err
	}

	l.Logger.Debugf("LicenseKeyGenerate req=%v", req)
	return &pb.LicenseKeyGenerateResp{
		Status: 1,
	}, nil
}
