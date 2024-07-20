package apiroutes

import (
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/apps/api/apicontrollers"
)

func Ping(api *gin.RouterGroup) {
	api.GET("/ping", apicontrollers.PingController)
}
