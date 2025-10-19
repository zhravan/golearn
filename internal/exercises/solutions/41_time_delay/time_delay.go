package timedelay

import "time"

func WaitFor(ms int) {
	duration := time.Duration(ms) * time.Millisecond
	time.Sleep(duration)
}

func NotifyAfter(ms int) chan bool {
	ch := make(chan bool, 1)
	duration := time.Duration(ms) * time.Millisecond

	go func() {
		time.Sleep(duration)
		ch <- true
	}()

	return ch
}
