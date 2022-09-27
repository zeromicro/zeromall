// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	licenseKeyAssignFieldNames          = builder.RawFieldNames(&LicenseKeyAssign{})
	licenseKeyAssignRows                = strings.Join(licenseKeyAssignFieldNames, ",")
	licenseKeyAssignRowsExpectAutoSet   = strings.Join(stringx.Remove(licenseKeyAssignFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	licenseKeyAssignRowsWithPlaceHolder = strings.Join(stringx.Remove(licenseKeyAssignFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheLicenseKeyAssignIdPrefix                  = "cache:licenseKeyAssign:id:"
	cacheLicenseKeyAssignPublicKeyConsumerIdPrefix = "cache:licenseKeyAssign:publicKey:consumerId:"
	cacheLicenseKeyAssignPublicKeyPrefix           = "cache:licenseKeyAssign:publicKey:"
)

type (
	licenseKeyAssignModel interface {
		Insert(ctx context.Context, data *LicenseKeyAssign) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*LicenseKeyAssign, error)
		FindOneByPublicKeyConsumerId(ctx context.Context, publicKey string, consumerId string) (*LicenseKeyAssign, error)
		FindOneByPublicKey(ctx context.Context, publicKey string) (*LicenseKeyAssign, error)
		Update(ctx context.Context, data *LicenseKeyAssign) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLicenseKeyAssignModel struct {
		sqlc.CachedConn
		table string
	}

	LicenseKeyAssign struct {
		Id         int64     `db:"id"`
		CreatedAt  time.Time `db:"created_at"`
		UpdatedAt  time.Time `db:"updated_at"`
		DeletedAt  time.Time `db:"deleted_at"`
		Status     int64     `db:"status"`      // 状态： <0=异常状态, >0=正常状态, 1=已分配, -1=封禁
		Desc       string    `db:"desc"`        // 描述信息
		PublicKey  string    `db:"public_key"`  // 公钥
		ProductId  string    `db:"product_id"`  // 产品id
		ShopId     string    `db:"shop_id"`     // 店铺id
		ConsumerId string    `db:"consumer_id"` // 消费者id
		OrderNo    string    `db:"order_no"`    // 订单号
	}
)

func newLicenseKeyAssignModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultLicenseKeyAssignModel {
	return &defaultLicenseKeyAssignModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`license_key_assign`",
	}
}

func (m *defaultLicenseKeyAssignModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	licenseKeyAssignIdKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignIdPrefix, id)
	licenseKeyAssignPublicKeyConsumerIdKey := fmt.Sprintf("%s%v:%v", cacheLicenseKeyAssignPublicKeyConsumerIdPrefix, data.PublicKey, data.ConsumerId)
	licenseKeyAssignPublicKeyKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignPublicKeyPrefix, data.PublicKey)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, licenseKeyAssignIdKey, licenseKeyAssignPublicKeyConsumerIdKey, licenseKeyAssignPublicKeyKey)
	return err
}

func (m *defaultLicenseKeyAssignModel) FindOne(ctx context.Context, id int64) (*LicenseKeyAssign, error) {
	licenseKeyAssignIdKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignIdPrefix, id)
	var resp LicenseKeyAssign
	err := m.QueryRowCtx(ctx, &resp, licenseKeyAssignIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", licenseKeyAssignRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLicenseKeyAssignModel) FindOneByPublicKeyConsumerId(ctx context.Context, publicKey string, consumerId string) (*LicenseKeyAssign, error) {
	licenseKeyAssignPublicKeyConsumerIdKey := fmt.Sprintf("%s%v:%v", cacheLicenseKeyAssignPublicKeyConsumerIdPrefix, publicKey, consumerId)
	var resp LicenseKeyAssign
	err := m.QueryRowIndexCtx(ctx, &resp, licenseKeyAssignPublicKeyConsumerIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `public_key` = ? and `consumer_id` = ? limit 1", licenseKeyAssignRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, publicKey, consumerId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLicenseKeyAssignModel) FindOneByPublicKey(ctx context.Context, publicKey string) (*LicenseKeyAssign, error) {
	licenseKeyAssignPublicKeyKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignPublicKeyPrefix, publicKey)
	var resp LicenseKeyAssign
	err := m.QueryRowIndexCtx(ctx, &resp, licenseKeyAssignPublicKeyKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `public_key` = ? limit 1", licenseKeyAssignRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, publicKey); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLicenseKeyAssignModel) Insert(ctx context.Context, data *LicenseKeyAssign) (sql.Result, error) {
	licenseKeyAssignIdKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignIdPrefix, data.Id)
	licenseKeyAssignPublicKeyConsumerIdKey := fmt.Sprintf("%s%v:%v", cacheLicenseKeyAssignPublicKeyConsumerIdPrefix, data.PublicKey, data.ConsumerId)
	licenseKeyAssignPublicKeyKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignPublicKeyPrefix, data.PublicKey)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, licenseKeyAssignRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.CreatedAt, data.UpdatedAt, data.DeletedAt, data.Status, data.Desc, data.PublicKey, data.ProductId, data.ShopId, data.ConsumerId, data.OrderNo)
	}, licenseKeyAssignIdKey, licenseKeyAssignPublicKeyConsumerIdKey, licenseKeyAssignPublicKeyKey)
	return ret, err
}

func (m *defaultLicenseKeyAssignModel) Update(ctx context.Context, newData *LicenseKeyAssign) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	licenseKeyAssignIdKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignIdPrefix, data.Id)
	licenseKeyAssignPublicKeyConsumerIdKey := fmt.Sprintf("%s%v:%v", cacheLicenseKeyAssignPublicKeyConsumerIdPrefix, data.PublicKey, data.ConsumerId)
	licenseKeyAssignPublicKeyKey := fmt.Sprintf("%s%v", cacheLicenseKeyAssignPublicKeyPrefix, data.PublicKey)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, licenseKeyAssignRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.CreatedAt, newData.UpdatedAt, newData.DeletedAt, newData.Status, newData.Desc, newData.PublicKey, newData.ProductId, newData.ShopId, newData.ConsumerId, newData.OrderNo, newData.Id)
	}, licenseKeyAssignIdKey, licenseKeyAssignPublicKeyConsumerIdKey, licenseKeyAssignPublicKeyKey)
	return err
}

func (m *defaultLicenseKeyAssignModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLicenseKeyAssignIdPrefix, primary)
}

func (m *defaultLicenseKeyAssignModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", licenseKeyAssignRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLicenseKeyAssignModel) tableName() string {
	return m.table
}
