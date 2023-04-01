package domain

import (
	"time"
)

type FeeService interface {
	GetParkingFee(vehicleType VehicleType, entryTime, exitTime time.Time) float64
}

type mallFeeModel struct {
	vehicleTypeToFee map[VehicleType]float64
}

func (feeModel mallFeeModel) GetParkingFee(vehicleType VehicleType, entryTime, exitTime time.Time) float64 {
	perHourFee := feeModel.vehicleTypeToFee[vehicleType]
	duration := int(exitTime.Sub(entryTime).Hours())
	return float64(duration) * perHourFee
}

func DefaultMallFeeModel() mallFeeModel {
	feeModel := mallFeeModel{}
	feeModel.vehicleTypeToFee = make(map[VehicleType]float64)
	feeModel.vehicleTypeToFee[Motorcycle] = 10.0
	feeModel.vehicleTypeToFee[CarOrSuv] = 20.0
	feeModel.vehicleTypeToFee[BusOrTruck] = 50.0
	return feeModel
}
