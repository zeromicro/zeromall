package outer

import (
	"context"
	"mall/app/core/queue/internal/domain/queue"
	"mall/app/core/queue/proto/config"
)

/*
	对外暴露的 API server: HTTP/WebSocket/GraphQL

*/
type Service struct {
	d *queue.Domain // 引入业务单元
}

func NewService(cfg config.Config, ctx context.Context) *Service {
	return &Service{
		d: queue.NewDomain(cfg, false),
	}
}

func (m *Service) Close() {
	m.d.Close()
}
