package channel_select

import (
	"math/rand/v2"
	"time"
)

func sendRequest(ch chan<- string) {
	sleepFor := rand.IntN(10) + 1
	time.Sleep(time.Duration(sleepFor) * time.Microsecond)
	ch <- "Success!"
}

func NetworkRequest() string {
	ch := make(chan string, 1)

	go sendRequest(ch)

	var result string
	select {
	case <-ch:
		result = "Success!"
	case <-time.After(5 * time.Microsecond):
		result = "Timeout!"
	}

	return result
}
