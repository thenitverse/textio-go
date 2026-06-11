package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	if status == "" {
		return errors.New("status cannot be empty")
	}
	if len(status) > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil
}

func main() {
	// Test case 1: Empty string
	err1 := validateStatus("")
	fmt.Printf("Test 1 (Empty): %v\n", err1)

	// Test case 2: Valid string
	err2 := validateStatus("Hello, world!")
	fmt.Printf("Test 2 (Valid): %v\n", err2)

	// Test case 3: Too long
	longStatus := "This is a very long status update that is designed to exceed the one hundred and forty character limit that we have strictly imposed on our users for this specific exercise."
	err3 := validateStatus(longStatus)
	fmt.Printf("Test 3 (Too long): %v\n", err3)
}

/*output;
Test 1 (Empty): status cannot be empty
Test 2 (Valid): <nil>
Test 3 (Too long): status exceeds 140 characters
*/
