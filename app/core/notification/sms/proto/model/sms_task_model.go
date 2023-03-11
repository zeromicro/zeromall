package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SmsTaskModel = (*customSmsTaskModel)(nil)

type (
	// SmsTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsTaskModel.
	SmsTaskModel interface {
		smsTaskModel
	}

	customSmsTaskModel struct {
		*defaultSmsTaskModel
	}
)

// NewSmsTaskModel returns a model for the database table.
func NewSmsTaskModel(conn sqlx.SqlConn, c cache.CacheConf) SmsTaskModel {
	return &customSmsTaskModel{
		defaultSmsTaskModel: newSmsTaskModel(conn, c),
	}
}
