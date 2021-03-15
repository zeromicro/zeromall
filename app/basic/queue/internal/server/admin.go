package server

import (
	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	"mall/app/basic/queue/internal/router"
	"mall/app/basic/queue/internal/service/admin"
	"mall/app/basic/queue/proto/config"
)

// internal admin HTTP/API server:
type AdminServer struct{}

func (m *AdminServer) Run(configFile string) {
	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)

	// admin api:
	svc := admin.NewService(cfg)
	defer svc.Close()

	// new engine:
	engine := rest.MustNewServer(cfg.Server.Admin.RestConf)
	defer engine.Stop()

	// register routes:
	router.RegisterAdminRoutes(engine, svc)

	log.Infof("starting admin api engine at %v\n", cfg.Server.Admin.RestConf)
	engine.Start()
}
