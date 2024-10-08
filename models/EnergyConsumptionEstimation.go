package models

type EnergyConsumptionEstimation struct {
	CityColdWeather     uint64 `json:"city_cold_weather"`
	HighwayColdWeather  uint64 `json:"highway_cold_weather"`
	CombinedColdWeather uint64 `json:"combined_cold_weather"`
	CityMildWeather     uint64 `json:"city_mild_weather"`
	HighwayMildWeather  uint64 `json:"highway_mild_weather"`
	CombinedMildWeather uint64 `json:"combined_mild_weather"`
}
