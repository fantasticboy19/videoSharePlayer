package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	s := gin.Default()

	// 跨域请求允许
	// 一般只需要在服务器端设置
	s.Use(cors.Middleware(cors.Config{
		ValidateHeaders: false,
		Origins:         "*",
		RequestHeaders:  "",
		ExposedHeaders:  "",
		Methods:         "",
		MaxAge:          0,
		Credentials:     false,
	}))

	// 开启静态文件托管
	// videos目录下是 /video1/我们的转化之后的文件
	s.Static("/videos/","videos")
	s.Static("/htmls/", "htmls")
	_ = s.Run(":8000")
}
