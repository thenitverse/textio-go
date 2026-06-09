package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {
	// Create a new instance of the struct
	msg := messageToSend{
		phoneNumber: 15555551234,
		message:     "Hello from the realm of Go!",
	}

	// Print the fields to verify they are set correctly
	fmt.Println("Phone Number:", msg.phoneNumber)
	fmt.Println("Message:", msg.message)
}

/*output:
Phone Number: 15555551234
Message: Hello from the realm of Go!
*/
