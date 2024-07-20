package Carservice

import (
	"gorm.io/gorm"
	"sandbox.101.icibot.net/models"
)

type CarService struct {
	Db *gorm.DB
}

func New(db *gorm.DB) *CarService {
	return &CarService{
		Db: db,
	}
}

func (s *CarService) GetCarById(id uint64) (cars models.Car, err error) {
	s.Db.Where("id = ?", id).First(&cars)
	return
}
