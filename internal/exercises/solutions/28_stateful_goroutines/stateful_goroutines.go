package stateful_goroutines

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
	c := &Counter{
		reads:  make(chan readOp),
		writes: make(chan writeOp),
	}
	
	// Start the state-owning goroutine
	go func() {
		var state int
		for {
			select {
			case read := <-c.reads:
				read.resp <- state
			case write := <-c.writes:
				state += write.amount
				write.resp <- true
			}
		}
	}()
	
	return c
}

// Increment increments the counter by the given amount
func (c *Counter) Increment(amount int) {
	write := writeOp{
		amount: amount,
		resp:   make(chan bool),
	}
	c.writes <- write
	<-write.resp
}

// GetValue returns the current counter value
func (c *Counter) GetValue() int {
	read := readOp{
		resp: make(chan int),
	}
	c.reads <- read
	return <-read.resp
}
