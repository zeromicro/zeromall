package job

import (
	"context"
	"time"

	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/collection"

	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/internal/domain/demo"
	"mall/app/basic/queue/proto/config"
)

/*
	Job 服务: 定时任务/异步任务
		- TimingWheel: https://zeromicro.github.io/go-zero/timing-wheel.html
			- src: github.com/tal-tech/go-zero/core/collection
			- usage: /go-zero@v1.1.4/example/timingwheel/main.go

*/
type Server struct {
	d *demo.Domain // 引入业务单元

	cfg *config.JobUnit
	tw  *collection.TimingWheel
}

func NewServer(cfg *config.Config, ctx *dao.ServiceContext) *Server {
	return &Server{
		// todo: need fix
		d:   demo.NewDomain(cfg, context.TODO(), ctx),
		cfg: cfg.Job,
	}
}

func (m *Server) RunTask() (err error) {
	interval := time.Duration(m.cfg.IntervalMinute) * time.Second
	m.tw, err = collection.NewTimingWheel(
		interval,
		m.cfg.SlotNum,
		func(key, value interface{}) {
			m.TaskQueryBlock()
		},
	)
	if err != nil {
		log.Errorf("job - register task error: %v", err)
		return
	}
	defer m.tw.Stop()

	//
	for i := 0; ; i++ {
		m.tw.SetTimer(i, i, interval)
		time.Sleep(time.Millisecond)
	}
}

func (m *Server) TaskQueryBlock() {
	log.Infof("post graphql server, query block info, every 12s")
	time.Sleep(1 * time.Second)
}
