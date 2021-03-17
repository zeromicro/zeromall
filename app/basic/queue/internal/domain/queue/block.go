package queue

import (
	"context"

	"github.com/better-go/pkg/log"

	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/proto/model/graphql"
)

// demo:
type BlockScope struct {
	*dao.MetaResource
}

func newBlockScope(g *dao.MetaResource) *BlockScope {
	return &BlockScope{g}
}

func (m *BlockScope) BlockParse(ctx context.Context) (resp graphql.Query, err error) {
	err = m.Graph.Chain.Block(ctx, &resp)
	log.Info("graphql query resp: %+v", resp)
	return resp, err
}

func (m *BlockScope) CreateMarketEvents(ctx context.Context) (resp graphql.QueryCreateMarketEvents, err error) {
	// graphql query:
	err = m.Graph.Chain.CreateMarketEvents(ctx, &resp)

	// json string:
	msg, err := resp.ToJsonString()
	if err != nil {
		log.Errorf("CreateMarketEvents invalid json format, %+v", resp)
		return
	}

	// queue pub:
	err = m.MQ.Block.PublishCreateMarketEvents(msg)
	return resp, err
}

func (m *BlockScope) FinalizeMarketEvents(ctx context.Context) (resp graphql.QueryFinalizeMarketEvents, err error) {
	err = m.Graph.Chain.FinalizeMarketEvents(ctx, &resp)
	log.Info("graphql query resp: %+v", resp)
	return resp, err
}
