package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAuthnOauthModel = (*customUserAuthnOauthModel)(nil)

type (
	// UserAuthnOauthModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthnOauthModel.
	UserAuthnOauthModel interface {
		userAuthnOauthModel
	}

	customUserAuthnOauthModel struct {
		*defaultUserAuthnOauthModel
	}
)

// NewUserAuthnOauthModel returns a model for the database table.
func NewUserAuthnOauthModel(conn sqlx.SqlConn, c cache.CacheConf) UserAuthnOauthModel {
	return &customUserAuthnOauthModel{
		defaultUserAuthnOauthModel: newUserAuthnOauthModel(conn, c),
	}
}
