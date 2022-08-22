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

	ledState := false // off by default
	ledPin.Output()

	rpio.DetectEdge(buttonPin, rpio.FallEdge)

	for {
		isEventOccurred := rpio.EdgeDetected(buttonPin)
		if isEventOccurred {
			ledState = !ledState
			if ledState {
				ledPin.Low()
			} else {
				ledPin.High()
			}
		}
	}

}
