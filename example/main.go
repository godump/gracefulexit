package main

import (
	"log"
	"time"

	"github.com/mohanson/gracefulexit"
)

func main() {
	log.Println("wait for signal")
	gracefulexit.Wait()
	log.Println("wait 10 seconds")
	time.Sleep(time.Second * 10)
	log.Println("done")
}
