package main

import "fmt"

func maxMessages(thresh int) int {
	total := 0
	for i := 0; ; i++ {
		cost := 100 + i
		total += cost
		if total > thresh {
			return i
		}
	}
	return total
}

func main() {
	testCases := []struct {
		thresh   int
		expected int
	}{
		{103, 1},
		{205, 2},
		{1000, 9},
		{100, 1},
		{3000, 26},
		{4000, 34},
		{5000, 41},
		{0, 0},
	}

	for _, tc := range testCases {
		result := maxMessages(tc.thresh)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("thresh=%d | expected=%d | got=%d | %s\n", tc.thresh, tc.expected, result, status)
	}
}

/*output:
thresh=100 | expected=1 | got=1 | PASS
thresh=3000 | expected=26 | got=26 | PASS
thresh=4000 | expected=34 | got=34 | PASS
thresh=5000 | expected=41 | got=41 | PASS
thresh=0 | expected=0 | got=0 | PASS
*/
