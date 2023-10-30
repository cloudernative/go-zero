package main

import (
	"github.com/cloudernative/go-zero/core/load"
	"github.com/cloudernative/go-zero/core/logx"
	"github.com/cloudernative/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
