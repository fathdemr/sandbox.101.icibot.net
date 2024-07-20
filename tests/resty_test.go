/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 20.07.2024 Saturday
	Time        : 18:18
	Notes       :

*/

package tests

import (
	"github.com/go-resty/resty/v2"
	"sandbox.101.icibot.net/libs"
	"testing"
	"time"
)

type PingResponse struct {
	Amsterdam  time.Time `json:"amsterdam"`
	Istanbul   time.Time `json:"istanbul"`
	London     time.Time `json:"london"`
	Madrid     time.Time `json:"madrid"`
	Malta      time.Time `json:"malta"`
	ServerTime time.Time `json:"server_time"`
}

func TestResty(t *testing.T) {
	// Create a Resty Client
	client := resty.New()

	var response PingResponse

	resp, err := client.R().
		SetResult(&response).
		SetAuthToken("b1c1b0t").
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"User-Agent":   "My custom User Agent String",
		}).
		Get("https://api.icibot.net/ping")

	if err != nil {
		t.Error(err)
	}

	t.Log(resp.StatusCode())
	libs.PrintPrettyStruct(response)

}
