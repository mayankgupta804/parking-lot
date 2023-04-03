package parking_test

import (
	"parking-lot/internal/domain"
	"parking-lot/internal/usecases/parking"
	"testing"
	"time"
)

func TestMallFeeModel_GetParkingFee_Motorcycle_59_minutes_59_seconds(t *testing.T) {
	mallFeeModel := parking.Mall()
	got := mallFeeModel.GetParkingFee(domain.Motorcycle, time.Now().Add(-59*time.Minute-59*time.Second), time.Now())
	expected := 10.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestMallFeeModel_GetParkingFee_CarOrSuv_60_minutes(t *testing.T) {
	mallFeeModel := parking.Mall()
	got := mallFeeModel.GetParkingFee(domain.CarOrSuv, time.Now().Add(-60*time.Minute), time.Now())
	expected := 40.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestStadiumFeeModel_GetParkingFee_Motorcycle_3_hours_59_minutes_59_seconds(t *testing.T) {
	stadiumFeeModel := parking.Stadium()
	got := stadiumFeeModel.GetParkingFee(domain.Motorcycle, time.Now().Add(-59*time.Minute-59*time.Second), time.Now())
	expected := 30.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestStadiumFeeModel_GetParkingFee_CarOrSuv_11_hours_59_minutes_59_seconds(t *testing.T) {
	stadiumFeeModel := parking.Stadium()
	got := stadiumFeeModel.GetParkingFee(domain.CarOrSuv, time.Now().Add(-11*time.Hour-59*time.Minute-59*time.Second), time.Now())
	expected := 180.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestStadiumFeeModel_GetParkingFee_Motorcycle_12_hours_59_minutes_59_seconds(t *testing.T) {
	stadiumFeeModel := parking.Stadium()
	got := stadiumFeeModel.GetParkingFee(domain.Motorcycle, time.Now().Add(-12*time.Hour-59*time.Minute-59*time.Second), time.Now())
	expected := 190.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestStadiumFeeModel_GetParkingFee_CarOrSuv_14_hours_59_minutes_59_seconds(t *testing.T) {
	stadiumFeeModel := parking.Stadium()
	got := stadiumFeeModel.GetParkingFee(domain.CarOrSuv, time.Now().Add(-14*time.Hour-59*time.Minute-59*time.Second), time.Now())
	expected := 780.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestAirportFeeModel_GetParkingFee_Motorcycle_59_minutes_59_seconds(t *testing.T) {
	airportFeeModel := parking.Airport()
	got := airportFeeModel.GetParkingFee(domain.Motorcycle, time.Now().Add(-59*time.Minute-59*time.Second), time.Now())
	expected := 0.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestAirportFeeModel_GetParkingFee_Motorcycle_7_hours_59_minutes_59_seconds(t *testing.T) {
	airportFeeModel := parking.Airport()
	got := airportFeeModel.GetParkingFee(domain.Motorcycle, time.Now().Add(-7*time.Hour-59*time.Minute-59*time.Second), time.Now())
	expected := 40.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestAirportFeeModel_GetParkingFee_Motorcycle_23_hours_59_minutes_59_seconds(t *testing.T) {
	airportFeeModel := parking.Airport()
	got := airportFeeModel.GetParkingFee(domain.Motorcycle, time.Now().Add(-23*time.Hour-59*time.Minute-59*time.Second), time.Now())
	expected := 60.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestAirportFeeModel_GetParkingFee_CarOrSuv_1_day(t *testing.T) {
	airportFeeModel := parking.Airport()
	got := airportFeeModel.GetParkingFee(domain.CarOrSuv, time.Now().Add(-24*time.Hour), time.Now())
	expected := 100.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}

func TestAirportFeeModel_GetParkingFee_CarOrSuv_1_day_23_hours_59_minutes_59_seconds(t *testing.T) {
	airportFeeModel := parking.Airport()
	got := airportFeeModel.GetParkingFee(domain.CarOrSuv, time.Now().Add(-47*time.Hour-59*time.Minute-59*time.Second), time.Now())
	expected := 200.0

	if got != expected {
		t.Errorf("Expected: %v. Got: %v", expected, got)
	}
}
