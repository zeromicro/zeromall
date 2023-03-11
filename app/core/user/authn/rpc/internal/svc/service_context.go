package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/authn/proto/model"
	"user/authn/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	// db model
	MUser    model.UserModel
	MProfile model.UserProfileModel
	MKyc     model.UserKycModel
	MAuthn   model.UserAuthnModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		MUser:    model.NewUserModel(sqlx.NewMysql(c.DataSource), c.Cache),
		MProfile: model.NewUserProfileModel(sqlx.NewMysql(c.DataSource), c.Cache),
		MKyc:     model.NewUserKycModel(sqlx.NewMysql(c.DataSource), c.Cache),
		MAuthn:   model.NewUserAuthnModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
