package logic

import (
	"context"

	"notification/sms/rpc/internal/svc"
	"notification/sms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendBatchLogic {
	return &SendBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// SendBatch 批量发送:
func (l *SendBatchLogic) SendBatch(in *pb.SendBatchReq) (*pb.SendBatchResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SendBatchResp{}, nil
}
