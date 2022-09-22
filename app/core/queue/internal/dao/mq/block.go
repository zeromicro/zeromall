package mq

import (
	"github.com/better-go/pkg/mq/rabbitmq"
)

// block:
type BlockQueue struct {
	g *ConnGroup
}

func newBlockQueue(g *ConnGroup) *BlockQueue {
	return &BlockQueue{g: g}
}

// CreateMarketEvents
func (m *BlockQueue) PublishCreateMarketEvents(message string) error {
	queueName := "CreateMarketEventsQueue"
	return m.publish(message, queueName)
}

// FinalizeMarketEvents
func (m *BlockQueue) PublishFinalizeMarketEvents(message string) error {
	queueName := "FinalizeMarketEventsQueue"
	return m.publish(message, queueName)
}

// pub:
func (m *BlockQueue) publish(message string, queueName string) error {
	exchangeName := "BlockExchange"
	return m.g.cli.Publish(
		&rabbitmq.Exchange{
			Name:       exchangeName,
			Type:       "direct",
			Durable:    false,
			AutoDelete: false,
			Internal:   false,
			NoWait:     false,
			Args:       nil,
		},
		&rabbitmq.Queue{
			Name:       queueName,
			Durable:    true,
			AutoDelete: false,
			Exclusive:  false, // 是否具有排他性
			NoWait:     false, // 是否阻塞
			Args:       nil,   // 额外属性
		},
		queueName,
		message,
		true,
	)
}
