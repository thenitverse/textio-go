package main

import "fmt"

func concurrentFib(n int) []int {
	s := make([]int, 0)
	ch := make(chan int)
	go fibonacci(n, ch)
	for val := range ch {
		s = append(s, val)
	}
	return s
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {

	testCases := []int{5, 3, 0, 1, 7}

	for _, n := range testCases {
		result := concurrentFib(n)
		fmt.Printf("fibonacci(%d) -> %v\n", n, result)
	}
}

/*output:
fibonacci(5) -> [0 1 1 2 3]
fibonacci(3) -> [0 1 1]
fibonacci(0) -> []
fibonacci(1) -> [0]
fibonacci(7) -> [0 1 1 2 3 5 8]*/
