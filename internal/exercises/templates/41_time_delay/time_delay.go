package timedelay

// WaitFor pauses the program for the given duration (in milliseconds).
// Use time.Sleep with time.Millisecond conversion.
func WaitFor(ms int) {
	// TODO: implement using time.Sleep
}

// NotifyAfter waits for the given duration (in milliseconds)
// and then sends true on a channel.
// Use time.After to create a delay.
func NotifyAfter(ms int) chan bool {
	// TODO: implement simple channel send after delay
	return nil
}
