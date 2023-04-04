package domain_test

import (
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

	expectedErr := domain.ErrParkingLotFull

	expectedSpotNumber := 0

	if spotNumber != expectedSpotNumber {
		t.Errorf("Expected: %d. Got: %d", expectedSpotNumber, spotNumber)
	}

	if expectedErr != err {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr, err)
	}
}

func TestParkingLot_Park_NoParkingSpotsAvailable(t *testing.T) {
	pl := domain.NewParkingLot()
	vehicleType := domain.Motorcycle
	spotNumber, err := pl.Park(vehicleType)

	expectedErr := domain.ErrNoParkingSpotsAvailable(vehicleType)

	expectedSpotNumber := 0

	if spotNumber != expectedSpotNumber {
		t.Errorf("Expected: %d. Got: %d", expectedSpotNumber, spotNumber)
	}

	if expectedErr.Error() != err.Error() {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr, err)
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
	expectedErr := domain.ErrVehicleInSpotDoesNotExist

	if err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v. Got error: %v", expectedErr, err)
	}
}
