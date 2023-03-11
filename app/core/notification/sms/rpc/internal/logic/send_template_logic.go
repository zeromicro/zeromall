package logic

import (
	"context"

	"notification/sms/rpc/internal/svc"
	"notification/sms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendTemplateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendTemplateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendTemplateLogic {
	return &SendTemplateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SendTemplate 模板发送:
func (l *SendTemplateLogic) SendTemplate(in *pb.SendTemplateReq) (*pb.SendTemplateResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SendTemplateResp{}, nil
}
