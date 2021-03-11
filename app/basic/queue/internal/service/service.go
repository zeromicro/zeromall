package service

import (
	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/internal/service/admin"
	"mall/app/basic/queue/internal/service/inner"
	"mall/app/basic/queue/internal/service/job"
	"mall/app/basic/queue/internal/service/outer"
	"mall/app/basic/queue/proto/config"
)

/*
	统一对外暴露 Service 入口: 收敛内部服务资源

*/
type Service struct {
	// common:
	cfg    *config.Config
	srvCtx *dao.ServiceContext

	// biz:
	Outer *outer.Server // 对外 HTTP API Server
	Inner *inner.Server // 对内 RPC Server
	Admin *admin.Server // 对内 Admin HTTP API Server
	Job   *job.Server   // 对内 Job Server
}

func NewService(cfg *config.Config, srvCtx *dao.ServiceContext) *Service {
	return &Service{
		cfg:    cfg,
		srvCtx: srvCtx,

		// biz:
		Outer: outer.NewServer(cfg, srvCtx),
		Inner: inner.NewServer(),
		Admin: admin.NewServer(),
		Job:   job.NewServer(cfg, srvCtx),
	}
}
