package models

type EnergyConsumptionRangeTel struct {
	RangeTel                 uint64  `json:"range"`
	Co2EmissionsTel          uint64  `json:"co2_emissions"`
	VehicleConsumptionTel    uint64  `json:"vehicle_consumption"`
	VehicleFuelEquivalentTel float64 `json:"vehicle_fuel_equivalent"`
}
