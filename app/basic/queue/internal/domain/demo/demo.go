package demo

import (
	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/proto/config"
)

/*
一个业务单元聚合:
	- 聚合一块业务逻辑

*/
type Domain struct {
	// inner global use:
	g *dao.MetaResource

	// biz:
	Hello *HelloScope
	Graph *GraphScope
}

//
func NewDomain(cfg config.Config, isRpcServer bool) *Domain {
	// global:
	g := dao.NewMetaResource(cfg, isRpcServer)

	return &Domain{

		g: g,

		// biz:
		Hello: newAuthScope(g),
		Graph: newGraphScope(g),
	}
}

func (m *Domain) Close() {
	m.g.Close()
	return
}
