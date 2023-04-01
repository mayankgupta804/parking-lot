package usecases

import "time"

type Receipt struct {
	Number        int
	EntryDateTime time.Time
	ExitDateTime  time.Time
	Fee           float64
}
