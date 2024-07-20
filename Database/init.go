package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	appDBHost         = os.Getenv("APP_DB_HOST")
	appDBUserName     = os.Getenv("APP_DB_USER_NAME")
	appDBUserPassword = os.Getenv("APP_DB_USER_PASSWORD")
	appDBName         = os.Getenv("APP_DB_NAME")
	appDBPort         = os.Getenv("APP_DB_PORT")
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
