package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	length := [3]int{len(primary), len(primary) + len(secondary), len(primary) + len(secondary) + len(tertiary)}
	return messages, length
}

func main() {
	// Example test case
	m1 := "Hello"
	m2 := "World"
	m3 := "!"

	msgs, costs := getMessageWithRetries(m1, m2, m3)

	fmt.Println("Messages:", msgs)
	// Expected: [Hello World !]

	fmt.Println("Costs:", costs)
	// Expected: [5 10 11]
}

/*output:
Messages: [Hello World !]
Costs: [5 10 11]
*/
