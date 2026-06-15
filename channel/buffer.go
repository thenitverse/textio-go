package main

import "fmt"

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for i := 0; i < len(emails); i++ {
		ch <- emails[i]
	}
	return ch
}

func main() {
	testCases := [][]string{
		{
			"To boldly go where no man has gone before.",
			"Live long and prosper.",
		},
		{
			"The needs of the many outweigh the needs of the few, or the one.",
			"Change is the essential process of all existence.",
			"Resistance is futile.",
		},
		{
			"It's life, Jim, but not as we know it.",
			"Infinite diversity in infinite combinations.",
			"Make it so.",
			"Engage!",
		},
		{},
	}

	for i, emails := range testCases {
		ch := addEmailsToQueue(emails)
		fmt.Printf("Test %d: expected %d, got %d\n", i+1, len(emails), len(ch))

		for len(ch) > 0 {
			fmt.Println(" ", <-ch)
		}
		fmt.Println("---")
	}
}

/*output:
Test 1: expected 2, got 2
  To boldly go where no man has gone before.
  Live long and prosper.
---
Test 2: expected 3, got 3
  The needs of the many outweigh the needs of the few, or the one.
  Change is the essential process of all existence.
  Resistance is futile.
---
Test 3: expected 4, got 4
  It's life, Jim, but not as we know it.
  Infinite diversity in infinite combinations.
  Make it so.
  Engage!
---
Test 4: expected 0, got 0
---*/
