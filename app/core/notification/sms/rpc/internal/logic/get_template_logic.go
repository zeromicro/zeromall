package logic

import (
	"context"

	"notification/sms/rpc/internal/svc"
	"notification/sms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplateLogic {
	return &GetTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetTemplate 查询模板
func (l *GetTemplateLogic) GetTemplate(in *pb.GetTemplateReq) (*pb.GetTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTemplateResp{}, nil
}
