package models

type ChargingPoints struct {
	Id              uint   `gorm:"primary_key" json:"id"`
	CarID           uint   `json:"car_id"`
	ChargeTypeId    uint   `json:"charge_type_id"`
	ChargePointName string `json:"charge_point_name"`
	MaxPower        string `json:"max_power"`
	AvgPower        string `json:"avg_power"`
	Time            string `json:"time"`
	Rate            string `json:"rate"`
}
