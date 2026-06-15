package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func test(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func main() {
	test("Hello there Kaladin!")
	test("Hi there Shallan!")
	test("Hey there Dalinar!")
}

/*output:
Email sent: 'Hello there Kaladin!'
Email received: 'Hello there Kaladin!'
========================
Email sent: 'Hi there Shallan!'
Email received: 'Hi there Shallan!'
========================
Email sent: 'Hey there Dalinar!'
Email received: 'Hey there Dalinar!'
========================
*/
