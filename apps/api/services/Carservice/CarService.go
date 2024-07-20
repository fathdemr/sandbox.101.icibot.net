package Carservice

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"sandbox.101.icibot.net/models"
)

type CarService struct {
	Db         *gorm.DB
	CurrentCar models.Car
	Cars       []models.Car
	Actor      models.User
}

func New(db *gorm.DB) (carService *CarService, err error) {
	carService = &CarService{
		Db: db,
	}

	err = carService.Db.Find(&carService.Cars).Error
	return
}

func (s *CarService) SetActor(user models.User) {
	s.Actor = user
}

func (s *CarService) GetCarById(id uint64) (err error) {

	//err = s.Db.
	//	Where("id = ?",
	//		id).
	//	First(&s.Car).Error
	var isExist bool
	s.CurrentCar, isExist = lo.Find(s.Cars, func(car models.Car) bool {
		return car.Id == id
	})

	if !isExist {
		err = fmt.Errorf("car not found")
	}

	return
}
