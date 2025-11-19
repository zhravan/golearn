package range_over_channels

func getEvents(ch chan<- int) {
	for i := range 10 {
		ch <- i
	}
	close(ch)
}

func processStream() int {
	ch := make(chan int)
	go getEvents(ch)

	sum := 0
	for event := range ch {
		sum += event
	}
	return sum
}
