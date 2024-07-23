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
		//&models.ChargeTypes{},
		&models.ChargingPoints{},
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

			if point.AvgPowerKW == 0 {
				if err = Config.Db.Delete(&point).Error; err != nil {
					t.Error(err)
				}
			}

			if point.RateKM == 0 {
				if err = Config.Db.Delete(&point).Error; err != nil {
					t.Error(err)
				}
			}

			wg.Done()

		}(point)
	}

	wg.Wait()

	fmt.Println(len(chargingPoints))

	/*

		var performance []models.Performance

		Config.Db.Find(&performance)

		var wg sync.WaitGroup
		for _, point := range performance {
			wg.Add(1)
			go func(pt models.Performance) {
				Config.Db.Save(&point)

				if point.Acceleration0To100SEC == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.TotalPowerKW == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.TopSpeedKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.ElectricRangeKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				wg.Done()

			}(point)

		}

		wg.Wait()

	*/
	/*

		var realRange []models.RealRangeEstimation

		var wg sync.WaitGroup

		Config.Db.Find(&realRange)

		for _, point := range realRange {
			wg.Add(1)
			go func(point models.RealRangeEstimation) {
				Config.Db.Save(&point)

				if point.CityColdWeatherKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.HighwayColdWeatherKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.CombinedColdWeatherKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.CityMildWeatherKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.HighwayMildWeatherKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				if point.CombinedMildWeatherKM == 0 {
					if err = Config.Db.Delete(&point).Error; err != nil {
						t.Error(err)
					}
				}

				wg.Done()
			}(point)
		}

		wg.Wait()

	*/

}
