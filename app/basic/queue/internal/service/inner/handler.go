package inner

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/better-go/pkg/log"

	"mall/app/basic/queue/proto/rpc"
)

func (m *Service) Grace(ctx context.Context, req *rpc.Request) (*rpc.Response, error) {
	log.Infof("inner rpc call: %+v", req)

	time.Sleep(time.Millisecond * 10)

	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	return &rpc.Response{
		Host: hostname,
	}, nil
}

func (m *Service) AdminCall(ctx context.Context, req *rpc.AdminReq) (*rpc.AdminResp, error) {
	log.Infof("inner rpc AdminCall req: %+v", req)

	return &rpc.AdminResp{
		Message: fmt.Sprintf("hello from rpc admin api, now = %+v", time.Now()),
		Status:  "ok",
	}, nil
}

func (m *Service) RestCall(ctx context.Context, req *rpc.RestReq) (*rpc.RestResp, error) {
	log.Infof("inner rpc RestCall req: %+v", req)

	return &rpc.RestResp{
		Message: fmt.Sprintf("hello from rpc rest api, now = %+v", time.Now()),
		Status:  "ok",
	}, nil
}
