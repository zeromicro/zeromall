package db

import (
	"github.com/better-go/pkg/database/orm/mysql"
	"github.com/better-go/pkg/log"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/app/core/queue/proto/config"
)

/*
所有 DBX CRUD 操作

*/
type Dao struct {
	g *ConnGroup

	// biz group:
	Hello *HelloStorage // demo db
	User  *UserStorage  //
}

func New(cfg config.Config) *Dao {
	// db meta:
	g := newConnGroup(cfg)

	// biz logic:
	d := &Dao{
		g: g,

		// biz group:
		Hello: newHelloStorage(g),
		User:  newUserStorage(g),
	}
	return d
}

func (m *Dao) Close() {
	m.g.Close()
	return
}

//////////////////////////////////////////////////////////////////////////////////////////

// db conn:
type ConnGroup struct {
	//DBCached *mysql.Client
	//DB2 *mysql.Client
	DB       *mysql.Client
	DBX      sqlx.SqlConn
	DBCached sqlc.CachedConn

	// for close:
	//closer []*mysql.Client
}

// todo:
func newConnGroup(cfg config.Config) *ConnGroup {
	// item:
	db1 := mysql.NewClient(cfg.DB.Demo)
	//db2 := mysql.NewClient(cfg.DB2)

	// redis config:
	redisCfg := make(cache.CacheConf, 0, 0)
	redisCfg = append(redisCfg, cache.NodeConf{
		RedisConf: redis.RedisConf{
			Host: cfg.Cache.Demo.Addr,
			Type: "",
			Pass: "",
		},
	})

	log.Infof("redis config: %+v", redisCfg)

	db := sqlx.NewMysql(cfg.DB.Demo.DSN)
	//conn := sqlc.NewConn(db, redisCfg)

	// sync closer:
	//closer := []*mysql.Client{
	//db1,
	//db2,
	//}

	return &ConnGroup{
		DB:  db1,
		DBX: db,
		//DBCached: conn,
		//DB2: db2,

		// for close:
		//closer: closer,
	}
}

func (m *ConnGroup) Close() {
	defer m.DB.Close()

	// close all:
	//for i, item := range m.closer {
	//	if err := item.Close(); err != nil {
	//		log.Errorf("db[%v] conn close error: %v", i, err)
	//	}
	//}
	return
}
