package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LicenseKeyAttrModel = (*customLicenseKeyAttrModel)(nil)

type (
	// LicenseKeyAttrModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLicenseKeyAttrModel.
	LicenseKeyAttrModel interface {
		licenseKeyAttrModel
	}

	customLicenseKeyAttrModel struct {
		*defaultLicenseKeyAttrModel
	}
)

// NewLicenseKeyAttrModel returns a model for the database table.
func NewLicenseKeyAttrModel(conn sqlx.SqlConn, c cache.CacheConf) LicenseKeyAttrModel {
	return &customLicenseKeyAttrModel{
		defaultLicenseKeyAttrModel: newLicenseKeyAttrModel(conn, c),
	}
}
