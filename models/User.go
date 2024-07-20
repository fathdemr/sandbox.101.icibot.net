/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 20.07.2024 Saturday
	Time        : 14:34
	Notes       :

*/

package models

type User struct {
	ID               uint64 `json:"id,omitempty" gorm:"primary_key"`
	Name             string `json:"name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	CurrentIpAddress string `json:"current_ip_address,omitempty"`
	BaseRecordFields
}
