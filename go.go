// The `exit` module defines a single function to register cleanup functions.
// Functions thus registered are automatically executed upon normal interpreter
// termination. atexit runs these functions in the reverse order in which they
// were registered; if you register A, B, and C, at interpreter termination
// time they will be run in the order C, B, A.
//
// Note: The functions registered via this module are not called when the
// program is killed by a signal not handled by `os.Interrupt`, when a Go
// fatal internal error is detected.
package exit

import (
	"os"
	"os/signal"
	"sync"
)

var (
	exitCome     bool
	exitHandlers []func()
	exitWg       sync.WaitGroup
)

// Handle register func as a function to be executed at termination(usually
// CTRL+C). The callback function has no arguments and does not need to return
// anything. Exit runs these functions in the reverse order in which they were
// registered; if you register A, B, and C, at interpreter termination time
// they will be run in the order C, B, A.
func Handle(f func()) {
	exitWg.Add(1)
	exitHandlers = append(exitHandlers, f)
}

// Come return true if the SIGINT has been received.
func Come() bool {
	return exitCome
}

// Wait for all callback functions to complete
func Wait() {
	exitWg.Wait()
}

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		exitCome = true
		size := len(exitHandlers)
		for i := 0; i < size; i++ {
			h := exitHandlers[size-i-1]
			h()
			exitWg.Done()
		}
	}()
}
