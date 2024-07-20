package models

type Battery struct {
	Id                   uint    `gorm:"primary_key" json:"id"`
	CarId                uint    `json:"car_id"`
	NominalCapacity      float64 `json:"nominal_capacity"`
	UsableCapacity       float64 `json:"usable_capacity"`
	BatteryType          string  `json:"battery_type"`
	CathodeMaterial      string  `json:"cathode_material"`
	NumberOfCells        string  `json:"number_of_cells"`
	PackConfiguration    string  `json:"pack_configuration"`
	Architecture         string  `json:"architecture"`
	NominalVoltage       string  `json:"nominal_voltage"`
	WarrantyPeriod       string  `json:"warranty_period"`
	FormFactor           string  `json:"form_factor"`
	WarrantyMileage      string  `json:"warranty_mileage"`
	BatteryNameReference string  `json:"battery_name_reference"`
}
