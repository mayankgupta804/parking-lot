package integration_test

import (
	"parking-lot/internal/usecases/parking"
	"testing"
	"time"
)

func TestAirportParking_CarOrSuv_For_4_Days(t *testing.T) {
	airportParkingService := parking.NewService(
		parking.WithParkingSpots(parking.CarOrSuv, 10),
		parking.WithParkingFeeModel(parking.Airport()))

	ticket, err := airportParkingService.GetTicket(parking.CarOrSuv, time.Now().Add(-73*time.Hour))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	receipt, err := airportParkingService.GenerateReceipt(ticket.Number)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFee := 400.0

	if receipt.Fee != expectedFee {
		t.Errorf("Expected: %v. Got: %v", expectedFee, receipt.Fee)
	}
}

func TestAirportParking_Motorcycle_For_23_Hours(t *testing.T) {
	airportParkingService := parking.NewService(
		parking.WithParkingSpots(parking.Motorcycle, 10),
		parking.WithParkingFeeModel(parking.Airport()))

	ticket, err := airportParkingService.GetTicket(parking.Motorcycle, time.Now().Add(-23*time.Hour))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	receipt, err := airportParkingService.GenerateReceipt(ticket.Number)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFee := 60.0

	if receipt.Fee != expectedFee {
		t.Errorf("Expected: %v. Got: %v", expectedFee, receipt.Fee)
	}
}

func TestStadiumParking_Motorcycle_For_11_Hours_59_Minutes_59_Seconds(t *testing.T) {
	stadiumParkingService := parking.NewService(
		parking.WithParkingSpots(parking.Motorcycle, 10),
		parking.WithParkingFeeModel(parking.Stadium()))

	ticket, err := stadiumParkingService.GetTicket(parking.Motorcycle, time.Now().Add(-11*time.Hour-59*time.Minute-59*time.Second))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	receipt, err := stadiumParkingService.GenerateReceipt(ticket.Number)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFee := 90.0

	if receipt.Fee != expectedFee {
		t.Errorf("Expected: %v. Got: %v", expectedFee, receipt.Fee)
	}
}

func TestStadiumParking_For_CarOrSuv_For_12_Hours(t *testing.T) {
	stadiumParkingService := parking.NewService(
		parking.WithParkingSpots(parking.CarOrSuv, 10),
		parking.WithParkingFeeModel(parking.Stadium()))

	ticket, err := stadiumParkingService.GetTicket(parking.CarOrSuv, time.Now().Add(-12*time.Hour))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	receipt, err := stadiumParkingService.GenerateReceipt(ticket.Number)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFee := 380.0

	if receipt.Fee != expectedFee {
		t.Errorf("Expected: %v. Got: %v", expectedFee, receipt.Fee)
	}
}

func TestMallParking_For_BusOrTruck_For_2_Hours_59_minutes_59_seconds(t *testing.T) {
	mallParkingService := parking.NewService(
		parking.WithParkingSpots(parking.BusOrTruck, 10),
		parking.WithParkingFeeModel(parking.Mall()))

	ticket, err := mallParkingService.GetTicket(parking.BusOrTruck, time.Now().Add(-2*time.Hour-59*time.Minute-59*time.Second))

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	receipt, err := mallParkingService.GenerateReceipt(ticket.Number)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedFee := 150.0

	if receipt.Fee != expectedFee {
		t.Errorf("Expected: %v. Got: %v", expectedFee, receipt.Fee)
	}
}
