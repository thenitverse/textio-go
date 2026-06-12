package main

import "fmt"

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func analyzeMessage(analytics *Analytics, msg Message) {
	if msg.Success {
		(*analytics).MessagesSucceeded++
	} else {
		(*analytics).MessagesFailed++
	}
	(*analytics).MessagesTotal++
}
func main() {
	// 1. Initialize our Analytics struct
	stats := Analytics{
		MessagesTotal:     0,
		MessagesFailed:    0,
		MessagesSucceeded: 0,
	}

	// 2. Create some dummy messages
	msg1 := Message{Recipient: "Hunter", Success: true}
	msg2 := Message{Recipient: "Gatherer", Success: false}

	// 3. Call the function.
	// Note: We use '&' to pass the address (pointer) of stats
	analyzeMessage(&stats, msg1)
	analyzeMessage(&stats, msg2)

	// 4. Print results to verify the original 'stats' was updated
	// You'll need to add this import at the top of your file
	fmt.Printf("Total: %d\n", stats.MessagesTotal)
	fmt.Printf("Succeeded: %d\n", stats.MessagesSucceeded)
	fmt.Printf("Failed: %d\n", stats.MessagesFailed)
}

/*output:
Total: 2
Succeeded: 1
Failed: 1
*/
