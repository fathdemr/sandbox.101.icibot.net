package Config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sandbox.101.icibot.net/apps/api/services/Carservice"
	"time"
)

var (
	appDBHost           string
	appDBUserName       string
	appDBUserPassword   string
	appDBName           string
	appDBPort           string
	Db                  *gorm.DB
	CarService          *Carservice.CarService
	Viper               = viper.New()
	AppSecretKey        = ""
	TokenExpireDuration time.Duration
)

func InitDb() error {

	err := ReadConfig()
	if err != nil {
		return err
	}

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

	fmt.Println("DB Connected")

	Db = db

	CarService, err = Carservice.New(Db)
	return nil
}

func ReadConfig() error {
	Viper.SetConfigFile(".env")
	err := Viper.ReadInConfig()
	if err != nil {
		return err
	}

	appDBHost = Viper.GetString("APP_DB_HOST")
	appDBUserName = Viper.GetString("APP_DB_USER_NAME")
	appDBUserPassword = Viper.GetString("APP_DB_USER_PASSWORD")
	appDBName = Viper.GetString("APP_DB_NAME")
	appDBPort = Viper.GetString("APP_DB_PORT")
	AppSecretKey = Viper.GetString("APP_TOKEN_SECRET")
	TokenExpireDuration = time.Duration(Viper.GetInt("APP_TOKEN_EXPIRE_DURATION")) * time.Minute

	return nil
}
