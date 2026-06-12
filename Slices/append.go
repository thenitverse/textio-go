package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	d := []float64{}
	for _, cost := range costs {
		if cost.day == day {
			d = append(d, cost.value)
		}
	}
	return d
}

func main() {
	// Sample data to test with
	data := []cost{
		{day: 1, value: 12.50},
		{day: 2, value: 5.00},
		{day: 1, value: 3.00},
		{day: 3, value: 25.00},
	}

	// Test Case 1: Multiple costs on day 1
	res1 := getDayCosts(data, 1)
	fmt.Printf("Day 1 - Expected: [12.5, 3], Got: %v\n", res1)

	// Test Case 2: One cost on day 2
	res2 := getDayCosts(data, 2)
	fmt.Printf("Day 2 - Expected: [5], Got: %v\n", res2)

	// Test Case 3: No costs on day 4
	res3 := getDayCosts(data, 4)
	fmt.Printf("Day 4 - Expected: [], Got: %v\n", res3)
}

/*output:
Day 1 - Expected: [12.5, 3], Got: [12.5 3]
Day 2 - Expected: [5], Got: [5]
Day 4 - Expected: [], Got: []
*/
