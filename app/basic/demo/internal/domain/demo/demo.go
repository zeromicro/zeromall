package demo

import (
	"context"
	svc "mall/app/basic/demo/internal/dao"

	types "mall/app/basic/demo/proto/api"

	"github.com/tal-tech/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) DemoLogic {
	return DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line
	return &types.Response{
		Message: "hello world",
	}, nil
}
