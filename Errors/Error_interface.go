package main

import (
	"fmt"
)

func main() {
	// Test case 1: Successful messages
	cost, err := sendSMSToCouple("Hello customer", "Hello spouse")
	if err != nil {
		fmt.Printf("Error 1: %v\n", err)
	} else {
		fmt.Printf("Success 1! Total cost: %v\n", cost)
	}

	// Test case 2: Message too long
	cost2, err2 := sendSMSToCouple("This message is way too long for our system to handle", "Short msg")
	if err2 != nil {
		fmt.Printf("Error 2: %v\n", err2)
	} else {
		fmt.Printf("Success 2! Total cost: %v\n", cost2)
	}
}

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	msg, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}
	m, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}
	return msg + m, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}

/*output:
Success 1! Total cost: 52
Error 2: can't send texts over 25 characters
*/
