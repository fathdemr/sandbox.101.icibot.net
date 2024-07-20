package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	appDBHost         = os.Getenv("APP_DB_HOST")
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
