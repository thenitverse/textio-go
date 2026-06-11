package main

import "fmt"

func countConnections(groupSize int) int {
	connections := 0
	for i := 0; i < groupSize; i++ {
		connections += i

	}
	return connections
}

func main() {
	// Test case 1: 2 people should have 1 connection
	fmt.Printf("Group size 2: expected 1, got %d\n", countConnections(2))

	// Test case 2: 3 people should have 3 connections
	fmt.Printf("Group size 3: expected 3, got %d\n", countConnections(3))

	// Test case 3: 4 people should have 6 connections
	fmt.Printf("Group size 4: expected 6, got %d\n", countConnections(4))

	// Test case 4: 10 people should have 45 connections
	fmt.Printf("Group size 10: expected 45, got %d\n", countConnections(10))
}

/*output:
Group size 2: expected 1, got 1
Group size 3: expected 3, got 3
Group size 4: expected 6, got 6
Group size 10: expected 45, got 45
*/
