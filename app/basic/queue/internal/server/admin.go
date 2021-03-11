package server

import (
	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/internal/router"
	"mall/app/basic/queue/proto/config"
)

// internal admin HTTP/API server:
type AdminServer struct {
}

func (m *AdminServer) Run(configFile string) {
	// parse config:
	var c config.Config
	conf.MustLoad(configFile, &c)

	ctx := dao.NewServiceContext(c)

	// new server:
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// register routes:
	router.RegisterHandlers(server, &c, ctx)

	log.Infof("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
