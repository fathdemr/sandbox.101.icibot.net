package libs

/*
   ©️ 2022 B1 Digital
   User    : ODI
   Name    : Özlem DEMİRCİ
   Date    : 15.02.2022  12:47
   Notes   :
*/

import (
	"encoding/json"
	"fmt"
)

func PrintPrettyStruct(value interface{}) {
	valueJSON, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%s\n", string(valueJSON))
}
