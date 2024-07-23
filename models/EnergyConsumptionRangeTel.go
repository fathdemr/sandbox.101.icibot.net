package models

type EnergyConsumptionRangeTel struct {
	Id                    uint64  `gorm:"primary_key" json:"id"`
	CarId                 uint64  `json:"car_id"`
	Range                 uint64  `json:"range"`
	Co2Emissions          uint64  `json:"co2_emissions"`
	VehicleConsumption    uint64  `json:"vehicle_consumption"`
	VehicleFuelEquivalent float64 `json:"vehicle_fuel_equivalent"`
}
