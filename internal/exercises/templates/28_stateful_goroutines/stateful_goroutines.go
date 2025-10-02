package stateful_goroutines

// TODO:
// - Implement a Counter that manages state using a single goroutine and channels.
// - The counter should support Increment and GetValue operations.
// - State must be owned by a single goroutine to avoid race conditions.
// - Other goroutines communicate via channels to read or modify the state.

// readOp represents a read request
type readOp struct {
	resp chan int
}

// writeOp represents a write request (increment)
type writeOp struct {
	amount int
	resp   chan bool
}

type Counter struct {
	reads  chan readOp
	writes chan writeOp
}

// NewCounter creates and starts a new stateful counter
func NewCounter() *Counter {
	// TODO: initialize channels and start the state-owning goroutine
	return &Counter{}
}

// Increment increments the counter by the given amount
func (c *Counter) Increment(amount int) {
	// TODO: send a write operation and wait for confirmation
}

// GetValue returns the current counter value
func (c *Counter) GetValue() int {
	// TODO: send a read operation and return the value
	return 0
}
