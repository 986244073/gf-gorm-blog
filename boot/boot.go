package boot

import (
	_ "gf-app/packed"
	"github.com/gogf/gf/frame/g"
)

func init() {
	g.Cfg().SetFileName("config.dev.toml")
}
