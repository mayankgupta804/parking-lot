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
