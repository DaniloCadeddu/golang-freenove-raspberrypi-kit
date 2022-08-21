package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	pin = rpio.Pin(19)
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatalf("Error opening gpio")
	}
	defer rpio.Close()

	pin.Output()
	for {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
