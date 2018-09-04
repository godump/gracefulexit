# exit - Exit handlers

The `exit` module defines a single function to register cleanup functions. Functions thus registered are automatically executed upon normal interpreter termination. atexit runs these functions in the reverse order in which they were registered; if you register A, B, and C, at interpreter termination time they will be run in the order C, B, A.

Note: The functions registered via this module are not called when the program is killed by a signal not handled by `os.Interrupt`(usually CTRL+C), when a Go fatal internal error is detected.

# Installation

```sh
$ go get github.com/mohanson/exit
```

# Example

The following simple example demonstrates how a module can initialize a loop and exit it gracefully.

```go
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
```
