package main

import "fmt"

func main() {
	// Case 1: Bad word is in the middle
	msg1 := []string{"the", "quick", "brown", "crap", "fox"}
	bad1 := []string{"crap", "shoot"}
	fmt.Println("Test 1 (Expected 3):", indexOfFirstBadWord(msg1, bad1))

	// Case 2: No bad words at all
	msg2 := []string{"hello", "world"}
	bad2 := []string{"dang", "frick"}
	fmt.Println("Test 2 (Expected -1):", indexOfFirstBadWord(msg2, bad2))

	// Case 3: Bad word is the very first word
	msg3 := []string{"shoot", "I", "forgot", "my", "keys"}
	bad3 := []string{"crap", "shoot"}
	fmt.Println("Test 3 (Expected 0):", indexOfFirstBadWord(msg3, bad3))

	// Case 4: Multiple bad words (should return the index of the FIRST one)
	msg4 := []string{"this", "dang", "code", "is", "crap"}
	bad4 := []string{"dang", "crap"}
	fmt.Println("Test 4 (Expected 1):", indexOfFirstBadWord(msg4, bad4))
}

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, m := range msg {
		for _, badWord := range badWords {
			if m == badWord {
				return i
			}
		}
	}
	return -1
}

/*output:
Test 1 (Expected 3): 3
Test 2 (Expected -1): -1
Test 3 (Expected 0): 0
Test 4 (Expected 1): 1
*/
