package hello

import (
	"gf-app/app/api/response"
	"github.com/gogf/gf/net/ghttp"
)

// Hello is a demonstration route handler for output "Hello World!".
func Hello(r *ghttp.Request) {
	response.Success(r, 200, "HelloWord")
}
