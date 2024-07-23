package models

type Performance struct {
	Id                    uint64  `gorm:"primary_key" json:"id"`
	CarId                 uint64  `json:"car_id"`
	Acceleration0To100    string  `json:"acceleration_0_to_100"`
	Acceleration0To100SEC float64 `json:"acceleration_0_to_100_sec"`
	TotalPower            string  `json:"total_power"`
	TotalPowerKW          uint64  `json:"total_power_kw"`
	TopSpeed              string  `json:"top_speed"`
	TopSpeedKM            uint64  `json:"top_speed_km"`
	TotalTorque           string  `json:"total_torque"`
	ElectricRange         string  `json:"electric_range"`
	ElectricRangeKM       uint64  `json:"electric_range_km"`
	Drive                 string  `json:"drive"`
}
