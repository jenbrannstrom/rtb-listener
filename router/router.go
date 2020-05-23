package router

import (
	"google-rtb/api"
	"google-rtb/config"

	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", api.StatusCheck)
	r.POST("api/rtb", api.RtbListener)

	return r
}

func GetPort() string {
	return config.Cfg.ServerConfigurations.Port
}
