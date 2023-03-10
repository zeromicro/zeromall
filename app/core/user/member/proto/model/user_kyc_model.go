package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserKycModel = (*customUserKycModel)(nil)

type (
	// UserKycModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserKycModel.
	UserKycModel interface {
		userKycModel
	}

	customUserKycModel struct {
		*defaultUserKycModel
	}
)

// NewUserKycModel returns a model for the database table.
func NewUserKycModel(conn sqlx.SqlConn, c cache.CacheConf) UserKycModel {
	return &customUserKycModel{
		defaultUserKycModel: newUserKycModel(conn, c),
	}
}
