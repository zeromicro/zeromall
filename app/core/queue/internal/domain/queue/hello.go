package queue

import (
	"mall/app/core/queue/internal/dao"
	"mall/app/core/queue/proto/api"
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

// one api:
func (m *HelloScope) Publish(req api.MessageReq) (resp *api.MessageResp, err error) {
	resp = new(api.MessageResp)

	err = m.MQ.Hello.Publish(req.Data)
	return
}
