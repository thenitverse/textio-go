package main

import "fmt"

func reformat(message string, formatter func(string) string) string {
	first := formatter(message)
	second := formatter(first)
	third := formatter(second)
	result := "TEXTIO: " + third
	return result
}
func main() {
	// Test Case 1: Using a named function that adds a star
	addStar := func(s string) string {
		return s + "*"
	}
	fmt.Println(reformat("Magic", addStar))
	// Expected: TEXTIO: Magic***

	// Test Case 2: Using an anonymous function to make text uppercase
	// Note: In a real scenario, you'd import "strings" and use strings.ToUpper
	shout := func(s string) string {
		return s + "!"
	}
	fmt.Println(reformat("hello", shout))
	// Expected: TEXTIO: hello!!!

	// Test Case 3: A formatter that doubles the string
	double := func(s string) string {
		return s + s
	}
	fmt.Println(reformat("Go", double))
	// Expected: TEXTIO: GoGoGoGoGoGoGoGo
}

/*output:
TEXTIO: Magic***
TEXTIO: hello!!!
TEXTIO: GoGoGoGoGoGoGoGo
*/
