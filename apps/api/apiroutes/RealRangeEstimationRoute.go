package apiroutes

import (
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/apps/api/apicontrollers"
)

func RealRangeEstimationRoute(api *gin.RouterGroup) {
	api.GET("/realRange/:id", apicontrollers.RealRangeEstimationController)
}
