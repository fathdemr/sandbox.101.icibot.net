package models

type Performance struct {
	Id                 uint   `gorm:"primary_key" json:"id"`
	CarId              uint   `json:"car_id"`
	Acceleration0To100 string `json:"acceleration_0_to_100"`
	TotalPower         string `json:"total_power"`
	TopSpeed           string `json:"top_speed"`
	TotalTorque        string `json:"total_torque"`
	ElectricRange      string `json:"electric_range"`
	Drive              string `json:"drive"`
}
