package usecases

import (
	"parking-lot/internal/domain"
	"time"
)

type ParkingService interface {
	GetParkingTicket(vehicleType domain.VehicleType, entryDateTime time.Time) (Ticket, error)
	GenerateParkingReceipt(ticketNumber int) (Receipt, error)
}

type MallParkingService struct {
	currentTicketNumber      int
	currentReceiptNumber     int
	ticketNumberToSpotNumber map[int]int
	ticketNumberToTicket     map[int]Ticket
	domain.FeeService
	domain.ParkingLot
}

func NewMallParkingService(options ...func(*MallParkingService)) ParkingService {
	mlp := &MallParkingService{}

	mlp.ticketNumberToSpotNumber = make(map[int]int)
	mlp.ticketNumberToTicket = make(map[int]Ticket)

	mlp.ParkingLot = domain.NewParkingLot()
	mlp.FeeService = domain.DefaultMallFeeModel()

	for _, opt := range options {
		opt(mlp)
	}

	return mlp
}

func WithFeeModel(feeService domain.FeeService) func(*MallParkingService) {
	return func(mlp *MallParkingService) {
		mlp.FeeService = feeService
	}
}

func WithSpots(vehicleType domain.VehicleType, size int) func(*MallParkingService) {
	return func(mlp *MallParkingService) {
		if mlp.ParkingLot == nil {
			mlp.ParkingLot = domain.NewParkingLot()
		}
		mlp.AddSpots(vehicleType, size)
	}
}

func (mallParkingLot *MallParkingService) GetParkingTicket(vehicleType domain.VehicleType, entryDateTime time.Time) (Ticket, error) {
	ticket := Ticket{}
	spotNumber, err := mallParkingLot.Park(vehicleType)
	if err != nil {
		// handle domain error
		return ticket, nil
	}
	currentTicketNumber := mallParkingLot.currentTicketNumber

	ticket.SpotNumber = spotNumber
	ticket.EntryDateTime = entryDateTime
	ticket.VehicleType = vehicleType
	// use mutex if op will be concurrent
	ticket.Number = currentTicketNumber + 1

	mallParkingLot.currentTicketNumber += 1
	mallParkingLot.ticketNumberToSpotNumber[ticket.Number] = spotNumber
	mallParkingLot.ticketNumberToTicket[ticket.Number] = ticket
	return ticket, nil
}

func (mallParkingLot *MallParkingService) GenerateParkingReceipt(ticketNumber int) (Receipt, error) {
	receipt := Receipt{}
	spotNumber := mallParkingLot.ticketNumberToSpotNumber[ticketNumber]

	if err := mallParkingLot.Unpark(spotNumber); err != nil {
		return receipt, err
	}

	ticket := mallParkingLot.ticketNumberToTicket[ticketNumber]

	entryDateTime := ticket.EntryDateTime
	exitDateTime := time.Now()

	fee := mallParkingLot.GetParkingFee(ticket.VehicleType, entryDateTime, exitDateTime)
	receipt.EntryDateTime = entryDateTime
	receipt.ExitDateTime = exitDateTime
	receipt.Fee = fee
	receipt.Number = mallParkingLot.currentReceiptNumber + 1
	mallParkingLot.currentReceiptNumber += 1

	return receipt, nil
}
