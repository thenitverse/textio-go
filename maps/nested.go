package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	n := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		first := runes[0]
		if _, ok := n[first]; !ok {
			n[first] = make(map[string]int)
		}
		n[first][name]++
	}
	return n
}

func main() {
	// Test Case 1: Standard names
	names1 := []string{"billy", "billy", "bob", "joe"}
	result1 := getNameCounts(names1)
	fmt.Println("Test Case 1:")
	for char, nestedMap := range result1 {
		fmt.Printf("%c: %v\n", char, nestedMap)
	}

	fmt.Println()

	// Test Case 2: Names starting with special/unicode characters (runes)
	names2 := []string{"😊", "😊", "bear", "boots", "🐟"}
	result2 := getNameCounts(names2)
	fmt.Println("Test Case 2:")
	for char, nestedMap := range result2 {
		fmt.Printf("%c: %v\n", char, nestedMap)
	}
}

/*output:
Test Case 1:
b: map[billy:2 bob:1]
j: map[joe:1]

Test Case 2:
😊: map[😊:2]
b: map[bear:1 boots:1]
🐟: map[🐟:1]
*/
