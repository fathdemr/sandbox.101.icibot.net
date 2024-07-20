package migration

import (
	"fmt"
	"sandbox.101.icibot.net/Database"
	"sandbox.101.icibot.net/models"
	"testing"
)

func TestAutoMigrate(t *testing.T) {

	err := Database.InitDb()
	if err != nil {
		fmt.Println("init db fail")
	}

	err = Database.Db.AutoMigrate(
		&models.Car{},
		&models.Battery{},
		&models.BidirectionalCharging{},
		&models.Charging{},
		&models.ChargeTypes{},
		&models.ChargingPoints{},
		&models.DimensionsAndWeight{},
		&models.EnergyConsumptionEstimation{},
		&models.EnergyConsumptionRangeReal{},
		&models.EnergyConsumptionRangeTeh{},
		&models.EnergyConsumptionRangeTel{},
		&models.Miscellaneous{},
		&models.Performance{},
		&models.RealRangeEstimation{})
	if err != nil {
		t.Error(err)
	}
}
