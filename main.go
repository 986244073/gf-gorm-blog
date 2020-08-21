package main

import (
	_ "gf-app/boot"
	_ "gf-app/router"
	"gf-app/utils/db"
	"github.com/gogf/gf/frame/g"
)

func main() {
	db.InitDb()
	g.Server().Run()
}
