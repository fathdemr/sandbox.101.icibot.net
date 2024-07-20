/*
	B1 Yönetim Sistemleri Yazılım ve Danışmanlık Ltd. Şti.
	User        : ICI
	Name        : Ibrahim ÇOBANİ
	Date        : 20.07.2024 Saturday
	Time        : 14:32
	Notes       :

*/

package Carservice

import (
	"fmt"
	"sandbox.101.icibot.net/models"
)

func (s *CarService) Create(car models.Car) (response models.Car, err error) {

	if car.Id == 0 {
		err = fmt.Errorf("id field is not valid for create operation")
		return
	}

	// Create Car
	car.SetCreatedUser(
		s.Actor.ID,
		fmt.Sprintf("%s %s", s.Actor.Name, s.Actor.LastName),
		s.Actor.CurrentIpAddress)

	err = s.Db.Create(&car).Error
	if err != nil {
		return
	}

	response = car

	return
}
