package main

import "fmt"

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func main() {
	// Test Case 1: Standard message update
	testEmail := email{
		message:     "Old Message",
		fromAddress: "wizard@boot.dev",
		toAddress:   "apprentice@boot.dev",
	}

	fmt.Printf("Before: %s\n", testEmail.message)
	testEmail.setMessage("New Message")
	fmt.Printf("After:  %s\n", testEmail.message)

	// Test Case 2: Emptying a message
	testEmail.setMessage("")
	fmt.Printf("Empty:  '%s'\n", testEmail.message)
}

/*output:
Before: Old Message
After:  New Message
Empty:  ''
*/
