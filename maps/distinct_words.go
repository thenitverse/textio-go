package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]struct{})

	for _, message := range messages {
		lowermessage := strings.ToLower(message)
		words := strings.Fields(lowermessage)
		for _, word := range words {
			distinctWords[word] = struct{}{}
		}
	}
	return len(distinctWords)
}

func main() {
	// Test case 1: Basic distinct words and casing
	test1 := []string{"Hello world", "hello there", "General Kenobi"}
	fmt.Printf("Test 1 - Expected: 5, Got: %d\n", countDistinctWords(test1))

	// Test case 2: Duplicate words with different casing and punctuation attached
	test2 := []string{"LFG UBRS", "lfg ubrs", "LFG Ubrs"}
	fmt.Printf("Test 2 - Expected: 2, Got: %d\n", countDistinctWords(test2))

	// Test case 3: Empty strings
	test3 := []string{"", "   "}
	fmt.Printf("Test 3 - Expected: 0, Got: %d\n", countDistinctWords(test3))

	// Test case 4: Shared words and different formatting
	test4 := []string{"I'm out of range", "I'm out of mana"}
	fmt.Printf("Test 4 - Expected: 5, Got: %d\n", countDistinctWords(test4))
}

/*output:
Test 1 - Expected: 5, Got: 5
Test 2 - Expected: 2, Got: 2
Test 3 - Expected: 0, Got: 0
Test 4 - Expected: 5, Got: 5
*/
