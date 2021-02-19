package service

import (
	"mall/app/basic/demo/internal/service/admin"
	"mall/app/basic/demo/internal/service/inner"
	"mall/app/basic/demo/internal/service/job"
	"mall/app/basic/demo/internal/service/outer"
)

/*
	统一对外暴露 Service 入口: 收敛内部服务资源

*/
type Service struct {
	Outer *outer.Server // 对外 HTTP API Server
	Inner *inner.Server // 对内 RPC Server
	Admin *admin.Server // 对内 Admin HTTP API Server
	Job   *job.Server   // 对内 Job Server
}

func NewService() *Service {
	return &Service{
		Outer: outer.NewServer(),
		Inner: inner.NewServer(),
		Admin: admin.NewServer(),
		Job:   job.NewServer(),
	}
}
