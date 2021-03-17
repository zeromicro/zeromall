package queue

import (
	"context"
	"testing"
)

func TestBlockScope_CreateMarketEvents(t *testing.T) {
	resp, err := testDomain.Block.CreateMarketEvents(context.Background())
	t.Logf("domain test - resp: %+v, err: %v", resp, err)
}
