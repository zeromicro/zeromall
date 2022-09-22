package server

import (
	"context"
	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	"mall/app/core/queue/internal/service/job"
	"mall/app/core/queue/proto/config"
)

// internal job task server:
type JobServer struct{}

func (m *JobServer) Run(configFile string) {
	ctx := context.Background()

	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)

	// new biz service:
	svc := job.NewService(cfg, ctx)
	defer svc.Close()

	// run job:
	svc.RunAsync()

	// new engine:
	engine := rest.MustNewServer(cfg.Server.Job.RestConf)
	defer engine.Stop()

	log.Infof("starting job engine at %v\n", cfg.Server.Job.RestConf)
	engine.Start()
}
