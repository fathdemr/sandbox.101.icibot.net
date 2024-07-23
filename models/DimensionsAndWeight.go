package models

type DimensionsAndWeight struct {
	Id                     uint64 `gorm:"primary_key" json:"id"`
	CarId                  uint64 `json:"car_id"`
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
