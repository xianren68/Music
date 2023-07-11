package router

import (
	"Music/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.New()
	r.Use(gin.Recovery())      // 添加错误恢复中间件
	r.Use(middleware.Logger()) // 添加日志中间件
}
