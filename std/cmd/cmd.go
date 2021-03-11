package cmd

import (
	"os"
	"strings"

	"github.com/better-go/pkg/log"
	"github.com/urfave/cli/v2"

	"mall/std/proto/server"
)

// run server:
func Runner(inner server.Server, outer server.Server, job server.Server, admin server.Server) {
	// parse cmd args:
	r := &cli.App{
		Name:    "Queue Service",
		Version: "v0.0.1",

		//
		// 参数解析:
		//
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "serverType",
				Aliases: []string{"r"},
				Usage:   "run server type: [inner/outer/job/admin]",
				EnvVars: []string{"SERVER_TYPE", "serverType"},
			},
			&cli.StringFlag{
				Name:    "configFile",
				Aliases: []string{"f"},
				Usage:   "run server type: [inner/outer/job/admin]",
				EnvVars: []string{"CONFIG_FILE", "configFile"},
				//FilePath: "configs/configs.yaml", // 会自动解析文件内容
			},
		},

		//
		// 执行动作:
		//
		Action: func(ctx *cli.Context) error {
			// 服务类型:(小写转换)
			serverType := strings.ToLower(ctx.String("serverType"))
			// 配置文件路径:(可以改成读内容)
			configFile := ctx.String("configFile")

			log.Infof("run action: serverType=%v, configFile=%v", serverType, configFile)

			// dispatch:
			switch serverType {
			case "inner":
				log.Infof("run <internal service> server...")
				inner.Run(configFile)
			case "outer":
				log.Infof("run <interface service> server...")
				outer.Run(configFile)
			case "job":
				log.Infof("run <job service> server...")
				job.Run(configFile)
			case "admin":
				log.Infof("run <job service> server...")
				admin.Run(configFile)
			}
			return nil
		},

		Before:   nil,
		After:    nil,
		Commands: nil,
	}

	log.Infof("ready to start server: cmd.args=%+v", os.Args)

	// do run:
	if err := r.Run(os.Args); err != nil {
		log.Errorf("start server failed, error: %v", err)
	}
}
