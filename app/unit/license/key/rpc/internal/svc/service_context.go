package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"license/key/rpc/internal/config"
)

import "license/key/proto/model"

type ServiceContext struct {
	Config config.Config

	ModelKeyVolume model.LicenseKeyVolumeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		ModelKeyVolume: model.NewLicenseKeyVolumeModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
