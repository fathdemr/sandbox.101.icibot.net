package Carservice

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sandbox.101.icibot.net/models"
)

func readJSONFile(filename string) ([]models.Car, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var cars []models.Car
	err = json.Unmarshal(byteValue, &cars)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (s *CarService) Create() {
	cars, err := readJSONFile("Data.json")
	if err != nil {
		log.Fatal("failed to read JSON file", err)
	}

	for _, car := range cars {
		result := s.Db.Create(&car)
		if result.Error != nil {
			log.Fatalf("failed to create car: %v", result.Error)
		}
	}
}
