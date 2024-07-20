package models

import (
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type ChargingPoints struct {
	Id              uint64 `gorm:"primary_key" json:"id"`
	CarID           uint64 `json:"car_id"`
	ChargeTypeId    uint64 `json:"charge_type_id"`
	ChargePointName string `json:"charge_point_name"`
	MaxPower        string `json:"max_power"`
	AvgPower        string `json:"avg_power"`
	Time            string `json:"time"`
	ChargeMinutes   uint64 `json:"charge_minutes"`
	Rate            string `json:"rate"`
	BaseRecordFields
}

func (cp *ChargingPoints) BeforeSave(tx *gorm.DB) (err error) {
	err = cp.Validate()
	return

}

func (cp *ChargingPoints) Validate() (err error) {
	var hours int64
	var minutes int64

	if cp.Time != "" {
		if strings.Contains(cp.Time, " hours") {
			cp.Time = strings.ReplaceAll(cp.Time, " hours", "h")
			cp.Time += "0m"
		}
		timeValues := strings.Split(cp.Time, "h")
		if len(timeValues) == 2 {
			hours, _ = strconv.ParseInt(timeValues[0], 10, 64)

			timeValues[1] = strings.ReplaceAll(timeValues[1], "m", "")
			minutes, _ = strconv.ParseInt(timeValues[1], 10, 64)

			cp.ChargeMinutes = uint64(hours*60 + minutes)
		}
	}
	return
}

func (cp *ChargingPoints) BeforeCreate(tx *gorm.DB) (err error) {
	err = cp.Validate()
	return
}
