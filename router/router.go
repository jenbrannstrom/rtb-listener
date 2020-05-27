package router

import (
	"google-rtb/api"
	"google-rtb/config"

	"github.com/gin-gonic/gin"
)

// GetRouter build all routes
func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", api.StatusCheck)
	r.POST("api/rtb", api.RtbListener)
	r.POST("api/check", api.RtbListenCheck)

	return r
}

// GetPort get port from config
func GetPort() string {
	return config.Cfg.ServerConfigurations.Port
}
