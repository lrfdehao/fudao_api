package main

import (
	"crawler_api/dao"
	"crawler_api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Init(r)

	dao.InitDb()
	r.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务
}
