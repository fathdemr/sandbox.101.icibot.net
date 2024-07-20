package apiroutes

import (
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/pkg/apicontrollers"
)

func RealRangeEstimationRoute(api *gin.RouterGroup) {
	api.GET("/realRange/:id", apicontrollers.RealRangeEstimationController)
}
