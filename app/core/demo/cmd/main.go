package main

import (
	"flag"
	"fmt"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"

	"mall/app/core/demo/internal/dao"
	"mall/app/core/demo/internal/router"
	"mall/app/core/demo/proto/config"
)

// fix for air reload:
var configFile = flag.String("f", "configs/configs.yaml", "the config file")

func main() {
	flag.Parse()

	// parse config:
	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := dao.NewServiceContext(c)

	// new server:
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// register routes:
	router.RegisterHandlers(server, &c, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
