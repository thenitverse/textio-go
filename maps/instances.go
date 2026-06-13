package main

import "fmt"

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, message := range messagedUsers {
		if _, ok := validUsers[message]; ok {
			validUsers[message] = validUsers[message] + 1
		}
	}
}

func main() {
	// Test Case 1: Standard users
	users1 := []string{"alice", "bob", "alice", "charlie"}
	valid1 := map[string]int{"alice": 0, "bob": 0}
	updateCounts(users1, valid1)
	fmt.Printf("Test 1 Result: %v\n", valid1)
	// Expected: map[alice:2 bob:1] (charlie is ignored)

	// Test Case 2: No users found
	users2 := []string{"guest", "anonymous"}
	valid2 := map[string]int{"admin": 0}
	updateCounts(users2, valid2)
	fmt.Printf("Test 2 Result: %v\n", valid2)
	// Expected: map[admin:0]

	// Test Case 3: Empty slice
	users3 := []string{}
	valid3 := map[string]int{"alice": 0}
	updateCounts(users3, valid3)
	fmt.Printf("Test 3 Result: %v\n", valid3)
	// Expected: map[alice:0]
}

/*output:
Test 1 Result: map[alice:2 bob:1]
Test 2 Result: map[admin:0]
Test 3 Result: map[alice:0]*/
