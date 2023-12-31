package parking

import (
	"parking-lot/internal/domain"

	"math"
	"time"
)

//go:generate moq -out fee_service_mock.go . FeeService
type FeeService interface {
	GetParkingFee(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64
}

const hourUnknown = -1

type interval struct {
	start int
	end   int
}

type intervalBasedFeeModel struct {
	motorcyleToIntervals map[domain.VehicleType]map[interval]float64
	carOrSuvToIntervals  map[domain.VehicleType]map[interval]float64

	vehicleTypeToIntervalFeeContainer map[domain.VehicleType]map[domain.VehicleType]map[interval]float64
}

type stadiumFeeModel struct {
	intervalBasedFeeModel
}

func Stadium() FeeService {
	feeModel := stadiumFeeModel{}

	feeModel.motorcyleToIntervals = make(map[domain.VehicleType]map[interval]float64)
	feeModel.carOrSuvToIntervals = make(map[domain.VehicleType]map[interval]float64)

	feeModel.vehicleTypeToIntervalFeeContainer = make(map[domain.VehicleType]map[domain.VehicleType]map[interval]float64)

	feeModel.vehicleTypeToIntervalFeeContainer[domain.Motorcycle] = feeModel.motorcyleToIntervals
	feeModel.vehicleTypeToIntervalFeeContainer[domain.CarOrSuv] = feeModel.carOrSuvToIntervals

	feeModel.motorcyleToIntervals[domain.Motorcycle] = make(map[interval]float64)
	feeModel.carOrSuvToIntervals[domain.CarOrSuv] = make(map[interval]float64)

	i1 := interval{start: 0, end: 4}
	i2 := interval{start: 4, end: 12}
	i3 := interval{start: 12, end: hourUnknown}
	feeModel.motorcyleToIntervals[domain.Motorcycle] = map[interval]float64{i1: 30, i2: 60, i3: 100}

	i1 = interval{start: 0, end: 4}
	i2 = interval{start: 4, end: 12}
	i3 = interval{start: 12, end: hourUnknown}
	feeModel.carOrSuvToIntervals[domain.CarOrSuv] = map[interval]float64{i1: 60, i2: 120, i3: 200}

	return feeModel
}

func (feeModel stadiumFeeModel) GetParkingFee(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
	vehicleTypeToIntervalFee, ok := feeModel.vehicleTypeToIntervalFeeContainer[vehicleType]
	if !ok {
		return 0.0
	}
	intervalFee, ok := vehicleTypeToIntervalFee[vehicleType]
	if !ok {
		return 0.0
	}
	duration := int(math.Ceil(exitTime.Sub(entryTime).Hours()))

	var totalFee float64
	durationLeft := duration

	for interval, fee := range intervalFee {
		intervalStart := interval.start
		intervalEnd := interval.end

		if intervalEnd == hourUnknown && duration > 0 {
			for i := intervalStart; i < duration; i++ {
				totalFee += fee
				durationLeft -= 1
			}
			continue
		}

		if durationLeft >= intervalStart {
			totalFee += fee
			durationLeft -= 1
		}

	}
	return totalFee
}

type airportFeeModel struct {
	intervalBasedFeeModel
}

func Airport() airportFeeModel {
	feeModel := airportFeeModel{}

	feeModel.motorcyleToIntervals = make(map[domain.VehicleType]map[interval]float64)
	feeModel.carOrSuvToIntervals = make(map[domain.VehicleType]map[interval]float64)

	feeModel.vehicleTypeToIntervalFeeContainer = make(map[domain.VehicleType]map[domain.VehicleType]map[interval]float64)

	feeModel.vehicleTypeToIntervalFeeContainer[domain.Motorcycle] = feeModel.motorcyleToIntervals
	feeModel.vehicleTypeToIntervalFeeContainer[domain.CarOrSuv] = feeModel.carOrSuvToIntervals

	feeModel.motorcyleToIntervals[domain.Motorcycle] = make(map[interval]float64)
	feeModel.carOrSuvToIntervals[domain.CarOrSuv] = make(map[interval]float64)

	i1 := interval{start: 0, end: 1}
	i2 := interval{start: 1, end: 8}
	i3 := interval{start: 8, end: 24}
	i4 := interval{start: 24, end: hourUnknown}
	feeModel.motorcyleToIntervals[domain.Motorcycle] = map[interval]float64{i1: 0, i2: 40, i3: 60, i4: 80}

	i1 = interval{start: 0, end: 12}
	i2 = interval{start: 12, end: 24}
	i3 = interval{start: 24, end: hourUnknown}
	feeModel.carOrSuvToIntervals[domain.CarOrSuv] = map[interval]float64{i1: 60, i2: 80, i3: 100}

	return feeModel
}

func (feeModel airportFeeModel) GetParkingFee(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
	vehicleTypeToIntervalFee, ok := feeModel.vehicleTypeToIntervalFeeContainer[vehicleType]
	if !ok {
		return 0.0
	}
	intervalFee, ok := vehicleTypeToIntervalFee[vehicleType]
	if !ok {
		return 0.0
	}

	duration := exitTime.Sub(entryTime).Hours()
	durationLeftInHours := int(duration)
	durationInDays := int(duration / 24.0)

	if durationInDays > 0 {
		durationLeftInHours = int(duration - (float64(durationInDays) * 24.0))
	}

	var totalFee float64
	durationLeft := durationLeftInHours

	for interval, fee := range intervalFee {
		intervalStart := interval.start
		intervalEnd := interval.end

		if intervalEnd == hourUnknown && durationInDays > 0 {
			totalFee = 0
			if durationLeft > 0 {
				durationInDays += 1
			}
			for i := (intervalStart / 24); i <= durationInDays; i++ {
				totalFee += fee
			}
			break
		}

		if durationLeft >= intervalStart && totalFee < fee {
			totalFee = fee
		}
	}
	return totalFee
}

type flatRateFeeModel struct {
	vehicleTypeToFee map[domain.VehicleType]float64
}

type mallFeeModel struct {
	flatRateFeeModel
}

func Mall() FeeService {
	feeModel := mallFeeModel{}
	feeModel.vehicleTypeToFee = make(map[domain.VehicleType]float64)
	feeModel.vehicleTypeToFee[domain.Motorcycle] = 10.0
	feeModel.vehicleTypeToFee[domain.CarOrSuv] = 20.0
	feeModel.vehicleTypeToFee[domain.BusOrTruck] = 50.0
	return feeModel
}

func (feeModel mallFeeModel) GetParkingFee(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
	perHourFee := feeModel.vehicleTypeToFee[vehicleType]
	duration := int(math.Ceil(exitTime.Sub(entryTime).Hours()))
	return float64(duration) * perHourFee
}
