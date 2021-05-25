package main

import (
	_ "chart/boot"
	_ "chart/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
