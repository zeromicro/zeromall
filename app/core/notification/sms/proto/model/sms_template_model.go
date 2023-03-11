package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SmsTemplateModel = (*customSmsTemplateModel)(nil)

type (
	// SmsTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSmsTemplateModel.
	SmsTemplateModel interface {
		smsTemplateModel
	}

	customSmsTemplateModel struct {
		*defaultSmsTemplateModel
	}
)

// NewSmsTemplateModel returns a model for the database table.
func NewSmsTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) SmsTemplateModel {
	return &customSmsTemplateModel{
		defaultSmsTemplateModel: newSmsTemplateModel(conn, c),
	}
}
