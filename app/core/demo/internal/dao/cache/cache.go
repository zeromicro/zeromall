package cache

import "demo/proto/config"

/*
所有 Cache Set/Get 操作

*/
type Dao struct {
	g *ConnGroup

	// biz group:
	Hello *HelloCache // demo db
}

func New(cfg *config.CacheUnit) *Dao {
	// db meta:
	g := newConnGroup(cfg)

	// biz logic:
	d := &Dao{
		g: g,

		// biz group:
		Hello: newHelloCache(g),
	}
	return d
}

func (m *Dao) Close() {
	m.g.Close()
	return
}

//////////////////////////////////////////////////////////////////////////////////////////

// redis conn group:
type ConnGroup struct {
	//RDS1 *redis.Client // storage: redis1
	//RDS2 *redis.Client // storage: redis2

	// expire option func:
	//Expire *redis.Expiration

	// for close:
	//closer []*redis.Client
}

func newConnGroup(cfg *config.CacheUnit) *ConnGroup {
	// item:
	//rds1 := redis.NewClient(cfg.R1)
	//rds2 := redis.NewClient(cfg.R2)

	// sync closer item:
	//closer := []*redis.Client{
	//	rds1,
	//	rds2,
	//}

	return &ConnGroup{
		//RDS1: rds1,
		//RDS2: rds2,
		//
		// expire:
		//Expire: redis.NewExpiration(),
		//
		// for close:
		//closer: closer,
	}
}

func (m *ConnGroup) Close() {
	// close all:
	//for i, item := range m.closer {
	//	if err := item.Close(); err != nil {
	//		log.Errorf("redis[%v] conn close error: %v", i, err)
	//	}
	//}
	return
}
