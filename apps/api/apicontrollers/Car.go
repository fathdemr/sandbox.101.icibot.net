package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sandbox.101.icibot.net/Config"
	"sandbox.101.icibot.net/apps/api/services/Carservice"
	"sandbox.101.icibot.net/middlewares"
	"sandbox.101.icibot.net/models"
	"strconv"
)

//err = Database.Db.First(&car, id).Error

func GETCarByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "no record found")
		return
	}

	err = Config.CarService.GetCarById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "no record found")
		return
	}

	c.JSON(http.StatusOK, Config.CarService.CurrentCar)
}

func POSTCalculateRoute(c *gin.Context) {

	var request Carservice.CalculateRootRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	response, err := Config.CarService.CalculateRoot(request)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, response)

}

func POSTCar(c *gin.Context) {

	actor := middlewares.GetUser(c)
	actor.CurrentIpAddress = c.ClientIP()

	var request models.Car
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	service, _ := Carservice.New(Config.Db)
	service.SetActor(actor)
	response, err := service.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, response)
}
