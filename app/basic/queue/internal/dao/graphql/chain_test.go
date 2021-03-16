package graphql

import (
	"context"
	"testing"

	"mall/app/basic/queue/proto/model/graphql"
)

func TestChainGraph_Block(t *testing.T) {
	var qs graphql.Query

	err := testDao.Chain.Block(context.Background(), &qs)
	t.Logf("query resp: %v, err: %v", qs, err)
}
