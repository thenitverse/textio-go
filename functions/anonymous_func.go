package main

import "fmt"

func printReports(intro, body, outro string) {
	printCostReport(func(a string) int {
		return 2 * len(a)
	}, intro)
	printCostReport(func(b string) int {
		return 3 * len(b)
	}, body)
	printCostReport(func(c string) int {
		return 4 * len(c)
	}, outro)
}

func main() {
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

/*output:
Message: "Welcome to the Hotel California" Cost: 62 cents
Message: "Such a lovely place" Cost: 57 cents
Message: "Plenty of room at the Hotel California" Cost: 152 cents
*/
