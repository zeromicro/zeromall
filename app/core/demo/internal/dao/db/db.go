package db

import "mall/app/core/demo/proto/config"

/*
所有 DB CRUD 操作

*/
type Dao struct {
	g *ConnGroup

	// biz group:
	Hello *HelloStorage // demo db
}

func New(cfg *config.DBUnit) *Dao {
	// db meta:
	g := newConnGroup(cfg)

	// biz logic:
	d := &Dao{
		g: g,

		// biz group:
		Hello: newHelloStorage(g),
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
	//DB1 *mysql.Client
	//DB2 *mysql.Client

	// for close:
	//closer []*mysql.Client
}

// todo:
func newConnGroup(cfg *config.DBUnit) *ConnGroup {
	// item:
	//db1 := mysql.NewClient(cfg.DB1)
	//db2 := mysql.NewClient(cfg.DB2)

	// sync closer:
	//closer := []*mysql.Client{
	//db1,
	//db2,
	//}

	return &ConnGroup{
		//DB1: db1,
		//DB2: db2,

		// for close:
		//closer: closer,
	}
}

func (m *ConnGroup) Close() {
	// close all:
	//for i, item := range m.closer {
	//	if err := item.Close(); err != nil {
	//		log.Errorf("db[%v] conn close error: %v", i, err)
	//	}
	//}
	return
}
