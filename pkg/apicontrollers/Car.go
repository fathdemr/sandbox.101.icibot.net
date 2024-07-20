package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sandbox.101.icibot.net/Database"
	"sandbox.101.icibot.net/models"
	"sandbox.101.icibot.net/services/Carservice"
	"strconv"
)

//err = Database.Db.First(&car, id).Error

func GETCarByID(c *gin.Context) {

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "no record found")
		return
	}

	var car models.Car
	s := Carservice.New(Database.Db)
	car, err = s.GetCarById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "no record found")
		return
	}

	c.JSON(http.StatusOK, car)
}
