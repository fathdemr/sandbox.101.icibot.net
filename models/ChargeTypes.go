package models

type ChargeTypes struct {
	Id             uint64           `gorm:"primary_key" json:"id"`
	CarId          uint64           `json:"car_id"`
	ChargeType     string           `json:"charge_type"`
	ImageUrl       string           `json:"image_url"`
	ChargingPoints []ChargingPoints `json:"charging_points" gorm:"-"`
	Content        string           `json:"content"`
	Info           string           `json:"info"`
	BaseRecordFields
}
