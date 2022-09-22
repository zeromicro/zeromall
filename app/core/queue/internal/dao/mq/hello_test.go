package mq

import (
	"fmt"
	"testing"
	"time"
)

func TestHelloQueue_Publish(t *testing.T) {
	// required start rabbitmq before
	msg := fmt.Sprintf("unit test pub msg: %v", time.Now())
	err := testDao.Hello.Publish(msg)
	t.Logf("mq publish msg: %v, error: %v", msg, err)
}
