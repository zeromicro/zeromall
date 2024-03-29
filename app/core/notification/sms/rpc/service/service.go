// Code generated by goctl. DO NOT EDIT.
// Source: main.proto

package service

import (
	"context"

	"notification/sms/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CheckMobileStatusReq  = pb.CheckMobileStatusReq
	CheckMobileStatusResp = pb.CheckMobileStatusResp
	GetTemplateReq        = pb.GetTemplateReq
	GetTemplateResp       = pb.GetTemplateResp
	SendBatchReq          = pb.SendBatchReq
	SendBatchResp         = pb.SendBatchResp
	SendReq               = pb.SendReq
	SendResp              = pb.SendResp
	SendTemplateReq       = pb.SendTemplateReq
	SendTemplateResp      = pb.SendTemplateResp
	UserMobileEntity      = pb.UserMobileEntity

	Service interface {
		// Send 发送短信:
		Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendResp, error)
		// SendTemplate 模板发送:
		SendTemplate(ctx context.Context, in *SendTemplateReq, opts ...grpc.CallOption) (*SendTemplateResp, error)
		// GetTemplate 查询模板
		GetTemplate(ctx context.Context, in *GetTemplateReq, opts ...grpc.CallOption) (*GetTemplateResp, error)
		// SendBatch 批量发送:
		SendBatch(ctx context.Context, in *SendBatchReq, opts ...grpc.CallOption) (*SendBatchResp, error)
		// CheckMobileStatus 检查手机号码状态:
		CheckMobileStatus(ctx context.Context, in *CheckMobileStatusReq, opts ...grpc.CallOption) (*CheckMobileStatusResp, error)
	}

	defaultService struct {
		cli zrpc.Client
	}
)

func NewService(cli zrpc.Client) Service {
	return &defaultService{
		cli: cli,
	}
}

// Send 发送短信:
func (m *defaultService) Send(ctx context.Context, in *SendReq, opts ...grpc.CallOption) (*SendResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.Send(ctx, in, opts...)
}

// SendTemplate 模板发送:
func (m *defaultService) SendTemplate(ctx context.Context, in *SendTemplateReq, opts ...grpc.CallOption) (*SendTemplateResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.SendTemplate(ctx, in, opts...)
}

// GetTemplate 查询模板
func (m *defaultService) GetTemplate(ctx context.Context, in *GetTemplateReq, opts ...grpc.CallOption) (*GetTemplateResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.GetTemplate(ctx, in, opts...)
}

// SendBatch 批量发送:
func (m *defaultService) SendBatch(ctx context.Context, in *SendBatchReq, opts ...grpc.CallOption) (*SendBatchResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.SendBatch(ctx, in, opts...)
}

// CheckMobileStatus 检查手机号码状态:
func (m *defaultService) CheckMobileStatus(ctx context.Context, in *CheckMobileStatusReq, opts ...grpc.CallOption) (*CheckMobileStatusResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.CheckMobileStatus(ctx, in, opts...)
}
