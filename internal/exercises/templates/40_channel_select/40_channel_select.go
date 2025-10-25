package channel_select

import (
	"math/rand/v2"
	"time"
)

// Use case: Network request timeouts

func sendRequest(ch chan<- string) {
	sleepFor := rand.IntN(10) + 1
	time.Sleep(time.Duration(sleepFor) * time.Microsecond)
	ch <- "Success!"
}

// TODO: Use select operator to implement a timeout
// a network request should timeout after 5 microseconds.
func NetworkRequest() string {
	ch := make(chan string, 1)

	go sendRequest(ch)

	var result string

	return result
}
