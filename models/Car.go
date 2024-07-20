package models

type Car struct {
	Id                          uint                        `json:"id" gorm:"primary_key" `
	CarModel                    string                      `json:"car_model"`
	Images                      []string                    `json:"images" gorm:"-"`
	AvailableSince              string                      `json:"available_since"`
	UsableBattery               string                      `json:"usable_battery"`
	RealRange                   string                      `json:"real_range"`
	Efficiency                  string                      `json:"efficiency"`
	RealRangeEstimation         RealRangeEstimation         `json:"real_range_estimation" gorm:"-"`
	Performance                 Performance                 `json:"performance" gorm:"-"`
	Battery                     Battery                     `json:"battery" gorm:"-" `
	Charging                    Charging                    `json:"charging" gorm:"-"`
	BidirectionalCharging       BidirectionalCharging       `json:"bidirectional_charging" gorm:"-"`
	EnergyConsumptionRangeReal  EnergyConsumptionRangeReal  `json:"energy_consumption_range_real" gorm:"-"`
	EnergyConsumptionRangeTel   EnergyConsumptionRangeTel   `json:"energy_consumption_range_tel" gorm:"-"`
	EnergyConsumptionRangeTeh   EnergyConsumptionRangeTeh   `json:"energy_consumption_range_teh" gorm:"-"`
	EnergyConsumptionEstimation EnergyConsumptionEstimation `json:"energy_consumption_estimation" gorm:"-"`
	DimensionsAndWeight         DimensionsAndWeight         `json:"dimensions_and_weight" gorm:"-"`
	Miscellaneous               Miscellaneous               `json:"miscellaneous" gorm:"-"`
	ChargeTypes                 []ChargeTypes               `json:"charge_types" gorm:"-"`
}
