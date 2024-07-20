package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sandbox.101.icibot.net/Config"
	"sandbox.101.icibot.net/models"
	"strconv"
	"strings"
)

func readJSONFiles(directory string) ([]CarRequest, error) {
	var cars []CarRequest

	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		filePath := directory + "/" + file.Name()

		jsonFile, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}

		byteValue, err := io.ReadAll(jsonFile)
		if err != nil {
			jsonFile.Close()
			return nil, err
		}
		jsonFile.Close()

		var fileCars []CarRequest
		err = json.Unmarshal(byteValue, &fileCars)
		if err != nil {
			var car CarRequest
			err = json.Unmarshal(byteValue, &car)
			if err != nil {
				return nil, err
			}
			fileCars = append(fileCars, car)
		}

		cars = append(cars, fileCars...)
	}

	return cars, nil
}

func convertCarRequestToCar(carRequest CarRequest) (models.Car, error) {
	NominalCapacity := strings.Replace(carRequest.Battery.NominalCapacity, " kWh", "", -1)
	UsableCapacity := strings.Replace(carRequest.Battery.UsableCapacity, " kWh", "", -1)

	ParseNominalCapacity, err := strconv.ParseFloat(NominalCapacity, 64)
	if err != nil {
		log.Fatal(err)
	}

	ParseUsableCapacity, err := strconv.ParseFloat(UsableCapacity, 64)
	if err != nil {
		log.Fatal(err)
	}

	/*HomeChargeTime0To415Km := strings.Replace(carRequest.Charging.HomeChargeTime0To415Km, "h", "", 2)
	HomeChargeTime0To415Km = strings.Replace(HomeChargeTime0To415Km, "m", "", 2)

	*/

	HomeChargeSpeed := strings.Replace(carRequest.Charging.HomeChargeSpeed, " km/h", "", -1)
	ParseHomeChargeSpeed, err := strconv.ParseUint(HomeChargeSpeed, 10, 64)

	HomeChargePower := strings.Replace(carRequest.Charging.HomeChargePower, " kW AC", "", -1)
	ParseHomeChargePower, err := strconv.ParseUint(HomeChargePower, 10, 64)

	FastChargeSpeed := strings.Replace(carRequest.Charging.FastChargeSpeed, " km/h", "", -1)
	ParseFastChargeSpeed, err := strconv.ParseUint(FastChargeSpeed, 10, 64)

	FastChargePowerMax := strings.Replace(carRequest.Charging.FastChargePowerMax, " kW DC", "", -1)
	ParseFastChargePowerMax, err := strconv.ParseUint(FastChargePowerMax, 10, 64)

	FastChargePower10To80P := strings.Replace(carRequest.Charging.FastChargePower10To80P, " kW DC", "", -1)
	ParseFastChargePower10To80P, err := strconv.ParseUint(FastChargePower10To80P, 10, 64)

	Range := strings.Replace(carRequest.EnergyConsumptionRangeReal.Range, " km", "", -1)
	ParseRange, err := strconv.ParseUint(Range, 10, 64)

	Co2Emissions := strings.Replace(carRequest.EnergyConsumptionRangeReal.Co2Emissions, " g/km", "", -1)
	ParseCo2Emissions, err := strconv.ParseUint(Co2Emissions, 10, 64)

	VehicleConsumption := strings.Replace(carRequest.EnergyConsumptionRangeReal.VehicleConsumption, " Wh/km", "", -1)
	ParseVehicleConsumption, err := strconv.ParseUint(VehicleConsumption, 10, 64)

	VehicleFuelEquivalent := strings.Replace(carRequest.EnergyConsumptionRangeReal.VehicleFuelEquivalent, " l/100km", "", -1)
	ParseVehicleFuelEquivalent, err := strconv.ParseFloat(VehicleFuelEquivalent, 64)

	TelRange := strings.Replace(carRequest.EnergyConsumptionRangeTel.Range, " km", "", -1)
	ParseTelRange, err := strconv.ParseUint(TelRange, 10, 64)

	TelCo2Emissions := strings.Replace(carRequest.EnergyConsumptionRangeTel.Co2Emissions, " g/km", "", -1)
	ParseTelCo2Emissions, err := strconv.ParseUint(TelCo2Emissions, 10, 64)

	TelVehicleConsumption := strings.Replace(carRequest.EnergyConsumptionRangeTel.VehicleConsumption, " Wh/km", "", -1)
	ParseTelVehicleConsumption, err := strconv.ParseUint(TelVehicleConsumption, 10, 64)

	TelVehicleFuelEquivalent := strings.Replace(carRequest.EnergyConsumptionRangeTel.VehicleFuelEquivalent, " l/100km", "", -1)
	ParseTelVehicleFuelEquivalent, err := strconv.ParseFloat(TelVehicleFuelEquivalent, 64)

	TehRange := strings.Replace(carRequest.EnergyConsumptionRangeTel.Range, " km", "", -1)
	ParseTehRange, err := strconv.ParseUint(TehRange, 10, 64)

	TehCo2Emissions := strings.Replace(carRequest.EnergyConsumptionRangeTel.Co2Emissions, " g/km", "", -1)
	ParseTehCo2Emissions, err := strconv.ParseUint(TehCo2Emissions, 10, 64)

	TehVehicleConsumption := strings.Replace(carRequest.EnergyConsumptionRangeTel.VehicleConsumption, " Wh/km", "", -1)
	ParseTehVehicleConsumption, err := strconv.ParseUint(TehVehicleConsumption, 10, 64)

	TehVehicleFuelEquivalent := strings.Replace(carRequest.EnergyConsumptionRangeTel.VehicleFuelEquivalent, " l/100km", "", -1)
	ParseTehVehicleFuelEquivalent, err := strconv.ParseFloat(TehVehicleFuelEquivalent, 64)

	CityColdWeather := strings.Replace(carRequest.EnergyConsumptionEstimation.CityColdWeather, " Wh/km", "", -1)
	ParseCityColdWeather, err := strconv.ParseUint(CityColdWeather, 10, 64)

	HighwayColdWeather := strings.Replace(carRequest.EnergyConsumptionEstimation.HighwayColdWeather, " Wh/km", "", -1)
	ParseHighwayColdWeather, err := strconv.ParseUint(HighwayColdWeather, 10, 64)

	CombinedColdWeather := strings.Replace(carRequest.EnergyConsumptionEstimation.CombinedColdWeather, " Wh/km", "", -1)
	ParseCombinedColdWeather, err := strconv.ParseUint(CombinedColdWeather, 10, 64)

	CityMildWeather := strings.Replace(carRequest.EnergyConsumptionEstimation.CityMildWeather, " Wh/km", "", -1)
	ParseCityMildWeather, err := strconv.ParseUint(CityMildWeather, 10, 64)

	HighwayMildWeather := strings.Replace(carRequest.EnergyConsumptionEstimation.HighwayMildWeather, " Wh/km", "", -1)
	ParseHighwayMildWeather, err := strconv.ParseUint(HighwayMildWeather, 10, 64)

	CombinedMildWeather := strings.Replace(carRequest.EnergyConsumptionEstimation.CombinedMildWeather, " Wh/km", "", -1)
	ParseCombinedMildWeather, err := strconv.ParseUint(CombinedMildWeather, 10, 64)

	RealRangeEstimation := models.RealRangeEstimation{
		CityColdWeather:     carRequest.RealRangeEstimation.CityColdWeather,
		HighwayColdWeather:  carRequest.RealRangeEstimation.HighwayColdWeather,
		CombinedColdWeather: carRequest.RealRangeEstimation.CombinedColdWeather,
		CityMildWeather:     carRequest.RealRangeEstimation.CityMildWeather,
		HighwayMildWeather:  carRequest.RealRangeEstimation.HighwayMildWeather,
		CombinedMildWeather: carRequest.RealRangeEstimation.CombinedMildWeather,
	}

	Performance := models.Performance{
		Acceleration0To100: carRequest.Performance.Acceleration0To100,
		TotalPower:         carRequest.Performance.TotalPower,
		TopSpeed:           carRequest.Performance.TopSpeed,
		TotalTorque:        carRequest.Performance.TotalTorque,
		ElectricRange:      carRequest.Performance.ElectricRange,
		Drive:              carRequest.Performance.Drive,
	}

	Battery := models.Battery{
		NominalCapacity:      ParseNominalCapacity,
		UsableCapacity:       ParseUsableCapacity,
		BatteryType:          carRequest.Battery.BatteryType,
		CathodeMaterial:      carRequest.Battery.CathodeMaterial,
		NumberOfCells:        carRequest.Battery.NumberOfCells,
		PackConfiguration:    carRequest.Battery.PackConfiguration,
		Architecture:         carRequest.Battery.Architecture,
		NominalVoltage:       carRequest.Battery.NominalVoltage,
		WarrantyPeriod:       carRequest.Battery.WarrantyPeriod,
		FormFactor:           carRequest.Battery.FormFactor,
		WarrantyMileage:      carRequest.Battery.WarrantyMileage,
		BatteryNameReference: carRequest.Battery.BatteryNameReference,
	}

	Charging := models.Charging{
		HomeChargePort: carRequest.Charging.HomeChargePort,
		//HomeChargeTime0To415Km
		HomePortLocation: carRequest.Charging.HomePortLocation,
		HomeChargeSpeed:  ParseHomeChargeSpeed,
		HomeChargePower:  ParseHomeChargePower,
		FastChargePort:   carRequest.Charging.FastChargePort,
		//FastChargeTime41To332
		FcPortLocation:                 carRequest.Charging.FcPortLocation,
		FastChargeSpeed:                ParseFastChargeSpeed,
		FastChargePowerMax:             ParseFastChargePowerMax,
		AutochargeSupported:            carRequest.Charging.AutochargeSupported,
		FastChargePower10To80P:         ParseFastChargePower10To80P,
		PlugAndChargeSupported:         carRequest.Charging.PlugAndChargeSupported,
		PlugAndChargeSupportedProtocol: carRequest.Charging.PlugAndChargeSupportedProtocol,
	}

	BidirectionalCharging := models.BidirectionalCharging{
		V2LSupported:        carRequest.BidirectionalCharging.V2LSupported,
		ExteriorOutlets:     carRequest.BidirectionalCharging.ExteriorOutlets,
		V2LMaxOutputPower:   carRequest.BidirectionalCharging.V2LMaxOutputPower,
		InteriorOutlets:     carRequest.BidirectionalCharging.InteriorOutlets,
		V2HViaAcSupported:   carRequest.BidirectionalCharging.V2HViaAcSupported,
		V2HViaDcSupported:   carRequest.BidirectionalCharging.V2HViaDcSupported,
		V2HMaxOutputPowerAc: carRequest.BidirectionalCharging.V2HMaxOutputPowerAc,
		V2HMaxOutputPowerDc: carRequest.BidirectionalCharging.V2HMaxOutputPowerDc,
		V2GViaAcSupported:   carRequest.BidirectionalCharging.V2GViaAcSupported,
		V2GMaxOutputPowerAc: carRequest.BidirectionalCharging.V2GMaxOutputPowerAc,
		V2GViaDcSupported:   carRequest.BidirectionalCharging.V2GViaDcSupported,
		V2GMaxOutputPowerDc: carRequest.BidirectionalCharging.V2GMaxOutputPowerDc,
	}

	EnergyConsumptionRangeReal := models.EnergyConsumptionRangeReal{
		Range:                 ParseRange,
		Co2Emissions:          ParseCo2Emissions,
		VehicleConsumption:    ParseVehicleConsumption,
		VehicleFuelEquivalent: ParseVehicleFuelEquivalent,
	}

	EnergyConsumptionRangeTel := models.EnergyConsumptionRangeTel{
		Range:                 ParseTelRange,
		Co2Emissions:          ParseTelCo2Emissions,
		VehicleConsumption:    ParseTelVehicleConsumption,
		VehicleFuelEquivalent: ParseTelVehicleFuelEquivalent,
	}

	EnergyConsumptionRangeTeh := models.EnergyConsumptionRangeTeh{
		Range:                 ParseTehRange,
		Co2Emissions:          ParseTehCo2Emissions,
		VehicleConsumption:    ParseTehVehicleConsumption,
		VehicleFuelEquivalent: ParseTehVehicleFuelEquivalent,
	}

	EnergyConsumptionEstimation := models.EnergyConsumptionEstimation{
		CityColdWeather:     ParseCityColdWeather,
		HighwayColdWeather:  ParseHighwayColdWeather,
		CombinedColdWeather: ParseCombinedColdWeather,
		CityMildWeather:     ParseCityMildWeather,
		HighwayMildWeather:  ParseHighwayMildWeather,
		CombinedMildWeather: ParseCombinedMildWeather,
	}

	DimensionsAndWeight := models.DimensionsAndWeight{
		Length:                 carRequest.DimensionsAndWeight.Length,
		Width:                  carRequest.DimensionsAndWeight.Width,
		WidthWithMirrors:       carRequest.DimensionsAndWeight.WidthWithMirrors,
		Height:                 carRequest.DimensionsAndWeight.Height,
		Wheelbase:              carRequest.DimensionsAndWeight.Wheelbase,
		WeightUnLadenEu:        carRequest.DimensionsAndWeight.WeightUnLadenEu,
		GrossVehicleWeightGvwr: carRequest.DimensionsAndWeight.GrossVehicleWeightGvwr,
		MaxPayload:             carRequest.DimensionsAndWeight.MaxPayload,
		CargoVolume:            carRequest.DimensionsAndWeight.CargoVolume,
		CargoVolumeMax:         carRequest.DimensionsAndWeight.CargoVolumeMax,
		CargoVolumeFrunk:       carRequest.DimensionsAndWeight.CargoVolumeFrunk,
		RoofLoad:               carRequest.DimensionsAndWeight.RoofLoad,
		TowHitchPossible:       carRequest.DimensionsAndWeight.TowHitchPossible,
		TowingWeightUnBraked:   carRequest.DimensionsAndWeight.TowingWeightUnBraked,
		TowingWeightBraked:     carRequest.DimensionsAndWeight.TowingWeightBraked,
		VerticalLoadMax:        carRequest.DimensionsAndWeight.VerticalLoadMax,
	}

	Miscellaneous := models.Miscellaneous{
		NumberSeats:         carRequest.Miscellaneous.NumberSeats,
		Isofix:              carRequest.Miscellaneous.Isofix,
		TurningCircle:       carRequest.Miscellaneous.TurningCircle,
		Platform:            carRequest.Miscellaneous.Platform,
		EvDedicatedPlatform: carRequest.Miscellaneous.EvDedicatedPlatform,
		CarBody:             carRequest.Miscellaneous.CarBody,
		Segment:             carRequest.Miscellaneous.Segment,
		RoofRails:           carRequest.Miscellaneous.RoofRails,
		HeatPump:            carRequest.Miscellaneous.HeatPump,
		HpStandardEquipment: carRequest.Miscellaneous.HpStandardEquipment,
	}

	var chargeTypes []models.ChargeTypes
	for _, ct := range carRequest.ChargeTypes {
		var chargingPoints []models.ChargingPoints
		for _, cp := range ct.ChargingPoints {
			chargingPoint := models.ChargingPoints{
				ChargePointName: cp.ChargePointName,
				MaxPower:        cp.MaxPower,
				AvgPower:        cp.AvgPower,
				Time:            cp.Time,
				Rate:            cp.Rate,
			}
			chargingPoints = append(chargingPoints, chargingPoint)
		}
		chargeType := models.ChargeTypes{
			ChargeType:     ct.ChargeType,
			ImageUrl:       ct.ImageUrl,
			ChargingPoints: chargingPoints,
			Content:        ct.Content,
			Info:           ct.Info,
		}
		chargeTypes = append(chargeTypes, chargeType)
	}

	car := models.Car{
		CarModel:                    carRequest.CarModel,
		AvailableSince:              carRequest.AvailableSince,
		UsableBattery:               carRequest.UsableBattery,
		RealRange:                   carRequest.RealRange,
		Efficiency:                  carRequest.Efficiency,
		RealRangeEstimation:         &RealRangeEstimation,
		Performance:                 &Performance,
		Battery:                     &Battery,
		Charging:                    &Charging,
		BidirectionalCharging:       &BidirectionalCharging,
		EnergyConsumptionRangeReal:  &EnergyConsumptionRangeReal,
		EnergyConsumptionRangeTel:   &EnergyConsumptionRangeTel,
		EnergyConsumptionRangeTeh:   &EnergyConsumptionRangeTeh,
		EnergyConsumptionEstimation: &EnergyConsumptionEstimation,
		DimensionsAndWeight:         &DimensionsAndWeight,
		Miscellaneous:               &Miscellaneous,
		ChargeTypes:                 chargeTypes,
	}

	return car, err

}

func main() {

	err := Config.InitDb()

	requestCars, err := readJSONFiles("cmd/EVDatabase")
	if err != nil {
		log.Fatal("failed to read JSON file:", err)
	}

	for _, car := range requestCars {

		modelCars, err := convertCarRequestToCar(car)
		if err != nil {
			log.Fatal("failed to convert car request to car model:", err)
		}

		result := Config.Db.Create(&modelCars)
		if result.Error != nil {
			log.Fatalf("failed to create car: %v", result.Error)
		}

		modelCars.RealRangeEstimation.CarId = modelCars.Id

		Config.Db.Create(&modelCars.RealRangeEstimation)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.Performance.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.Performance)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.Battery.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.Battery)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.Charging.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.Charging)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.BidirectionalCharging.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.BidirectionalCharging)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.EnergyConsumptionRangeReal.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.EnergyConsumptionRangeReal)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.EnergyConsumptionRangeTel.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.EnergyConsumptionRangeTel)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.EnergyConsumptionRangeTeh.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.EnergyConsumptionRangeTeh)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.EnergyConsumptionEstimation.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.EnergyConsumptionEstimation)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.DimensionsAndWeight.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.DimensionsAndWeight)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		modelCars.Miscellaneous.CarId = modelCars.Id

		result = Config.Db.Create(&modelCars.Miscellaneous)
		if result.Error != nil {
			log.Fatalf("failed to create battery: %v", result.Error)
		}

		for i := range modelCars.ChargeTypes {
			modelCars.ChargeTypes[i].CarId = modelCars.Id
			result = Config.Db.Create(&modelCars.ChargeTypes[i])
			if result.Error != nil {
				log.Fatalf("failed to create chargetypes: %v", result.Error)
			}
			for j := range modelCars.ChargeTypes[i].ChargingPoints {
				modelCars.ChargeTypes[i].ChargingPoints[j].ChargeTypeId = modelCars.ChargeTypes[i].Id
				result = Config.Db.Create(&modelCars.ChargeTypes[i].ChargingPoints[j])
				if result.Error != nil {
					log.Fatalf("failed to create chargepoints: %v", result.Error)
				}
			}
		}

	}

}
