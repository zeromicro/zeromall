package graphql

import (
	"github.com/better-go/pkg/net/graphql"

	"mall/app/basic/queue/proto/config"
)

type Dao struct {
	g *ConnGroup

	// biz:
	Chain *ChainGraph
}

func NewDao(cfg *config.GraphQLUnit) *Dao {
	// mq meta:
	g := newConnGroup(cfg)

	return &Dao{
		g: g,

		// biz:
		Chain: newChainGraph(g),
	}
}

func (m *Dao) Close() {
	m.g.Close()
	return
}

//////////////////////////////////////////////////////////////////////////////////////////

// queue conn:
type ConnGroup struct {
	cli *graphql.Client
}

func newConnGroup(cfg *config.GraphQLUnit) *ConnGroup {
	cli := graphql.NewClient(cfg.Url, nil)

	return &ConnGroup{
		cli: cli,
	}
}

func (m *ConnGroup) Close() {
	return
}
