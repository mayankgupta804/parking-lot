package domain

import (
	"errors"
	"fmt"
)

//go:generate moq -out parking_lot_mock.go . ParkingLot
type ParkingLot interface {
	Park(vehicleType VehicleType) (spotNumber int, err error)
	Unpark(spotNumber int) error
	AddSpots(vehicleType VehicleType, size int)
}

type parkingLot struct {
	vehicleTypeToTotalSpots map[VehicleType][]*Spot
	spotNumberToSpot        map[int]Spot
}

func NewParkingLot() ParkingLot {
	pl := &parkingLot{}
	pl.spotNumberToSpot = make(map[int]Spot)
	pl.vehicleTypeToTotalSpots = make(map[VehicleType][]*Spot)
	return pl
}

var (
	ErrParkingLotFull            = errors.New("parking lot is full")
	ErrVehicleInSpotDoesNotExist = errors.New("vehicle is not in spot")
	ErrNoParkingSpotsAvailable   = func(vehicleType VehicleType) error {
		return fmt.Errorf("no parking spots available for vehicle type: %s", vehicleType)
	}
)

func (parkingLot *parkingLot) Park(vehicleType VehicleType) (spotNumber int, err error) {
	spots := parkingLot.vehicleTypeToTotalSpots[vehicleType]
	if len(spots) == 0 {
		err = ErrNoParkingSpotsAvailable(vehicleType)
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
	err = ErrParkingLotFull
	return
}

func (parkingLot *parkingLot) Unpark(spotNumber int) error {
	occupiedSpot, ok := parkingLot.spotNumberToSpot[spotNumber]
	if !ok {
		return ErrVehicleInSpotDoesNotExist
	}
	spots := parkingLot.vehicleTypeToTotalSpots[occupiedSpot.typ]
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
		pl.vehicleTypeToTotalSpots[typ] = append(pl.vehicleTypeToTotalSpots[typ], &spot)
	}
}
