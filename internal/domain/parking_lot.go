package domain

import (
	"errors"
	"fmt"
)

type ParkingLot interface {
	Park(vehicleType VehicleType) (spotNumber int, err error)
	Unpark(spotNumber int) error
	AddSpots(vehicleType VehicleType, size int)
}

type parkingLot struct {
	vehicleTypeToSpot map[VehicleType][]*Spot
	spotNumberToSpot  map[int]Spot
}

func NewParkingLot() ParkingLot {
	pl := &parkingLot{}
	pl.spotNumberToSpot = make(map[int]Spot)
	pl.vehicleTypeToSpot = make(map[VehicleType][]*Spot)
	return pl
}

type ParkingLotError struct {
	Err error
}

func (parkingLotErr ParkingLotError) Error() string {
	return parkingLotErr.Err.Error()
}

var (
	ErrParkingLotFull            = errors.New("parking lot is full")
	ErrVehicleInSpotDoesNotExist = errors.New("vehicle is not in spot")
	ErrNoParkingSpotsAvailable   = func(vehicleType VehicleType) error {
		return fmt.Errorf("no parking spots available for vehicle type: %s", vehicleType)
	}
)

func (parkingLot *parkingLot) Park(vehicleType VehicleType) (spotNumber int, err error) {
	spots := parkingLot.vehicleTypeToSpot[vehicleType]
	if len(spots) == 0 {
		err = ParkingLotError{Err: ErrNoParkingSpotsAvailable(vehicleType)}
		return
	}
	var isVehicleParked bool
	for _, spot := range spots {
		if !spot.occupied {
			spot.occupied = true
			isVehicleParked = true
			spotNumber = spot.number
			parkingLot.spotNumberToSpot[spotNumber] = *spot
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
		return ParkingLotError{Err: ErrVehicleInSpotDoesNotExist}
	}
	spots := parkingLot.vehicleTypeToSpot[occupiedSpot.typ]
	for _, spot := range spots {
		if spot.number == spotNumber {
			spot.occupied = false
			delete(parkingLot.spotNumberToSpot, spotNumber)
			break
		}
	}
	return nil
}

func (pl *parkingLot) AddSpots(typ VehicleType, size int) {
	for i := 1; i <= size; i++ {
		spot := NewSpot(i, typ)
		pl.vehicleTypeToSpot[typ] = append(pl.vehicleTypeToSpot[typ], &spot)
	}
}
