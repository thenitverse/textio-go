package main

import "fmt"

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	msg := []Message{}
	for _, message := range messages {
		if message.Type() == filterType {
			msg = append(msg, message)
		}

	}
	return msg

}
func main() {
	messages := []Message{
		TextMessage{Sender: "Boots", Content: "Hello!"},
		MediaMessage{Sender: "Wizard", MediaType: "video", Content: "Fireball Tutorial"},
		TextMessage{Sender: "Apprentice", Content: "I'm learning Go!"},
	}

	// Test Case 1: Filter for "text"
	// Should return 2 messages
	textMsgs := filterMessages(messages, "text")

	// Test Case 2: Filter for "media"
	// Should return 1 message
	mediaMsgs := filterMessages(messages, "media")

	// Test Case 3: Filter for "link"
	// Should return 0 messages (empty slice)
	linkMsgs := filterMessages(messages, "link")
	fmt.Printf("Text messages found: %d\n", len(textMsgs))
	fmt.Printf("Media messages found: %d\n", len(mediaMsgs))
	fmt.Printf("Link messages found: %d\n", len(linkMsgs))
}

/*output:
Text messages found: 2
Media messages found: 1
Link messages found: 0
*/
