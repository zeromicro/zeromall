package demo

import (
	"context"

	"github.com/better-go/pkg/log"

	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/proto/model/graphql"
)

// demo:
type GraphScope struct {
	*dao.MetaResource
}

func newGraphScope(g *dao.MetaResource) *GraphScope {
	return &GraphScope{g}
}

func (m *GraphScope) BlockParse(ctx context.Context) (resp graphql.Query, err error) {
	err = m.Graph.Chain.Block(ctx, &resp)
	log.Info("graphql query resp: %+v", resp)
	return resp, err
}
