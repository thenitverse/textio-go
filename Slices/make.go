package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	myMsg := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		cost := float64(len(messages[i])) * 0.01
		myMsg[i] = cost

	}
	return myMsg
}
func main() {
	messages := []string{"Hello", "World"}
	costs := getMessageCosts(messages)
	fmt.Println(costs) // Expected: [0.05 0.05]
}
