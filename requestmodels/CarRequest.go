package requestmodels

type CarRequest struct {
	CarModel                    string                      `json:"car_model"`
	Images                      []string                    `json:"images" gorm:"-"`
	AvailableSince              string                      `json:"available_since"`
	UsableBattery               string                      `json:"usable_battery"`
	RealRange                   string                      `json:"real_range"`
	Efficiency                  string                      `json:"efficiency"`
	RealRangeEstimation         RealRangeEstimation         `gorm:"-" json:"real_range_estimation"`
	Performance                 Performance                 `gorm:"-" json:"performance"`
	Battery                     Battery                     `gorm:"-" json:"battery"`
	Charging                    Charging                    `gorm:"-" json:"charging"`
	BidirectionalCharging       BidirectionalCharging       `gorm:"-" json:"bidirectional_charging"`
	EnergyConsumptionRangeReal  EnergyConsumptionRangeReal  `gorm:"-" json:"energy_consumption_range_real"`
	EnergyConsumptionRangeTel   EnergyConsumptionRangeTel   `gorm:"-" json:"energy_consumption_range_tel"`
	EnergyConsumptionRangeTeh   EnergyConsumptionRangeTeh   `gorm:"-" json:"energy_consumption_range_teh"`
	EnergyConsumptionEstimation EnergyConsumptionEstimation `gorm:"-" json:"energy_consumption_estimation"`
	DimensionsAndWeight         DimensionsAndWeight         `gorm:"-" json:"dimensions_and_weight"`
	Miscellaneous               Miscellaneous               `gorm:"-" json:"miscellaneous"`
	ChargeTypes                 []ChargeTypes               `gorm:"-" json:"charge_types"`
}

type RealRangeEstimation struct {
	CityColdWeather     string `json:"city_cold_weather"`
	HighwayColdWeather  string `json:"highway_cold_weather"`
	CombinedColdWeather string `json:"combined_cold_weather"`
	CityMildWeather     string `json:"city_mild_weather"`
	HighwayMildWeather  string `json:"highway_mild_weather"`
	CombinedMildWeather string `json:"combined_mild_weather"`
}

type Performance struct {
	Acceleration0To100 string `json:"acceleration_0_to_100"`
	TotalPower         string `json:"total_power"`
	TopSpeed           string `json:"top_speed"`
	TotalTorque        string `json:"total_torque"`
	ElectricRange      string `json:"electric_range"`
	Drive              string `json:"drive"`
}

type Battery struct {
	NominalCapacity      string `json:"nominal_capacity"`
	UsableCapacity       string `json:"usable_capacity"`
	BatteryType          string `json:"battery_type"`
	CathodeMaterial      string `json:"cathode_material"`
	NumberOfCells        string `json:"number_of_cells"`
	PackConfiguration    string `json:"pack_configuration"`
	Architecture         string `json:"architecture"`
	NominalVoltage       string `json:"nominal_voltage"`
	WarrantyPeriod       string `json:"warranty_period"`
	FormFactor           string `json:"form_factor"`
	WarrantyMileage      string `json:"warranty_mileage"`
	BatteryNameReference string `json:"battery_name_reference"`
}

type Charging struct {
	HomeChargePort                 string `json:"home_charge_port"`
	HomeChargeTime0To415Km         string `json:"home_charge_time_0_to_415_km"`
	HomePortLocation               string `json:"home_port_location"`
	HomeChargeSpeed                string `json:"home_charge_speed"`
	HomeChargePower                string `json:"home_charge_power"`
	FastChargePort                 string `json:"fast_charge_port"`
	FastChargeTime41To332Km        string `json:"fast_charge_time_41_to_332_km"`
	FcPortLocation                 string `json:"fc_port_location"`
	FastChargeSpeed                string `json:"fast_charge_speed"`
	FastChargePowerMax             string `json:"fast_charge_power_max"`
	AutochargeSupported            string `json:"autocharge_supported"`
	FastChargePower10To80P         string `json:"fast_charge_power_10_to_80p"`
	PlugAndChargeSupported         string `json:"plug_and_charge_supported"`
	PlugAndChargeSupportedProtocol string `json:"plug_and_charge_supported_protocol"`
}

type BidirectionalCharging struct {
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
}

type EnergyConsumptionRangeReal struct {
	Range                 string `json:"range"`
	Co2Emissions          string `json:"co2_emissions"`
	VehicleConsumption    string `json:"vehicle_consumption"`
	VehicleFuelEquivalent string `json:"vehicle_fuel_equivalent"`
}

type EnergyConsumptionRangeTel struct {
	Range                 string `json:"range"`
	Co2Emissions          string `json:"co2_emissions"`
	VehicleConsumption    string `json:"vehicle_consumption"`
	VehicleFuelEquivalent string `json:"vehicle_fuel_equivalent"`
}

type EnergyConsumptionRangeTeh struct {
	Range                 string `json:"range"`
	Co2Emissions          string `json:"co2_emissions"`
	VehicleConsumption    string `json:"vehicle_consumption"`
	VehicleFuelEquivalent string `json:"vehicle_fuel_equivalent"`
}

type EnergyConsumptionEstimation struct {
	CityColdWeather     string `json:"city_cold_weather"`
	HighwayColdWeather  string `json:"highway_cold_weather"`
	CombinedColdWeather string `json:"combined_cold_weather"`
	CityMildWeather     string `json:"city_mild_weather"`
	HighwayMildWeather  string `json:"highway_mild_weather"`
	CombinedMildWeather string `json:"combined_mild_weather"`
}

type DimensionsAndWeight struct {
	Length                 string `json:"length"`
	Width                  string `json:"width"`
	WidthWithMirrors       string `json:"width_with_mirrors"`
	Height                 string `json:"height"`
	Wheelbase              string `json:"wheelbase"`
	WeightUnLadenEu        string `json:"weight_un_laden_eu"`
	GrossVehicleWeightGvwr string `json:"gross_vehicle_weight_gvwr"`
	MaxPayload             string `json:"max_payload"`
	CargoVolume            string `json:"cargo_volume"`
	CargoVolumeMax         string `json:"cargo_volume_max"`
	CargoVolumeFrunk       string `json:"cargo_volume_frunk"`
	RoofLoad               string `json:"roof_load"`
	TowHitchPossible       string `json:"tow_hitch_possible"`
	TowingWeightUnBraked   string `json:"towing_weight_un_braked"`
	TowingWeightBraked     string `json:"towing_weight_braked"`
	VerticalLoadMax        string `json:"vertical_load_max"`
}

type Miscellaneous struct {
	NumberSeats         string `json:"number_seats"`
	Isofix              string `json:"isofix"`
	TurningCircle       string `json:"turning_circle"`
	Platform            string `json:"platform"`
	EvDedicatedPlatform string `json:"ev_dedicated_platform"`
	CarBody             string `json:"car_body"`
	Segment             string `json:"segment"`
	RoofRails           string `json:"roof_rails"`
	HeatPump            string `json:"heat_pump"`
	HpStandardEquipment string `json:"hp_standard_equipment"`
}

type ChargingPoints struct {
	ChargePointName string `json:"charge_point_name"`
	MaxPower        string `json:"max_power"`
	AvgPower        string `json:"avg_power"`
	Time            string `json:"time"`
	Rate            string `json:"rate"`
}

type ChargeTypes struct {
	ChargeType     string           `json:"charge_type"`
	ImageUrl       string           `json:"image_url"`
	ChargingPoints []ChargingPoints `gorm:"foreignKey:ChargeTypeID" json:"charging_points"`
	Content        string           `json:"content"`
	Info           string           `json:"info"`
}
