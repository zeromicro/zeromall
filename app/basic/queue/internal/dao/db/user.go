package db

import (
	"context"

	"mall/app/basic/queue/proto/model"
)

// 用户登录认证:
type UserStorage struct {
	g *ConnGroup
}

func newUserStorage(g *ConnGroup) *UserStorage {
	return &UserStorage{g: g}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (m *UserStorage) CreateUserMeta(ctx context.Context, req *model.UserMeta) (int64, error) {
	ret := m.g.DB.DB().Create(req)
	return ret.RowsAffected, ret.Error
}

func (m *UserStorage) UserMeta(ctx context.Context, userID uint64) (*model.UserMeta, error) {
	resp := new(model.UserMeta)

	err := m.g.DB.DB().Where(&model.UserMeta{UserID: userID}).Find(&resp).Error
	return resp, err
}

func (m *UserStorage) UserMetaWithCache(ctx context.Context, userID uint64) (*model.UserMeta, error) {
	resp := new(model.UserMeta)

	//m.g.DBX.QueryRow()

	err := m.g.DB.DB().Where(&model.UserMeta{UserID: userID}).Find(&resp).Error
	return resp, err
}
