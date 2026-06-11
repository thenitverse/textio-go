package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that costs $%.2f to be sent to '%v' cannot be sent", cost, recipient)
}

func main() {
	output := getSMSErrorString(1.5, "+15555555")
	fmt.Println(output)
}

// SMS that costs $1.50 to be sent to '+15555555' cannot be sent
