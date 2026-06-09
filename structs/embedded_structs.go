package main

import "fmt"

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

func main() {
	// You can initialize the embedded struct by
	// nesting the user literal inside the sender literal
	s := sender{
		rateLimit: 100,
		user: user{
			name:   "Gandalf",
			number: 5551234,
		},
	}

	// Because 'user' is embedded, you can access its fields
	// directly on the 'sender' instance!
	fmt.Println("Sender Name:", s.name)
	fmt.Println("Sender Number:", s.number)
	fmt.Println("Rate Limit:", s.rateLimit)
}

/*output:
Sender Name: Gandalf
Sender Number: 5551234
Rate Limit: 100
*/
