package tests

import (
	"fmt"
	"sandbox.101.icibot.net/Config"
	"sandbox.101.icibot.net/models"
	"sync"
	"testing"
)

func TestAutoMigrate(t *testing.T) {

	err := Config.InitDb()
	if err != nil {
		panic(err)
	}

	err = Config.Db.AutoMigrate(
		//&models.Car{},
		//&models.Battery{},
		//&models.BidirectionalCharging{},
		//&models.Charging{},
		//&models.ChargeTypes{},
		&models.ChargingPoints{},
		//&models.DimensionsAndWeight{},
		//&models.EnergyConsumptionEstimation{},
		//&models.EnergyConsumptionRangeReal{},
		//&models.EnergyConsumptionRangeTeh{},
		//&models.EnergyConsumptionRangeTel{},
		//&models.Miscellaneous{},
		//&models.Performance{},
		//&models.RealRangeEstimation{},
		//&models.Station{},
		//&models.User{},
	)
	if err != nil {
		t.Error(err)
	}

	var chargingPoints []models.ChargingPoints

	if err = Config.Db.Where("time = ?", "").Delete(&chargingPoints).Error; err != nil {
		t.Error(err)

	}

	Config.Db.Find(&chargingPoints)

	var wg sync.WaitGroup

	for _, point := range chargingPoints {
		wg.Add(1)
		go func(point models.ChargingPoints) {
			Config.Db.Save(&point)

			if point.ChargeMinutes == 0 {
				if err = Config.Db.Delete(&point).Error; err != nil {
					t.Error(err)
				}
			}

			wg.Done()

		}(point)
	}

	wg.Wait()
	fmt.Println(len(chargingPoints))

}
