package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.recipient.number == 0 {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.name == "" {
		return false
	}
	return true
}

func main() {
	// Test Case 1: Valid message with all fields populated
	msg1 := messageToSend{
		message: "Hello, friend!",
		sender: user{
			name:   "Alice",
			number: 123456789,
		},
		recipient: user{
			name:   "Bob",
			number: 987654321,
		},
	}
	fmt.Printf("Test 1 (Should be true): %v\n", canSendMessage(msg1))

	// Test Case 2: Sender is missing a name (zero value "")
	msg2 := messageToSend{
		message: "Are you there?",
		sender: user{
			number: 123456789,
		},
		recipient: user{
			name:   "Bob",
			number: 987654321,
		},
	}
	fmt.Printf("Test 2 (Should be false): %v\n", canSendMessage(msg2))

	// Test Case 3: Recipient is missing a phone number (zero value 0)
	msg3 := messageToSend{
		message: "Hey!",
		sender: user{
			name:   "Alice",
			number: 123456789,
		},
		recipient: user{
			name: "Bob",
		},
	}
	fmt.Printf("Test 3 (Should be false): %v\n", canSendMessage(msg3))
}

/*output:
Test 1 (Should be true): true
Test 2 (Should be false): false
Test 3 (Should be false): false
*/
