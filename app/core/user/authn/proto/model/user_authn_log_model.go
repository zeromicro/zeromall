package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAuthnLogModel = (*customUserAuthnLogModel)(nil)

type (
	// UserAuthnLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthnLogModel.
	UserAuthnLogModel interface {
		userAuthnLogModel
	}

	customUserAuthnLogModel struct {
		*defaultUserAuthnLogModel
	}
)

// NewUserAuthnLogModel returns a model for the database table.
func NewUserAuthnLogModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthnLogModel {
	return &customUserAuthnLogModel{
		defaultUserAuthnLogModel: newUserAuthnLogModel(conn, c),
	}
}
