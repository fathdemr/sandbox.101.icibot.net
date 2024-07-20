package apicontrollers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
