package domain_test

import (
	"errors"
	"parking-lot/internal/domain"
	"testing"
)

func TestParkingLot_Park_SpotAvailableToPark(t *testing.T) {
	pl := domain.NewParkingLot()
	pl.AddSpots(domain.Motorcycle, 1)
	spotNumber, err := pl.Park(domain.Motorcycle)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 1

	if spotNumber != expected {
		t.Errorf("Expected: %d. Got: %d", expected, spotNumber)
	}
}

func TestParkingLot_Park_SpotUnavailableToPark(t *testing.T) {
	pl := domain.NewParkingLot()
	pl.AddSpots(domain.Motorcycle, 1)
	spotNumber, err := pl.Park(domain.Motorcycle)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 1

	if spotNumber != expected {
		t.Errorf("Expected: %d. Got: %d", expected, spotNumber)
	}

	spotNumber, err = pl.Park(domain.Motorcycle)

	expectedInternalErr := domain.ErrParkingLotFull
	expectedErr := domain.ParkingLotError{Err: expectedInternalErr}

	expectedSpotNumber := 0

	if spotNumber != expectedSpotNumber {
		t.Errorf("Expected: %d. Got: %d", expectedSpotNumber, spotNumber)
	}

	if !errors.As(err, &expectedErr) {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr, err)
	}

	if expectedInternalErr.Error() != err.Error() {
		t.Errorf("Expected error: %v. Got error: %v", expectedInternalErr.Error(), err.Error())
	}
}

func TestParkingLot_Park_NoParkingSpotsAvailable(t *testing.T) {
	pl := domain.NewParkingLot()
	vehicleType := domain.Motorcycle
	spotNumber, err := pl.Park(vehicleType)

	expectedInternalErr := domain.ErrNoParkingSpotsAvailable(vehicleType)
	expectedErr := domain.ParkingLotError{Err: expectedInternalErr}

	expectedSpotNumber := 0

	if spotNumber != expectedSpotNumber {
		t.Errorf("Expected: %d. Got: %d", expectedSpotNumber, spotNumber)
	}

	if !errors.As(err, &expectedErr) {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr, err)
	}

	if expectedInternalErr.Error() != err.Error() {
		t.Errorf("Expected error: %v. Got error: %v", expectedInternalErr.Error(), err.Error())
	}
}

func TestParkingLot_Unpark_Successful(t *testing.T) {
	pl := domain.NewParkingLot()
	pl.AddSpots(domain.Motorcycle, 1)
	spotNumber, err := pl.Park(domain.Motorcycle)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expected := 1

	if spotNumber != expected {
		t.Errorf("Expected: %d. Got: %d", expected, spotNumber)
	}

	err = pl.Unpark(spotNumber)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestParkingLot_Unpark_InvalidSpotNumber(t *testing.T) {
	pl := domain.NewParkingLot()
	spotNumber := 1

	err := pl.Unpark(spotNumber)
	expectedInternalErr := domain.ErrVehicleInSpotDoesNotExist
	expectedErr := domain.ParkingLotError{Err: expectedInternalErr}

	if !errors.As(err, &expectedErr) {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr, err)
	}

	if err.Error() != expectedInternalErr.Error() {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr.Error(), err.Error())
	}
}
