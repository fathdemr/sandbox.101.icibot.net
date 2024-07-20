/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 20.07.2024 Saturday
	Time        : 14:17
	Notes       :

*/

package models

type Station struct {
	StationName string  `json:"station_name"`
	StationType string  `json:"station_type"`
	StationLat  float32 `json:"station_lat"`
	StationLng  float32 `json:"station_lng"`
	BaseRecordFields
}
