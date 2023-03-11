package logic

import (
	"context"

	"notification/sms/rpc/internal/svc"
	"notification/sms/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Send 发送短信:
func (l *SendLogic) Send(in *pb.SendReq) (*pb.SendResp, error) {
	/*
		todo x: 发送短信流程
			1. 校验参数
			2. 校验短信模板
			3. 发送短信
			4. 记录发送日志
			5. 返回结果
	*/

	// mock
	resp := &pb.SendResp{
		Uuid:   "uuid123546",
		Status: "success",
	}

	// TODO: DO SEND SMS HERE

	l.Logger.Debugf("Send SMS: req: %+v, resp: %+v", in, resp)
	return resp, nil
}
