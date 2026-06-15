package main

import "fmt"

func countReports(numSentCh chan int) int {
	total := 0
	for {
		v, ok := <-numSentCh
		if !ok {
			break
		}
		total += v
	}
	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func main() {
	tests := []struct {
		numBatches int
		expected   int
	}{
		{3, 114},
		{4, 198},
		{0, 0},
		{1, 15},
		{6, 435},
	}

	for _, test := range tests {
		ch := make(chan int)
		go sendReports(test.numBatches, ch)
		result := countReports(ch)
		status := "PASS"
		if result != test.expected {
			status = "FAIL"
		}
		fmt.Printf("%s | numBatches: %d | expected: %d | actual: %d\n", status, test.numBatches, test.expected, result)
	}
}

/*output:
PASS | numBatches: 3 | expected: 114 | actual: 114
PASS | numBatches: 4 | expected: 198 | actual: 198
PASS | numBatches: 0 | expected: 0 | actual: 0
PASS | numBatches: 1 | expected: 15 | actual: 15
PASS | numBatches: 6 | expected: 435 | actual: 435*/
