package models

type Car struct {
	Id             uint64   `json:"id,omitempty" gorm:"primary_key" `
	CarModel       string   `json:"car_model,omitempty"`
	Images         []string `json:"images,omitempty" gorm:"-"`
	AvailableSince string   `json:"available_since,omitempty"`
	UsableBattery  string   `json:"usable_battery,omitempty"`
	RealRange      string   `json:"real_range,omitempty"`
	Efficiency     string   `json:"efficiency,omitempty"`
	RealRangeEstimation
	Performance
	Battery
	Charging
	BidirectionalCharging
	EnergyConsumptionRangeReal
	EnergyConsumptionRangeTel
	EnergyConsumptionRangeTeh
	EnergyConsumptionEstimation
	DimensionsAndWeight
	Miscellaneous
	ChargeTypes []ChargeTypes `json:"charge_types,omitempty" gorm:"-"`
	BaseRecordFields
}
