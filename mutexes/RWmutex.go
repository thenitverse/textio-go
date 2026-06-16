package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}
func main() {
	// Initialize the safeCounter with a map and a pointer to an RWMutex
	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.RWMutex{},
	}

	var wg sync.WaitGroup
	key := "clicks"
	iterations := 1000

	// Launch multiple goroutines to concurrently increment the counter
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			sc.inc(key)
			wg.Done()
		}()
	}

	// Launch concurrent readers to read the value while writes are happening
	for i := 0; i < 5; i++ {
		go func(readerID int) {
			for j := 0; j < 10; j++ {
				// Multiple readers can safely access this concurrently using RLock
				_ = sc.val(key)
				time.Sleep(time.Microsecond)
			}
		}(i)
	}

	// Wait for all writers to finish
	wg.Wait()

	// Final verification
	finalVal := sc.val(key)
	if finalVal == iterations {
		println("Success! Expected:", iterations, "Got:", finalVal)
	} else {
		println("Failure! Expected:", iterations, "Got:", finalVal)
	}
}

//Success! Expected: 1000 Got: 1000
