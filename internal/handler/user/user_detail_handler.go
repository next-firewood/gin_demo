package user

import (
	"gin_demo/common/response"
	"gin_demo/internal/api"
	"gin_demo/internal/logic/user"
	"gin_demo/internal/svc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserDetailHandler(svcCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &api.UuidForm{}

		if err := c.ShouldBindQuery(req); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		}

		logic := user.NewUserDetailLogic(c.Request.Context(), svcCtx)

		resp, err := logic.UserDetail(req)
		response.Response(c, resp, err)
	}
}
