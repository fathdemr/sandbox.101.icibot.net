package models

type Charging struct {
	Id                             uint   `gorm:"primary_key" json:"id"`
	CarId                          uint   `json:"car_id"`
	HomeChargePort                 string `json:"home_charge_port"`
	HomeChargeTime0To415Km         uint64 `json:"home_charge_time_0_to_415_km"`
	HomePortLocation               string `json:"home_port_location"`
	HomeChargeSpeed                uint64 `json:"home_charge_speed"`
	HomeChargePower                uint64 `json:"home_charge_power"`
	FastChargePort                 string `json:"fast_charge_port"`
	FastChargeTime41To332          uint64 `json:"fast_charge_time_41_to_332_km"`
	FcPortLocation                 string `json:"fc_port_location"`
	FastChargeSpeed                uint64 `json:"fast_charge_speed"`
	FastChargePowerMax             uint64 `json:"fast_charge_power_max"`
	AutochargeSupported            string `json:"autocharge_supported"`
	FastChargePower10To80P         uint64 `json:"fast_charge_power_10_to_80p"`
	PlugAndChargeSupported         string `json:"plug_and_charge_supported"`
	PlugAndChargeSupportedProtocol string `json:"plug_and_charge_supported_protocol"`
}
