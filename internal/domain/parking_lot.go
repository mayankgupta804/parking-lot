package domain

import (
	"errors"
)

type ParkingLot interface {
	Park(vehicleType VehicleType) (spotNumber int, err error)
	Unpark(spotNumber int) error
	AddSpots(vehicleType VehicleType, size int)
}

type parkingLot struct {
	vehicleTypeToSpot map[VehicleType][]Spot
	spotNumberToSpot  map[int]Spot
}

func NewParkingLot() ParkingLot {
	pl := &parkingLot{}
	pl.spotNumberToSpot = make(map[int]Spot)
	pl.vehicleTypeToSpot = make(map[VehicleType][]Spot)
	return pl
}

type ParkingLotError struct {
	Err error
}

func (parkingLotErr ParkingLotError) Error() string {
	return parkingLotErr.Err.Error()
}

func (parkingLotErr ParkingLotError) Unwrap() error {
	return parkingLotErr.Err
}

var (
	ErrParkingLotFull            = errors.New("parking lot is full")
	ErrVehicleAlreadyParked      = errors.New("vehicle is already parked")
	ErrVehicleInSlotDoesNotExist = errors.New("vehicle is not in slot")
)

func (parkingLot *parkingLot) Park(vehicleType VehicleType) (spotNumber int, err error) {
	spots := parkingLot.vehicleTypeToSpot[vehicleType]
	var isVehicleParked bool
	for _, spot := range spots {
		if !spot.occupied {
			spot.occupied = true
			isVehicleParked = true
			spotNumber = spot.number
			parkingLot.spotNumberToSpot[spotNumber] = spot
			break
		}
	}
	if isVehicleParked {
		return
	}
	err = ParkingLotError{Err: ErrParkingLotFull}
	return
}

func (parkingLot *parkingLot) Unpark(spotNumber int) error {
	occupiedSpot, ok := parkingLot.spotNumberToSpot[spotNumber]
	if !ok {
		return ParkingLotError{Err: ErrVehicleInSlotDoesNotExist}
	}
	spots := parkingLot.vehicleTypeToSpot[occupiedSpot.typ]
	for _, spot := range spots {
		if spot.number == spotNumber {
			spot.occupied = false
			break
		}
	}
	return nil
}

func (pl *parkingLot) AddSpots(typ VehicleType, size int) {
	for i := 1; i <= size; i++ {
		spot := NewSpot(i, typ)
		pl.vehicleTypeToSpot[typ] = append(pl.vehicleTypeToSpot[typ], spot)
	}
}
