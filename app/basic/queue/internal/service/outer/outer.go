package outer

import (
	"context"
	"mall/app/basic/queue/internal/domain/demo"
	"mall/app/basic/queue/proto/config"
)

/*
	对外暴露的 API server: HTTP/WebSocket/GraphQL

*/
type Service struct {
	d *demo.Domain // 引入业务单元
}

func NewService(cfg config.Config, ctx context.Context) *Service {
	return &Service{
		d: demo.NewDomain(cfg, false),
	}
}

func (m *Service) Close() {
	m.d.Close()
}
