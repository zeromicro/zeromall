package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAuthnModel = (*customUserAuthnModel)(nil)

type (
	// UserAuthnModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthnModel.
	UserAuthnModel interface {
		userAuthnModel
	}

	customUserAuthnModel struct {
		*defaultUserAuthnModel
	}
)

// NewUserAuthnModel returns a model for the database table.
func NewUserAuthnModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthnModel {
	return &customUserAuthnModel{
		defaultUserAuthnModel: newUserAuthnModel(conn, c),
	}
}
