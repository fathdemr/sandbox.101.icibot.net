package models

type EnergyConsumptionRangeReal struct {
	RangeReal                 uint64  `json:"range"`
	Co2EmissionsReal          uint64  `json:"co2_emissions"`
	VehicleConsumptionReal    uint64  `json:"vehicle_consumption"`
	VehicleFuelEquivalentReal float64 `json:"vehicle_fuel_equivalent"`
}
