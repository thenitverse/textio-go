package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func main() {
	// Create our first adder
	myAdder := adder()
	fmt.Println(myAdder(1))  // Expected: 1
	fmt.Println(myAdder(10)) // Expected: 11
	fmt.Println(myAdder(5))  // Expected: 16

	// Create a second, independent adder
	secondAdder := adder()
	fmt.Println(secondAdder(100)) // Expected: 100 (Starts at 0)

	// The first one still remembers its own sum
	fmt.Println(myAdder(1)) // Expected: 17
}

/*output;
1
11
16
100
17
*/
