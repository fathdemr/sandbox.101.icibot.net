package apiroutes

import (
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/apps/api/apicontrollers"
)

func CarRoute(api *gin.RouterGroup) {
	api.GET("/car/:id", apicontrollers.GETCarByID)
	api.POST("/car/calc_route", apicontrollers.POSTCalculateRoute)
	api.POST("/car", apicontrollers.POSTCar)
}
