package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		var err error = errors.New("no dividing by 0")
		return 0.0, err
	}
	return x / y, nil
}

func main() {
	// Test Case 1: Normal division
	res, err := divide(10, 2)
	fmt.Printf("Result: %v, Error: %v\n", res, err)
	// Expected: Result: 5, Error: <nil>

	// Test Case 2: Division by zero
	res, err = divide(10, 0)
	fmt.Printf("Result: %v, Error: %v\n", res, err)
	// Expected: Result: 0, Error: no dividing by 0
}

/*output:
Result: 5, Error: <nil>
Result: 0, Error: no dividing by 0
*/
