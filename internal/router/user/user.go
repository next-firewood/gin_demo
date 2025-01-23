package user

import (
	"gin_demo/internal/handler/user"
	"gin_demo/internal/middleware"
	"gin_demo/internal/svc"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup, svcCtx *svc.ServiceContext) {
	group := router.Group("/user")
	group.Use(middleware.LoggerMiddle()) // 中间件

	group.GET("/detail", user.UserDetailHandler(svcCtx))
}
