package viper

import (
	"gin_demo/internal/conf"
	"github.com/spf13/viper"
	"log"
)

func InitConfig(url string) *conf.Conf {
	v := viper.New()
	// 1. 读取配置文件
	v.SetConfigFile(url)
	if err := v.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	config := &conf.Conf{}
	if err := v.Unmarshal(config); err != nil {
		log.Fatal(err)
	}

	return config
}
