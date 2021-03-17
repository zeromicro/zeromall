package db

import (
	"context"

	"mall/app/basic/queue/proto/model"
)

// 用户登录认证:
type HelloStorage struct {
	g *ConnGroup
}

func newHelloStorage(g *ConnGroup) *HelloStorage {
	return &HelloStorage{g: g}
}

func (m *HelloStorage) Hello(ctx context.Context, req *model.HelloReq) (resp *model.HelloEntity, err error) {
	resp = new(model.HelloEntity)

	// todo: need query db here
	err = m.g.DB.DB().Where(req).Find(&resp).Error
	return resp, err
}
