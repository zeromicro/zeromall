package mq

import (
	"github.com/better-go/pkg/mq/rabbitmq"

	"mall/app/basic/queue/proto/config"
)

type Dao struct {
	g *ConnGroup

	// biz:
	Hello *HelloQueue
}

func New(cfg config.MQUnit) *Dao {
	// mq meta:
	g := newConnGroup(cfg)

	return &Dao{
		g: g,

		// biz:
		Hello: newHelloQueue(g),
	}
}

func (m *Dao) Close() {
	m.g.Close()
	return
}

//////////////////////////////////////////////////////////////////////////////////////////

// queue conn:
type ConnGroup struct {
	cli *rabbitmq.Producer
}

func newConnGroup(cfg config.MQUnit) *ConnGroup {
	cli, _ := rabbitmq.NewProducer(&rabbitmq.ConnOption{
		Uri: cfg.Demo.Url,
	})

	return &ConnGroup{
		cli: cli,
	}
}

func (m *ConnGroup) Close() {
	m.cli.Close()
}
