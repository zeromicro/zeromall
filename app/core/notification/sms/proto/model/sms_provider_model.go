package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SmsProviderModel = (*customSmsProviderModel)(nil)

type (
	// SmsProviderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsProviderModel.
	SmsProviderModel interface {
		smsProviderModel
	}

	customSmsProviderModel struct {
		*defaultSmsProviderModel
	}
)

// NewSmsProviderModel returns a model for the database table.
func NewSmsProviderModel(conn sqlx.SqlConn, c cache.CacheConf) SmsProviderModel {
	return &customSmsProviderModel{
		defaultSmsProviderModel: newSmsProviderModel(conn, c),
	}
}
