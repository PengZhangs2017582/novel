package main

import (
	_ "novel/internal/packed"

	// _ "novel/internal/logic/logic.go"

	_ "novel/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"novel/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
