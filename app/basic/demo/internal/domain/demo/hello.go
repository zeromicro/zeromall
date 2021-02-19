package demo

import (
	"mall/app/basic/demo/internal/dao"
	"mall/app/basic/demo/proto/api"
)

// demo:
type HelloScope struct {
	*dao.MetaResource
}

func newAuthScope(g *dao.MetaResource) *HelloScope {
	return &HelloScope{g}
}

// one api:
func (m *HelloScope) Demo(req api.Request) (*api.Response, error) {
	// todo: add your logic here and delete this line
	return &api.Response{
		Message: "hello world",
	}, nil
}
