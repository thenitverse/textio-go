package main

import "fmt"

func main() {
	const name = "Saul Goodman"
	const openRate = 30.54

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f percent\n", name, openRate) // 0.1f for  rounded tenths place

	fmt.Print(msg)
}

// output: Hi Saul Goodman, your open rate is 30.5 percent
