package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	*message = strings.ReplaceAll(*message, "fubb", "****")
	*message = strings.ReplaceAll(*message, "shiz", "****")
	*message = strings.ReplaceAll(*message, "witch", "*****")

}
func main() {
	// Test 1: contains "fubb"
	msg1 := "English, motherfubber, do you speak it?"
	removeProfanity(&msg1)
	fmt.Println(msg1)
	// Expected: "English, mother****er, do you speak it?"

	// Test 2: contains "shiz"
	msg2 := "Oh man I've seen some crazy ass shiz in my time..."
	removeProfanity(&msg2)
	fmt.Println(msg2)
	// Expected: "Oh man I've seen some crazy ass **** in my time..."

	// Test 3: contains "witch"
	msg3 := "Does he look like a witch?"
	removeProfanity(&msg3)
	fmt.Println(msg3)
	// Expected: "Does he look like a *****?"

	// Test 4: no profanity — should remain unchanged
	msg4 := "Hello, how are you?"
	removeProfanity(&msg4)
	fmt.Println(msg4)
	// Expected: "Hello, how are you?"

	// Test 5: multiple profanities in one string
	msg5 := "fubb and shiz and witch"
	removeProfanity(&msg5)
	fmt.Println(msg5)
	// Expected: "**** and **** and *****"
}

/*output:
English, mother****er, do you speak it?
Oh man I've seen some crazy ass **** in my time...
Does he look like a *****?
Hello, how are you?
**** and **** and ******/
