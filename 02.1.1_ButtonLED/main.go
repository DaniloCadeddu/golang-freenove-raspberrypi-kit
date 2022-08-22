package main

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	ledPin    = rpio.Pin(17)
	buttonPin = rpio.Pin(18)
)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatalf("Error opening gpio")
	}
	defer rpio.Close()

	ledPin.Output()
	for {
		if buttonPin.Read() == rpio.Low {
			ledPin.High()
		} else {
			ledPin.Low()
		}
	}

}
