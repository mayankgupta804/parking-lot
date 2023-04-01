package main

import (
	"fmt"
	"log"
	"parking-lot/internal/domain"
	"parking-lot/usecases"
	"time"
)

func main() {
	mallParkingService := usecases.NewMallParkingService(usecases.WithSpots(domain.BusOrTruck, 10))
	ticket, err := mallParkingService.GetParkingTicket(domain.BusOrTruck, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(ticket.Number)
	log.Println(ticket.SpotNumber)
	log.Println(ticket.EntryDateTime)
	log.Println(ticket.VehicleType)

	receipt, err := mallParkingService.GenerateParkingReceipt(ticket.Number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt)
}
