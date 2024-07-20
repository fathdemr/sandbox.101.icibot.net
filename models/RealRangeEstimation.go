package models

type RealRangeEstimation struct {
	Id                  uint64 `gorm:"primary_key" json:"id"`
	CarId               uint64 `json:"car_id"`
	CityColdWeather     string `json:"city_cold_weather"`
	HighwayColdWeather  string `json:"highway_cold_weather"`
	CombinedColdWeather string `json:"combined_cold_weather"`
	CityMildWeather     string `json:"city_mild_weather"`
	HighwayMildWeather  string `json:"highway_mild_weather"`
	CombinedMildWeather string `json:"combined_mild_weather"`
	BaseRecordFields
}
