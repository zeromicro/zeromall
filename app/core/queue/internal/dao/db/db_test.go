package db

import (
	"os"
	"testing"

	"github.com/better-go/pkg/log"
	"github.com/zeromicro/go-zero/core/conf"

	"mall/app/core/queue/proto/config"
)

var (
	testDao *Dao // for unit test case use
)

func TestMain(m *testing.M) {
	// default config path:
	configFile := "../../../configs/configs.yaml"

	// parse config:
	var cfg config.Config
	conf.MustLoad(configFile, &cfg)
	log.Infof("test dao config: %+v, %+v", cfg.DB, cfg.Cache)

	// new:
	testDao = New(cfg)
	defer testDao.Close()

	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
}
