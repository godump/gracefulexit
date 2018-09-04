package main

import (
	"log"
	"time"

	"github.com/mohanson/exit"
)

func main() {
	exit.Handle(func() { log.Println("Exit function(C)") })
	exit.Handle(func() { log.Println("Exit function(B)") })
	exit.Handle(func() { log.Println("Exit function(A)") })
	// loop forever and exit gracefully when CTRL+C
	for _ = range time.NewTicker(time.Second).C {
		if exit.Come() {
			break
		}
		log.Println("Wait for CTRL+C")
	}
}
