package non_blocking_channel_operations

import (
	"sync"
	"time"
)

// Assume case where you have high throughput system
// where you have to drop logs
// if they cannot be processed in time,
// in order to keep system throughput high.

type Logger struct {
	ch        chan string
	wg        sync.WaitGroup
	processed int
	dropped   int

	mu sync.Mutex
}

func NewLogger(bufferSize int) *Logger {
	l := &Logger{
		ch: make(chan string, bufferSize),
	}
	l.wg.Add(1)
	go l.run()
	return l
}

var processingDelay = 200 * time.Microsecond

// slow log processor
func (l *Logger) run() {
	defer l.wg.Done()
	for msg := range l.ch {
		_ = msg
		time.Sleep(processingDelay)

		l.mu.Lock()
		l.processed++
		l.mu.Unlock()
	}
}

// TODO: Implement logger function in a way
// 1. incoming messages are queued to the channel
// 2. if queue is full, increment dropped
func (l *Logger) Log(msg string) {

}

// graceful shutdown
func (l *Logger) Close() {
	close(l.ch)
	l.wg.Wait()
}

func (l *Logger) Stats() (processed, dropped int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.processed, l.dropped
}
