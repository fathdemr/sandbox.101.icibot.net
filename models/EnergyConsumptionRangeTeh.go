package models

type EnergyConsumptionRangeTeh struct {
	RangeTeh                 uint64  `json:"range"`
	Co2EmissionsTeh          uint64  `json:"co2_emissions"`
	VehicleConsumptionTeh    uint64  `json:"vehicle_consumption"`
	VehicleFuelEquivalentTeh float64 `json:"vehicle_fuel_equivalent"`
}
