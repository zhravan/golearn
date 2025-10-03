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
	// TODO: update following code
	sleepForOneHundredMillisecond()
	sleepForTwoHundredMilliseconds()
}
