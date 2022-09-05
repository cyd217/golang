package main

import (
	_ "gf_lua/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf_lua/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
