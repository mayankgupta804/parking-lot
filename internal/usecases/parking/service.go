package parking

import (
	"parking-lot/internal/domain"

	"time"
)

type Receipt struct {
	Number        int
	EntryDateTime time.Time
	ExitDateTime  time.Time
	Fee           float64
}

type Ticket struct {
	Number        int
	SpotNumber    int
	VehicleType   domain.VehicleType
	EntryDateTime time.Time
}

type VehicleType string

const (
	Motorcycle VehicleType = VehicleType(domain.Motorcycle)
	CarOrSuv   VehicleType = VehicleType(domain.CarOrSuv)
	BusOrTruck VehicleType = VehicleType(domain.BusOrTruck)
)

type ParkingService interface {
	GetTicket(vehicleType VehicleType, entryDateTime time.Time) (Ticket, error)
	GenerateReceipt(ticketNumber int) (Receipt, error)
}

type parkingService struct {
	currentTicketNumber      int
	currentReceiptNumber     int
	ticketNumberToSpotNumber map[int]int
	ticketNumberToTicket     map[int]Ticket
	domain.ParkingLot
	FeeService
}

func NewService(options ...func(*parkingService)) ParkingService {
	ps := &parkingService{}

	ps.ticketNumberToSpotNumber = make(map[int]int)
	ps.ticketNumberToTicket = make(map[int]Ticket)

	ps.ParkingLot = domain.NewParkingLot()

	for _, opt := range options {
		opt(ps)
	}

	if ps.FeeService == nil {
		ps.FeeService = Mall()
	}

	return ps
}

func WithFeeModel(feeService FeeService) func(*parkingService) {
	return func(ps *parkingService) {
		ps.FeeService = feeService
	}
}

func WithSpots(vehicleType VehicleType, size int) func(*parkingService) {
	return func(ps *parkingService) {
		if ps.ParkingLot == nil {
			ps.ParkingLot = domain.NewParkingLot()
		}
		ps.AddSpots(domain.VehicleType(vehicleType), size)
	}
}

func (ps *parkingService) GetTicket(vehicleType VehicleType, entryDateTime time.Time) (Ticket, error) {
	ticket := Ticket{}
	spotNumber, err := ps.Park(domain.VehicleType(vehicleType))
	if err != nil {
		// handle domain error
		return ticket, nil
	}
	currentTicketNumber := ps.currentTicketNumber

	ticket.SpotNumber = spotNumber
	ticket.EntryDateTime = entryDateTime
	ticket.VehicleType = domain.VehicleType(vehicleType)
	// use mutex if op will be concurrent
	ticket.Number = currentTicketNumber + 1

	ps.currentTicketNumber += 1
	ps.ticketNumberToSpotNumber[ticket.Number] = spotNumber
	ps.ticketNumberToTicket[ticket.Number] = ticket
	return ticket, nil
}

func (ps *parkingService) GenerateReceipt(ticketNumber int) (Receipt, error) {
	receipt := Receipt{}
	spotNumber := ps.ticketNumberToSpotNumber[ticketNumber]

	if err := ps.Unpark(spotNumber); err != nil {
		return receipt, err
	}

	ticket := ps.ticketNumberToTicket[ticketNumber]

	entryDateTime := ticket.EntryDateTime
	exitDateTime := time.Now()

	fee := ps.GetParkingFee(ticket.VehicleType, entryDateTime, exitDateTime)
	receipt.EntryDateTime = entryDateTime
	receipt.ExitDateTime = exitDateTime
	receipt.Fee = fee
	receipt.Number = ps.currentReceiptNumber + 1
	ps.currentReceiptNumber += 1

	return receipt, nil
}