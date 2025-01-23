package router

import (
	"gin_demo/internal/router/user"
	"gin_demo/internal/svc"
	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine, svcCtx *svc.ServiceContext) {
	globalRouter := router.Group("/api")

	user.Router(globalRouter, svcCtx)
}
