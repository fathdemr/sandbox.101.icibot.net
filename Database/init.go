package Database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	appDBHost         string
	appDBUserName     string
	appDBUserPassword string
	appDBName         string
	appDBPort         string
	Db                *gorm.DB
)

func InitDb() error {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	appDBHost = viper.GetString("APP_DB_HOST")
	appDBUserName = viper.GetString("APP_DB_USER_NAME")
	appDBUserPassword = viper.GetString("APP_DB_USER_PASSWORD")
	appDBName = viper.GetString("APP_DB_NAME")
	appDBPort = viper.GetString("APP_DB_PORT")

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
