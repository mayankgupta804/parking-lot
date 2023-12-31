// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package parking

import (
	"parking-lot/internal/domain"
	"sync"
	"time"
)

// Ensure, that FeeServiceMock does implement FeeService.
// If this is not the case, regenerate this file with moq.
var _ FeeService = &FeeServiceMock{}

// FeeServiceMock is a mock implementation of FeeService.
//
//	func TestSomethingThatUsesFeeService(t *testing.T) {
//
//		// make and configure a mocked FeeService
//		mockedFeeService := &FeeServiceMock{
//			GetParkingFeeFunc: func(vehicleType domain.VehicleType, entryTime time.Time, exitTime time.Time) float64 {
//				panic("mock out the GetParkingFee method")
//			},
//		}
//
//		// use mockedFeeService in code that requires FeeService
//		// and then make assertions.
//
//	}
type FeeServiceMock struct {
	// GetParkingFeeFunc mocks the GetParkingFee method.
	GetParkingFeeFunc func(vehicleType domain.VehicleType, entryTime time.Time, exitTime time.Time) float64

	// calls tracks calls to the methods.
	calls struct {
		// GetParkingFee holds details about calls to the GetParkingFee method.
		GetParkingFee []struct {
			// VehicleType is the vehicleType argument value.
			VehicleType domain.VehicleType
			// EntryTime is the entryTime argument value.
			EntryTime time.Time
			// ExitTime is the exitTime argument value.
			ExitTime time.Time
		}
	}
	lockGetParkingFee sync.RWMutex
}

// GetParkingFee calls GetParkingFeeFunc.
func (mock *FeeServiceMock) GetParkingFee(vehicleType domain.VehicleType, entryTime time.Time, exitTime time.Time) float64 {
	if mock.GetParkingFeeFunc == nil {
		panic("FeeServiceMock.GetParkingFeeFunc: method is nil but FeeService.GetParkingFee was just called")
	}
	callInfo := struct {
		VehicleType domain.VehicleType
		EntryTime   time.Time
		ExitTime    time.Time
	}{
		VehicleType: vehicleType,
		EntryTime:   entryTime,
		ExitTime:    exitTime,
	}
	mock.lockGetParkingFee.Lock()
	mock.calls.GetParkingFee = append(mock.calls.GetParkingFee, callInfo)
	mock.lockGetParkingFee.Unlock()
	return mock.GetParkingFeeFunc(vehicleType, entryTime, exitTime)
}

// GetParkingFeeCalls gets all the calls that were made to GetParkingFee.
// Check the length with:
//
//	len(mockedFeeService.GetParkingFeeCalls())
func (mock *FeeServiceMock) GetParkingFeeCalls() []struct {
	VehicleType domain.VehicleType
	EntryTime   time.Time
	ExitTime    time.Time
} {
	var calls []struct {
		VehicleType domain.VehicleType
		EntryTime   time.Time
		ExitTime    time.Time
	}
	mock.lockGetParkingFee.RLock()
	calls = mock.calls.GetParkingFee
	mock.lockGetParkingFee.RUnlock()
	return calls
}
