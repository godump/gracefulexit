package main

import (
	"log"
	"time"

	"github.com/godump/gracefulexit"
)

func main() {
	chanExit := gracefulexit.Chan()
	for {
		if len(chanExit) != 0 {
			break
		}
		log.Println("wait")
		time.Sleep(time.Second)
	}
	log.Println("done")
}
