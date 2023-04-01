package domain

type Spot struct {
	number   int
	typ      VehicleType
	occupied bool
}

func NewSpot(number int, typ VehicleType) Spot {
	spot := Spot{}
	spot.number = number
	spot.typ = typ
	spot.occupied = false
	return spot
}
