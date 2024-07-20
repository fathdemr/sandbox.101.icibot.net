package models

type BidirectionalCharging struct {
	Id                  uint64 `gorm:"primary_key" json:"id"`
	CarId               uint64 `json:"car_id"`
	V2LSupported        string `json:"v2l_supported"`
	ExteriorOutlets     string `json:"exterior_outlets"`
	V2LMaxOutputPower   string `json:"v2l_max_output_power"`
	InteriorOutlets     string `json:"interior_outlets"`
	V2HViaAcSupported   string `json:"v2h_via_ac_supported"`
	V2HViaDcSupported   string `json:"v2h_via_dc_supported"`
	V2HMaxOutputPowerAc string `json:"v2h_max_output_power_ac"`
	V2HMaxOutputPowerDc string `json:"v2h_max_output_power_dc"`
	V2GViaAcSupported   string `json:"v2g_via_ac_supported"`
	V2GMaxOutputPowerAc string `json:"v2g_max_output_power_ac"`
	V2GViaDcSupported   string `json:"v2g_via_dc_supported"`
	V2GMaxOutputPowerDc string `json:"v2g_max_output_power_dc"`
	BaseRecordFields
}
