package main

import (
	"strings"
)

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message

	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}
func main() {
	// Test Case 1: A normal string with profanity
	msg1 := "That fubb is a total witch!"
	removeProfanity(&msg1)
	println("Result 1:", msg1)

	// Test Case 2: A string with no profanity
	msg2 := "Have a wonderful day."
	removeProfanity(&msg2)
	println("Result 2:", msg2)

	// Test Case 3: Passing nil
	// This tests your safety check!
	var msg3 *string = nil
	removeProfanity(msg3)
	println("Result 3: (Successfully handled nil without crashing)")
}

/*output:
Result 1: That **** is a total *****!
Result 2: Have a wonderful day.
Result 3: (Successfully handled nil without crashing)
*/
