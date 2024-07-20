package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sandbox.101.icibot.net/Database"
	"sandbox.101.icibot.net/models"
	"strconv"
)

func RealRangeEstimationController(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "no record found")
		return
	}

	var car models.Car
	err = Database.Db.First(&car, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, "no record found")
		return
	}

	Database.Db.Where("car_id = ?", id).First(&car.RealRangeEstimation)

	c.JSON(http.StatusOK, car)
}
