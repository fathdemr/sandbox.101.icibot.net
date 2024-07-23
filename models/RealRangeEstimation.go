package models

type RealRangeEstimation struct {
	CityColdWeather       string `json:"city_cold_weather"`
	CityColdWeatherKM     uint64 `json:"city_cold_weather_km"`
	HighwayColdWeather    string `json:"highway_cold_weather"`
	HighwayColdWeatherKM  uint64 `json:"highway_cold_weather_km"`
	CombinedColdWeather   string `json:"combined_cold_weather"`
	CombinedColdWeatherKM uint64 `json:"combined_cold_weather_km"`
	CityMildWeather       string `json:"city_mild_weather"`
	CityMildWeatherKM     uint64 `json:"city_mild_weather_km"`
	HighwayMildWeather    string `json:"highway_mild_weather"`
	HighwayMildWeatherKM  uint64 `json:"highway_mild_weather_km"`
	CombinedMildWeather   string `json:"combined_mild_weather"`
	CombinedMildWeatherKM uint64 `json:"combined_mild_weather_km"`
}
