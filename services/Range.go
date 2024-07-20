package services

import (
	"fmt"
	"sandbox.101.icibot.net/Database"
)

var sql = `
		select c.id,
			   c.car_model,
				b.usable_capacity,
				e.vehicle_consumption
		from cars c
				 left join batteries b on c.id = b.car_id
		left join energy_consumption_range_reals e on c.id = e.car_id
		Where c.id = ?;
`

type ConsResponse struct {
	ID                 uint64  `json:"id" gorm:"primary_key"`
	CarModel           string  `json:"car_model"`
	UsableCapacity     float64 `json:"usable_capacity"`
	VehicleConsumption uint64  `json:"vehicle_consumption"`
}

type ChargeStations struct {
	Distance float64 `json:"distance"`
}

var chargeStations = []ChargeStations{
	ChargeStations{
		Distance: 200,
	},
	ChargeStations{
		Distance: 300,
	},
	ChargeStations{
		400,
	},
}

func airConsumption(Heat int64, Range float64) (maxRange float64) {
	switch Heat {
	case 24:
	case 27:
		Range = Range - (Range * 3 / 100)
	case 29:
		Range = Range - (Range * 3 / 100)
	case 32:
		Range = Range - (Range * 5 * 100)
	case 35:
		Range = Range - (Range * 15 / 100)
	case 38:
		Range = Range - (Range * 31 / 100)
	}
	return Range
}

func Consumption(id uint64, currentCharge uint64, Heat int64) (maxRange float64) {

	var result ConsResponse = ConsResponse{}

	err := Database.Db.
		Raw(sql, id).
		Scan(&result).Error

	if err != nil {
		fmt.Println(err)
	}

	var vehicleConsumption = float64(result.VehicleConsumption)

	var usableCapacity = result.UsableCapacity

	maxRange = (float64(currentCharge) * usableCapacity * 10) / vehicleConsumption

	maxRange = airConsumption(Heat, maxRange)

	return maxRange
}

func ChargeStationCalc(id uint64) {
	maxRange := Consumption(id, 100, 24)

}
