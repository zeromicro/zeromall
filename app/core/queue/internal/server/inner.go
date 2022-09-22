package server

import (
	"github.com/better-go/pkg/log"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"

	"mall/app/core/queue/internal/service/inner"
	"mall/app/core/queue/proto/config"
	"mall/app/core/queue/proto/rpc"
)

// internal gRPC server:
type InnerServer struct{}

func (m *InnerServer) Run(configFile string) {
	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)

	// rpc service:
	svc := inner.NewService(cfg)
	defer svc.Close()

	// new rpc engine:
	engine := zrpc.MustNewServer(cfg.Server.Inner.RpcServerConf, func(server *grpc.Server) {
		// register rpc server
		rpc.RegisterQueueServiceServer(server, svc)
	})
	defer engine.Stop()

	log.Infof("starting rpc engine at %v\n", cfg.Server.Inner.RpcServerConf)
	engine.Start()
}
