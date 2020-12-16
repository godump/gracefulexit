package gracefulexit

import (
	"os"
	"os/signal"
	"sync/atomic"
)

var (
	geiger uint32
)

// Come return true if the SIGTERM/SIGINT has been received.
func Come() bool {
	return atomic.LoadUint32(&geiger) != 0
}

// Wait for signals.
func Wait() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	<-c
	atomic.AddUint32(&geiger, 1)
	go func() {
		for {
			<-c
			if atomic.AddUint32(&geiger, 1) == 4 {
				signal.Stop(c)
				break
			}
		}
	}()
}
