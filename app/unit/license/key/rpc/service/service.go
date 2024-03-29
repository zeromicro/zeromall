// Code generated by goctl. DO NOT EDIT!
// Source: main.proto

package service

import (
	"context"

	"license/key/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LicenseKeyAssignReq    = pb.LicenseKeyAssignReq
	LicenseKeyAssignResp   = pb.LicenseKeyAssignResp
	LicenseKeyGenerateReq  = pb.LicenseKeyGenerateReq
	LicenseKeyGenerateResp = pb.LicenseKeyGenerateResp
	LicenseKeyGetReq       = pb.LicenseKeyGetReq
	LicenseKeyGetResp      = pb.LicenseKeyGetResp
	Request                = pb.Request
	Response               = pb.Response

	Service interface {
		//  注册码批量生成:
		LicenseKeyGenerate(ctx context.Context, in *LicenseKeyGenerateReq, opts ...grpc.CallOption) (*LicenseKeyGenerateResp, error)
		//  注册码分配:
		LicenseKeyAssign(ctx context.Context, in *LicenseKeyAssignReq, opts ...grpc.CallOption) (*LicenseKeyAssignResp, error)
		//  注册码查询:
		LicenseKeyGet(ctx context.Context, in *LicenseKeyGetReq, opts ...grpc.CallOption) (*LicenseKeyGetResp, error)
		//  注册码批量查询:
		LicenseKeyList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		//  注册码状态: 有效期
		LicenseKeyStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		//  注册码封禁:
		LicenseKeyDisable(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
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

//  注册码批量生成:
func (m *defaultService) LicenseKeyGenerate(ctx context.Context, in *LicenseKeyGenerateReq, opts ...grpc.CallOption) (*LicenseKeyGenerateResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LicenseKeyGenerate(ctx, in, opts...)
}

//  注册码分配:
func (m *defaultService) LicenseKeyAssign(ctx context.Context, in *LicenseKeyAssignReq, opts ...grpc.CallOption) (*LicenseKeyAssignResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LicenseKeyAssign(ctx, in, opts...)
}

//  注册码查询:
func (m *defaultService) LicenseKeyGet(ctx context.Context, in *LicenseKeyGetReq, opts ...grpc.CallOption) (*LicenseKeyGetResp, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LicenseKeyGet(ctx, in, opts...)
}

//  注册码批量查询:
func (m *defaultService) LicenseKeyList(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LicenseKeyList(ctx, in, opts...)
}

//  注册码状态: 有效期
func (m *defaultService) LicenseKeyStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LicenseKeyStatus(ctx, in, opts...)
}

//  注册码封禁:
func (m *defaultService) LicenseKeyDisable(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewServiceClient(m.cli.Conn())
	return client.LicenseKeyDisable(ctx, in, opts...)
}
