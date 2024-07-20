package models

type EnergyConsumptionRangeReal struct {
	Id                    uint    `gorm:"primary_key" json:"id"`
	CarId                 uint    `json:"car_id"`
	Range                 uint64  `json:"range"`
	Co2Emissions          uint64  `json:"co2_emissions"`
	VehicleConsumption    uint64  `json:"vehicle_consumption"`
	VehicleFuelEquivalent float64 `json:"vehicle_fuel_equivalent"`
}
