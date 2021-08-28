package main

import (
	"log"

	"github.com/godump/gracefulexit"
)

func main() {
	log.Println("wait for signal")
	gracefulexit.Wait()
	log.Println("done")
}
