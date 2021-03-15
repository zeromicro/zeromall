package job

import (
	"context"
	"time"

	"github.com/better-go/pkg/cronjob"
	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/collection"

	"mall/app/basic/queue/internal/domain/demo"
	"mall/app/basic/queue/proto/config"
)

/*
	Job 服务: 定时任务/异步任务
		- TimingWheel: https://zeromicro.github.io/go-zero/timing-wheel.html
			- src: github.com/tal-tech/go-zero/core/collection
			- usage: /go-zero@v1.1.4/example/timingwheel/main.go

*/
type Service struct {
	d *demo.Domain // 引入业务单元

	cfg config.JobUnit

	// internal:
	tw        *collection.TimingWheel
	cronJob   *cronjob.CronJob
	taskFuncs []func() // 任务函数列表
}

func NewService(cfg config.Config, ctx context.Context) *Service {
	return &Service{
		d:   demo.NewDomain(cfg, false),
		cfg: cfg.Job,

		//
		cronJob: cronjob.New(),
	}
}

func (m *Service) RunAsync() {
	// register + run:
	m.cronJob.RunAsync(
		cronjob.Task{
			Name:     "query and parse block",
			Schedule: m.cfg.Schedule.QueryBlock,
			TaskFunc: func() {
				m.taskQueryBlock()
			},
		},
	)
}

func (m *Service) RunTwTask() (err error) {
	interval := time.Duration(m.cfg.IntervalMinute) * time.Second
	m.tw, err = collection.NewTimingWheel(
		interval,
		m.cfg.SlotNum,
		func(key, value interface{}) {
			m.taskQueryBlock()
		},
	)
	if err != nil {
		log.Errorf("cronJob - register task error: %v", err)
		return
	}
	defer m.tw.Stop()

	//
	for i := 0; ; i++ {
		m.tw.SetTimer(i, i, interval)
		time.Sleep(time.Millisecond)
	}
}

// task1:
func (m *Service) taskQueryBlock() {
	log.Infof("post graphql server, query block info, every 12s")
	time.Sleep(1 * time.Second)
}

func (m *Service) Close() {
	m.cronJob.Stop()
}
