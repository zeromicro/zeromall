package mq

import (
	"github.com/better-go/pkg/log"
	"github.com/better-go/pkg/mq/rabbitmq"
)

// 用户登录认证:
type HelloQueue struct {
	g *ConnGroup
}

func newHelloQueue(g *ConnGroup) *HelloQueue {
	return &HelloQueue{g: g}
}

func (m *HelloQueue) Publish(message string) error {
	err := m.g.cli.Publish(
		&rabbitmq.Exchange{
			Name:       "test-exchange",
			Type:       "direct",
			Durable:    false,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		&rabbitmq.Queue{
			Name:       "test-queue",
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false, // 是否具有排他性
			NoWait:     false, // 是否阻塞
			Args:       nil,   // 额外属性
		},
		"test-routing-key",
		message,
		true,
	)

	log.Infof("producer publish message: %v", err)
	return err
}
