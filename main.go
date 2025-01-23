package main

import (
	"fmt"
	"gin_demo/common/viper"
	"gin_demo/internal/router"
	"gin_demo/internal/svc"
	"github.com/gin-gonic/gin"
)

const _url = "./resource/config.yaml"

func main() {
	config := viper.InitConfig(_url)

	r := gin.New()

	serviceContext := svc.NewServiceContext(config)

	router.InitRouter(r, serviceContext)

	r.Run(fmt.Sprintf(":%d", config.Port))
}
