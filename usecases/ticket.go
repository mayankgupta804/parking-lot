package usecases

import (
	"parking-lot/internal/domain"
	"time"
)

type Ticket struct {
	Number        int
	SpotNumber    int
	VehicleType   domain.VehicleType
	EntryDateTime time.Time
}
