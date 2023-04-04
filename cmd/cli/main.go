package main

import (
	"parking-lot/internal/usecases/parking"

	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

var headerFmt = color.New(color.FgGreen, color.Underline).SprintfFunc()
var columnFmt = color.New(color.FgYellow).SprintfFunc()

func main() {
	// Mall Parking Service execution code
	mallParkingService := parking.NewService(
		parking.WithParkingSpots(parking.BusOrTruck, 10),
		parking.WithParkingFeeModel(parking.Mall()))
	ticket, err := mallParkingService.GetTicket(parking.BusOrTruck, time.Now())
	if err != nil {
		log.Fatal(err)
	}

	printTicketInfo(ticket)

	receipt, err := mallParkingService.GenerateReceipt(ticket.Number)
	if err != nil {
		log.Fatal(err)
	}
	printReceiptInfo(receipt)

	// Stadium parking service execution code
	stadiumParkingService := parking.NewService(
		parking.WithParkingSpots(parking.Motorcycle, 10),
		parking.WithParkingFeeModel(parking.Stadium()))

	ticket, err = stadiumParkingService.GetTicket(parking.Motorcycle, time.Now().Add(-14*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	printTicketInfo(ticket)

	receipt, err = stadiumParkingService.GenerateReceipt(ticket.Number)
	if err != nil {
		log.Fatal(err)
	}
	printReceiptInfo(receipt)

	// Airport parking service execution code
	airportParkingService := parking.NewService(
		parking.WithParkingSpots(parking.CarOrSuv, 10),
		parking.WithParkingFeeModel(parking.Airport()))

	ticket, err = airportParkingService.GetTicket(parking.CarOrSuv, time.Now().Add(-73*time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	printTicketInfo(ticket)

	receipt, err = airportParkingService.GenerateReceipt(ticket.Number)
	if err != nil {
		log.Fatal(err)
	}
	printReceiptInfo(receipt)
}

func printTicketInfo(ticket parking.Ticket) {
	fmt.Printf("Ticket:\n\n")
	tbl := table.New("No.", "Spot Number", "Entry Time", "Vehicle Type")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.AddRow(ticket.Number, ticket.SpotNumber, ticket.EntryDateTime, ticket.VehicleType)
	tbl.Print()
}

func printReceiptInfo(receipt parking.Receipt) {
	fmt.Printf("\n\nReceipt:\n\n")
	tbl := table.New("No.", "Entry Time", "Exit Time", "Fee")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	tbl.AddRow(receipt.Number, receipt.EntryDateTime, receipt.ExitDateTime, receipt.Fee)
	tbl.Print()
}
