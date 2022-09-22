package cache

import (
	"context"
	"fmt"
)

const (
	// redis key fmt:
	_fmtKeyHello = "auth:%v"
)

type HelloCache struct {
	g *ConnGroup
}

func newHelloCache(g *ConnGroup) *HelloCache {
	return &HelloCache{g: g}
}

//////////////////////////////////////////////////////////////////////////////////////////

// redis key:
func keyHello(userID int64) string {
	return fmt.Sprintf(_fmtKeyHello, userID)
}

//////////////////////////////////////////////////////////////////////////////////////////

// set:
func (m *HelloCache) Set(ctx context.Context, userID int64, token string) error {
	//key := keyHello(userID)
	//expire := m.g.Expire.Minutes(1)

	// todo: need set cache
	return nil
	//return m.g.RDS1.V1().Set(ctx, key, token, expire).Err()
}

// get:
func (m *HelloCache) Get(ctx context.Context, userID int64) (int64, error) {
	//key := keyHello(userID)

	// todo: need query cache
	return 0, nil
	//ret, err := m.g.RDS1.V1().Get(ctx, key).Int64()
	//return ret, err
}
