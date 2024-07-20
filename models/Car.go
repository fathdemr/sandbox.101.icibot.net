package models

type Car struct {
	Id                          uint64                       `json:"id,omitempty" gorm:"primary_key" `
	CarModel                    string                       `json:"car_model,omitempty"`
	Images                      []string                     `json:"images,omitempty" gorm:"-"`
	AvailableSince              string                       `json:"available_since,omitempty"`
	UsableBattery               string                       `json:"usable_battery,omitempty"`
	RealRange                   string                       `json:"real_range,omitempty"`
	Efficiency                  string                       `json:"efficiency,omitempty"`
	RealRangeEstimation         *RealRangeEstimation         `json:"real_range_estimation,omitempty" gorm:"-"`
	Performance                 *Performance                 `json:"performance,omitempty" gorm:"-"`
	Battery                     *Battery                     `json:"battery,omitempty" gorm:"-" `
	Charging                    *Charging                    `json:"charging,omitempty" gorm:"-"`
	BidirectionalCharging       *BidirectionalCharging       `json:"bidirectional_charging,omitempty" gorm:"-"`
	EnergyConsumptionRangeReal  *EnergyConsumptionRangeReal  `json:"energy_consumption_range_real,omitempty" gorm:"-"`
	EnergyConsumptionRangeTel   *EnergyConsumptionRangeTel   `json:"energy_consumption_range_tel,omitempty" gorm:"-"`
	EnergyConsumptionRangeTeh   *EnergyConsumptionRangeTeh   `json:"energy_consumption_range_teh,omitempty" gorm:"-"`
	EnergyConsumptionEstimation *EnergyConsumptionEstimation `json:"energy_consumption_estimation,omitempty" gorm:"-"`
	DimensionsAndWeight         *DimensionsAndWeight         `json:"dimensions_and_weight,omitempty" gorm:"-"`
	Miscellaneous               *Miscellaneous               `json:"miscellaneous,omitempty" gorm:"-"`
	ChargeTypes                 []ChargeTypes                `json:"charge_types,omitempty" gorm:"-"`
	BaseRecordFields
}
