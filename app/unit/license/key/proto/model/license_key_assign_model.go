package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseKeyAssignModel = (*customLicenseKeyAssignModel)(nil)

type (
	// LicenseKeyAssignModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseKeyAssignModel.
	LicenseKeyAssignModel interface {
		licenseKeyAssignModel
	}

	customLicenseKeyAssignModel struct {
		*defaultLicenseKeyAssignModel
	}
)

// NewLicenseKeyAssignModel returns a model for the database table.
func NewLicenseKeyAssignModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseKeyAssignModel {
	return &customLicenseKeyAssignModel{
		defaultLicenseKeyAssignModel: newLicenseKeyAssignModel(conn, c),
	}
}
