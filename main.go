// Code generated by hertz generator.

package main

import (
	"TikTok/biz/dal"
	"TikTok/biz/middleware/jwt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Init() {
	dal.Init()
	jwt.Init()
}

func main() {
	Init()
	h := server.Default(
		server.WithMaxRequestBodySize(10 << 20), // 设置最大请求体大小为 10MB)
	)
	register(h)
	h.Spin()
}
