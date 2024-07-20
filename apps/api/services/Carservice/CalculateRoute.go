/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 20.07.2024 Saturday
	Time        : 14:06
	Notes       :

*/

package Carservice

import (
	"sandbox.101.icibot.net/models"
	"time"
)

type CalculateRootRequest struct {
	FromDestination string           `json:"from_destination"`
	ToDestination   string           `json:"to_destination"`
	CurrentBattery  float32          `json:"current_battery"`
	RouteDate       time.Time        `json:"route_date"`
	Stations        []models.Station `json:"stations"` //
}

type CalculateRootResponse struct {
	Stations                  []models.Station `json:"stations"`
	NeedReCalculateByNewPoint bool             `json:"need_re_calculate_by_new_point"` // Eğer araç yeni bir noktaya ulaştığında yeni bir hesaplama yapılması gerekiyorsa true döner. Son noktaya ulaşmak için tekrar bir şarj istasyonuna gitmesi gerekiyorsa true döner.
}

func (s *CarService) CalculateRoot(request CalculateRootRequest) (response CalculateRootResponse, err error) {

	return
}

func (s *CarService) GetCelsiusByDestination(destination string) (celsius float32, err error) {

	switch destination {
	case "Ankara":
		celsius = 25
	case "İstanbul":
		celsius = 30
	case "İzmir":
		celsius = 35
	case "Antalya":
		celsius = 40
	case "Adana":
		celsius = 45
	case "Bursa":
		celsius = 24
	case "Rize":
		celsius = 20
	}

	winter := false
	if time.Now().Month() == time.January || time.Now().Month() == time.February || time.Now().Month() == time.December {
		winter = true
	}

	if winter {
		celsius -= 15
	}

	return
}
