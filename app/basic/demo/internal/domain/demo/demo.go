package demo

import (
	"context"
	"mall/app/basic/demo/proto/config"

	"github.com/tal-tech/go-zero/core/logx"

	"mall/app/basic/demo/internal/dao"
)

/*
一个业务单元聚合:
	- 聚合一块业务逻辑

*/
type Domain struct {
	logx.Logger

	ctx    context.Context
	svcCtx *dao.ServiceContext

	///////////////////////////////////////

	// inner global use:
	g *dao.MetaResource

	// biz:
	Hello *HelloScope
}

//
func NewScope(cfg *config.Config, ctx context.Context, svcCtx *dao.ServiceContext) *Domain {
	// global:
	g := dao.NewMetaResource(cfg)

	return &Domain{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,

		g: g,

		// biz:
		Hello: newAuthScope(g),
	}
}

func (m *Domain) Close() {
	m.g.Close()
	return
}
