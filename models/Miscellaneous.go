package models

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
