package stateful_goroutines

import (
	"sync"
	"testing"
	"time"
)

func TestCounterInitialization(t *testing.T) {
	counter := NewCounter()
	if counter == nil {
		t.Fatal("NewCounter() returned nil")
	}
	
	// Give goroutine time to start
	time.Sleep(10 * time.Millisecond)
	
	value := counter.GetValue()
	if value != 0 {
		t.Errorf("Initial counter value = %d, want 0", value)
	}
}

func TestCounterIncrement(t *testing.T) {
	counter := NewCounter()
	
	counter.Increment(5)
	counter.Increment(3)
	
	value := counter.GetValue()
	if value != 8 {
		t.Errorf("Counter value = %d, want 8", value)
	}
}

func TestCounterConcurrentIncrements(t *testing.T) {
	counter := NewCounter()
	
	var wg sync.WaitGroup
	numGoroutines := 100
	incrementsPerGoroutine := 10
	
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment(1)
			}
		}()
	}
	
	wg.Wait()
	
	expected := numGoroutines * incrementsPerGoroutine
	value := counter.GetValue()
	if value != expected {
		t.Errorf("Counter value = %d, want %d", value, expected)
	}
}

func TestCounterConcurrentReadsAndWrites(t *testing.T) {
	counter := NewCounter()
	
	var wg sync.WaitGroup
	numReaders := 50
	numWriters := 50
	
	// Start writers
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				counter.Increment(1)
				time.Sleep(time.Microsecond)
			}
		}()
	}
	
	// Start readers
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				_ = counter.GetValue()
				time.Sleep(time.Microsecond)
			}
		}()
	}
	
	wg.Wait()
	
	// Verify final value
	expected := numWriters * 5
	value := counter.GetValue()
	if value != expected {
		t.Errorf("Counter value = %d, want %d", value, expected)
	}
}

func TestCounterNegativeIncrement(t *testing.T) {
	counter := NewCounter()
	
	counter.Increment(10)
	counter.Increment(-3)
	
	value := counter.GetValue()
	if value != 7 {
		t.Errorf("Counter value = %d, want 7", value)
	}
}
