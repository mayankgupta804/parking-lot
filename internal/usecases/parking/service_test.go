package parking_test

import (
	"errors"
	"parking-lot/internal/domain"
	"parking-lot/internal/usecases/parking"
	"testing"
	"time"
)

func TestParkingService_GetTicket(t *testing.T) {
	spotNumber := 1
	mockParkingLot := &domain.ParkingLotMock{
		ParkFunc: func(vehicleType domain.VehicleType) (int, error) {
			return spotNumber, nil
		},
	}
	mockFeeService := &parking.FeeServiceMock{
		GetParkingFeeFunc: func(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
			return 10.0
		},
	}
	parkingService := parking.NewService(parking.WithParkingFeeModel(mockFeeService), parking.WithParkingLotImplementation(mockParkingLot))
	ticket, err := parkingService.GetTicket(parking.CarOrSuv, time.Now().Add(-3*time.Hour))

	if err != nil {
		t.Errorf("Unexpected error: %v", ticket)
	}

	if spotNumber != ticket.SpotNumber {
		t.Errorf("Expected: %d. Got: %d", spotNumber, ticket.SpotNumber)
	}

	if ticket.VehicleType != domain.VehicleType(parking.CarOrSuv) {
		t.Errorf("Expected: %s. Got: %s", parking.CarOrSuv, ticket.VehicleType)
	}

	if ticket.Number != 1 {
		t.Errorf("Expected: %d. Got: %d", 1, ticket.Number)
	}
}

func TestParkingService_GetTicket_ParkingError_NoParkingSpotsAvailable(t *testing.T) {
	vehicleType := domain.CarOrSuv
	expectedUsecaseErr := parking.ParkingServiceErr{Err: parking.ErrParking}
	expectedDomainErr := domain.ErrNoParkingSpotsAvailable(vehicleType)
	mockParkingLot := &domain.ParkingLotMock{
		ParkFunc: func(vehicleType domain.VehicleType) (int, error) {
			return 0, expectedDomainErr
		},
	}
	mockFeeService := &parking.FeeServiceMock{
		GetParkingFeeFunc: func(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
			return 0.0
		},
	}
	parkingService := parking.NewService(parking.WithParkingFeeModel(mockFeeService), parking.WithParkingLotImplementation(mockParkingLot))
	_, err := parkingService.GetTicket(parking.CarOrSuv, time.Now().Add(-3*time.Hour))

	if !errors.As(err, &expectedUsecaseErr) {
		t.Errorf("Expected: %v. Got: %v", expectedUsecaseErr, err)
	}

	actualDomainErr := errors.Unwrap(err)

	if actualDomainErr != expectedDomainErr {
		t.Errorf("Expected: %v. Got: %v", expectedDomainErr, actualDomainErr)
	}
}

func TestParkingService_GetTicket_ParkingError_ParkingLotFull(t *testing.T) {
	expectedUsecaseErr := parking.ParkingServiceErr{Err: parking.ErrParking}
	expectedDomainErr := domain.ErrParkingLotFull
	mockParkingLot := &domain.ParkingLotMock{
		ParkFunc: func(vehicleType domain.VehicleType) (int, error) {
			return 0, expectedDomainErr
		},
	}
	mockFeeService := &parking.FeeServiceMock{
		GetParkingFeeFunc: func(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
			return 0.0
		},
	}
	parkingService := parking.NewService(parking.WithParkingFeeModel(mockFeeService), parking.WithParkingLotImplementation(mockParkingLot))
	_, err := parkingService.GetTicket(parking.CarOrSuv, time.Now().Add(-3*time.Hour))

	if !errors.As(err, &expectedUsecaseErr) {
		t.Errorf("Expected: %v. Got: %v", expectedUsecaseErr, err)
	}

	if expectedUsecaseErr.Error() != err.Error() {
		t.Errorf("Expected: %v. Got: %v", expectedUsecaseErr.Error(), err.Error())
	}
}

func TestParkingService_GetReceipt(t *testing.T) {
	expectedFee := 10.0
	expectedReceiptNumber := 1
	mockParkingLot := &domain.ParkingLotMock{
		UnparkFunc: func(spotNumber int) error {
			return nil
		},
	}
	mockFeeService := &parking.FeeServiceMock{
		GetParkingFeeFunc: func(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
			return expectedFee
		},
	}
	parkingService := parking.NewService(parking.WithParkingFeeModel(mockFeeService), parking.WithParkingLotImplementation(mockParkingLot))
	receipt, err := parkingService.GenerateReceipt(1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if expectedFee != receipt.Fee {
		t.Errorf("Expected: %v. Got: %v", expectedFee, receipt.Fee)
	}

	if receipt.Number != 1 {
		t.Errorf("Expected: %d. Got: %d", expectedReceiptNumber, receipt.Number)
	}
}

func TestParkingService_GetReceipt_VehicleInSpotDoesNotExist(t *testing.T) {
	expectedFee := 0.0
	expectedUsecaseErr := parking.ParkingServiceErr{Err: parking.ErrParking}
	mockParkingLot := &domain.ParkingLotMock{
		UnparkFunc: func(spotNumber int) error {
			return expectedUsecaseErr
		},
	}
	mockFeeService := &parking.FeeServiceMock{
		GetParkingFeeFunc: func(vehicleType domain.VehicleType, entryTime, exitTime time.Time) float64 {
			return expectedFee
		},
	}
	parkingService := parking.NewService(parking.WithParkingFeeModel(mockFeeService), parking.WithParkingLotImplementation(mockParkingLot))
	_, err := parkingService.GenerateReceipt(1)

	if !errors.As(err, &expectedUsecaseErr) {
		t.Errorf("Expected: %v. Got: %v", expectedUsecaseErr, err)
	}

	if expectedUsecaseErr.Error() != err.Error() {
		t.Errorf("Expected: %v. Got: %v", expectedUsecaseErr.Error(), err.Error())
	}
}
