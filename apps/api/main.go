package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/Config"
	"sandbox.101.icibot.net/apps/api/apiroutes"
	"sandbox.101.icibot.net/apps/api/services"
	"sandbox.101.icibot.net/middlewares"
)

func main() {

	err := Config.InitDb()
	if err != nil {
		panic(err)
	}

	app := gin.Default()

	// Public routes
	api := app.Group("/")
	apiroutes.RealRangeEstimationRoute(api)
	apiroutes.Ping(api)

	// Protected routes
	protectedRoots := app.Group("/api")
	protectedRoots.Use(middlewares.CheckToken)
	apiroutes.CarRoute(protectedRoots)

	var Heat int64 = 24
	var currentCharge uint64 = 100
	var id uint64 = 364

	maxRange := services.Consumption(id, currentCharge, Heat)

	fmt.Printf("%v şarj ile %v sıcaklıkta %v km yol alınabilir!\", currentCharge, Heat, maxRange")

	fmt.Println(maxRange)

	app.Run(":8094")
}
