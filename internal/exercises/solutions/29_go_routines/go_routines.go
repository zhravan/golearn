package go_routines

import (
	"time"
)

func sleepForOneHundredMillisecond() {
	time.Sleep(100 * time.Millisecond)
}

func sleepForTwoHundredMilliseconds() {
	time.Sleep(200 * time.Millisecond)
}

func RunConcurrently() {
	go sleepForOneHundredMillisecond()
	go sleepForTwoHundredMilliseconds()

	// Alternative solution using anonymous go routines
	// go func() {
	// 	sleepForOneHundredMillisecond()
	// }()
	// go func() {
	// 	sleepForTwoHundredMilliseconds()
	// }()
}
