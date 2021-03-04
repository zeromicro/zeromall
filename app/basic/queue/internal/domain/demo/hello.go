package demo

import (
	"mall/app/basic/queue/internal/dao"
	"mall/app/basic/queue/proto/api"
)

// demo:
type HelloScope struct {
	*dao.MetaResource
}

func newAuthScope(g *dao.MetaResource) *HelloScope {
	return &HelloScope{g}
}

// one api:
func (m *HelloScope) Demo(req api.HelloReq) (*api.Response, error) {
	// todo: add your logic here and delete this line
	return &api.Response{
		Message: "hello " + req.Name,
	}, nil
}
