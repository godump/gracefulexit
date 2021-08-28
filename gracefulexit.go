package gracefulexit

import (
	"os"
	"os/signal"
)

func Chan() chan os.Signal {
	buffer := make(chan os.Signal, 1)
	signal.Notify(buffer, os.Interrupt, os.Kill)
	return buffer
}

func Wait() {
	<-Chan()
}
