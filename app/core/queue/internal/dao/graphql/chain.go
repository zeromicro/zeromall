package graphql

import (
	"context"

	"mall/app/core/queue/proto/model/graphql"
)

// 链出块分析:
type ChainGraph struct {
	g *ConnGroup
}

func newChainGraph(g *ConnGroup) *ChainGraph {
	return &ChainGraph{g: g}
}

//
func (m *ChainGraph) Block(ctx context.Context, query *graphql.Query) error {
	return m.g.cli.Query(ctx, &query, nil)
}

// market create:
func (m *ChainGraph) CreateMarketEvents(ctx context.Context, query *graphql.QueryCreateMarketEvents) error {
	return m.g.cli.Query(ctx, &query, nil)
}

// QueryFinalizeMarketEvents
func (m *ChainGraph) FinalizeMarketEvents(ctx context.Context, query *graphql.QueryFinalizeMarketEvents) error {
	return m.g.cli.Query(ctx, &query, nil)
}
