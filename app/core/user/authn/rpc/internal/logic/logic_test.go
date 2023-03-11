package logic

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"os"
	"testing"
	"user/authn/rpc/internal/config"
	"user/authn/rpc/internal/svc"
)

var (
	// for unit test
	testSvcCtx *svc.ServiceContext
)

func TestMain(m *testing.M) {
	// todo x: 根据需要, 调整相对路径值
	var configFile = flag.String("f", "../../etc/test.yaml", "the test config file")

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	fmt.Printf("config: %+v", c)

	testSvcCtx = svc.NewServiceContext(c)

	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
}
