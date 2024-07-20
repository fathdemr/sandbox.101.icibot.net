package apiroutes

import (
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/pkg/apicontrollers"
)

func CarRoute(api *gin.RouterGroup) {
	api.GET("/car/:id", apicontrollers.GETCarByID)
}
