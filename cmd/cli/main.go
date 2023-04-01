package main

import (
	"parking-lot/internal/usecases/parking"

	"fmt"
	"log"
	"time"
)

func main() {
	mallParkingService := parking.NewService(
		parking.WithSpots(parking.BusOrTruck, 10),
		parking.WithFeeModel(parking.Mall()))
	ticket, err := mallParkingService.GetTicket(parking.BusOrTruck, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ticket.Number)
	log.Println(ticket.SpotNumber)
	log.Println(ticket.EntryDateTime)
	log.Println(ticket.VehicleType)

	receipt, err := mallParkingService.GenerateReceipt(ticket.Number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt)

}
