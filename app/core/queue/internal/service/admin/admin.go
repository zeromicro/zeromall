package admin

import (
	"mall/app/core/queue/internal/domain/queue"
	"mall/app/core/queue/proto/config"
)

/*
	Admin API Service: 提供内部 API

*/
type Service struct {
	d *queue.Domain // 引入业务单元
}

func NewService(cfg config.Config) *Service {
	return &Service{
		d: queue.NewDomain(cfg, false),
	}
}

func (m *Service) Close() {
	m.d.Close()
}
