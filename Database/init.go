package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	appDBHost         = "sandbox.s101.icibot.net"
	appDBUserName     = "fathdemr"
	appDBUserPassword = "12karakter5334"
	appDBName         = "dbsandbox"
	appDBPort         = "5432"
	Db                *gorm.DB
)

func InitDb() error {
	cnnString := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s sslmode=disable",
		appDBHost,
		appDBUserName,
		appDBUserPassword,
		appDBName,
		appDBPort)

	db, err := gorm.Open(postgres.Open(cnnString), &gorm.Config{})
	if err != nil {
		return err
	}

	Db = db
	return nil
}
