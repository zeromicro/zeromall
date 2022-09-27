package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseKeyVolumeModel = (*customLicenseKeyVolumeModel)(nil)

type (
	// LicenseKeyVolumeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseKeyVolumeModel.
	LicenseKeyVolumeModel interface {
		licenseKeyVolumeModel
	}

	customLicenseKeyVolumeModel struct {
		*defaultLicenseKeyVolumeModel
	}
)

// NewLicenseKeyVolumeModel returns a model for the database table.
func NewLicenseKeyVolumeModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseKeyVolumeModel {
	return &customLicenseKeyVolumeModel{
		defaultLicenseKeyVolumeModel: newLicenseKeyVolumeModel(conn, c),
	}
}
