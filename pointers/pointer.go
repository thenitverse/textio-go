package main

import "fmt"

type Message struct {
	Recipient string
	Text      string
}

func getMessageText(m Message) string {
	return fmt.Sprintf(`
To: %v
Message: %v
`, m.Recipient, m.Text)
}
func main() {
	// Test Case 1: Standard message
	msg1 := Message{
		Recipient: "Albus",
		Text:      "The lemon drops are in the tower.",
	}
	fmt.Println("Test 1:")
	fmt.Print(getMessageText(msg1))

	// Test Case 2: Empty strings
	msg2 := Message{
		Recipient: "",
		Text:      "",
	}
	fmt.Println("\nTest 2 (Empty):")
	fmt.Print(getMessageText(msg2))

	// Test Case 3: Long message
	msg3 := Message{
		Recipient: "The Council",
		Text:      "We must discuss the migration of the salmon before the frost sets in.",
	}
	fmt.Println("\nTest 3 (Long):")
	fmt.Print(getMessageText(msg3))
}

/*output:
Test 1:

To: Albus
Message: The lemon drops are in the tower.

Test 2 (Empty):

To:
Message:

Test 3 (Long):

To: The Council
Message: We must discuss the migration of the salmon before the frost sets in.
➜  pointers git:(main) ✗ */
