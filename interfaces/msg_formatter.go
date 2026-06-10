package main

import "fmt"

type formatter interface {
	format() string
}
type plainText struct {
	message string
}
type bold struct {
	message string
}
type code struct {
	message string
}

func (pt plainText) format() string {
	return pt.message
}
func (b bold) format() string {
	return "**" + b.message + "**"
}
func (c code) format() string {
	return "`" + c.message + "`"
}

func sendMessage(format formatter) string {
	return format.format() // Adjusted to call Format without an argument
}

func main() {
	p := plainText{message: "Hello world"}
	b := bold{message: "Welcome to the forest"}
	c := code{message: "fmt.Println(message)"}

	fmt.Println(sendMessage(p))
	fmt.Println(sendMessage(b))
	fmt.Println(sendMessage(c))
}

/*output:
Hello world
**Welcome to the forest**
`fmt.Println(message)`
*/
