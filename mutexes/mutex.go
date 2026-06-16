package main

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.slowVal(key)
}

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}

func (sc safeCounter) slowVal(key string) int {
	time.Sleep(time.Microsecond)
	return sc.counts[key]
}
func main() {
	// Initialize our safe counter with an empty map and a new Mutex
	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.Mutex{},
	}

	var wg sync.WaitGroup
	email := "apprentice@boot.dev"
	incrementCount := 100

	// Launch multiple concurrent goroutines to increment the same key
	for i := 0; i < incrementCount; i++ {
		wg.Add(1)
		go func() {
			sc.inc(email)
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Retrieve the final value
	finalVal := sc.val(email)

	// Verify the result
	if finalVal == incrementCount {
		println("Success! Count is correct:", finalVal)
	} else {
		println("Failure! Expected", incrementCount, "but got", finalVal)
	}
}

//Success! Count is correct: 100
