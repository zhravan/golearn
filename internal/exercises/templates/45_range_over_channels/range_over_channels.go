package range_over_channels

// Assume case where you receive stream of events
// through a channel till it is closed.
// And you process those events as they come through.

func getEvents(ch chan<- int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

// TODO: 1. Range over the channel, calculate the sum from events and return it.
func processStream() int {
	ch := make(chan int)
	go getEvents(ch)

	return 0
}
