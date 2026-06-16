package main

import "fmt"

func getLast[T any](s []T) T {
	var myZero T
	if len(s) == 0 {
		return myZero
	}

	last := s[len(s)-1]
	return last
}

func main() {
	// Test Case 1: A slice of integers
	intSlice := []int{10, 20, 30, 42}
	fmt.Printf("Last int: %v\n", getLast(intSlice))

	// Test Case 2: A slice of strings
	stringSlice := []string{"wizard", "bear", "boots"}
	fmt.Printf("Last string: %v\n", getLast(stringSlice))

	// Test Case 3: An empty slice of booleans (should return zero value: false)
	emptyBoolSlice := []bool{}
	fmt.Printf("Empty slice zero value: %v\n", getLast(emptyBoolSlice))
}

/*output:
Last int: 42
Last string: boots
Empty slice zero value: false*/
