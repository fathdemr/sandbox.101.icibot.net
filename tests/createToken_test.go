/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 20.07.2024 Saturday
	Time        : 17:09
	Notes       :

*/

package tests

import (
	"sandbox.101.icibot.net/Config"
	"sandbox.101.icibot.net/middlewares"
	"sandbox.101.icibot.net/models"
	"testing"
)

func TestCreateToken(t *testing.T) {
	err := Config.ReadConfig()
	if err != nil {
		t.Error("Config dosyası okunamadı. err : ", err)
	}

	var user = models.User{
		ID:       1,
		Name:     "Ibrahim",
		LastName: "ÇOBANİ",
	}

	if token, err := middlewares.CreateToken(user); err != nil {
		t.Error("Token oluşturulamadı. err : ", err)
	} else {
		t.Log("Token oluşturuldu. Token : ", token)
	}

}
