package main

import "fmt"

func bulkSend(numMessages int) float64 {
	total := 0.0
	for i := 0; i < numMessages; i++ {
		cost := 1.0 + float64(i)*0.01
		total += cost

	}
	return total
}

func main() {
	// 0 messages -> 0.00
	fmt.Printf("0 messages: %.2f (expected 0.00)\n", bulkSend(0))

	// 1 message: 1.0 + 0.00 = 1.00
	fmt.Printf("1 message:  %.2f (expected 1.00)\n", bulkSend(1))

	// 5 messages: 1.00 + 1.01 + 1.02 + 1.03 + 1.04 = 5.10
	fmt.Printf("5 messages: %.2f (expected 5.10)\n", bulkSend(5))

	// 10 messages: sum from i=0..9 -> 10.45
	fmt.Printf("10 messages: %.2f (expected 10.45)\n", bulkSend(10))

	// 20 messages -> 21.90
	fmt.Printf("20 messages: %.2f (expected 21.90)\n", bulkSend(20))

	// 30 messages -> 34.35
	fmt.Printf("30 messages: %.2f (expected 34.35)\n", bulkSend(30))
}

/*output:
0 messages: 0.00 (expected 0.00)
1 message:  1.00 (expected 1.00)
5 messages: 5.10 (expected 5.10)
10 messages: 10.45 (expected 10.45)
20 messages: 21.90 (expected 21.90)
30 messages: 34.35 (expected 34.35)
*/
