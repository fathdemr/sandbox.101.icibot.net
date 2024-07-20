package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sandbox.101.icibot.net/Database"
	"sandbox.101.icibot.net/pkg/apiroutes"
	"sandbox.101.icibot.net/services"
)

func main() {

	Database.InitDb()

	app := gin.Default()

	api := app.Group("/")
	apiroutes.CarRoute(api)
	apiroutes.RealRangeEstimationRoute(api)
	apiroutes.Ping(api)

	var Heat int64 = 24
	var currentCharge uint64 = 100
	var id uint64 = 364

	maxRange := services.Consumption(id, currentCharge, Heat)

	fmt.Printf("%v şarj ile %v sıcaklıkta %v km yol alınabilir!\", currentCharge, Heat, maxRange")

	fmt.Println(maxRange)

	app.Run(":8094")
}
