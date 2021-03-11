package main

import (
	"mall/app/basic/queue/internal/server"
	"mall/std/cmd"
)

func main() {
	// auto select:
	cmd.Runner(server.NewServerSelector())
}
