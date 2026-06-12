package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, msg := range messages {
		messages[i].tags = tagger(msg)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	lowerContent := strings.ToLower(msg.content)
	if strings.Contains(lowerContent, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(lowerContent, "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}

func main() {
	testMessages := []sms{
		{
			id:      "001",
			content: "URGENT: Your account has been compromised!",
		},
		{
			id:      "002",
			content: "Don't miss our summer sale! Huge discounts!",
		},
		{
			id:      "003",
			content: "Get your urgent sale items before they are gone!",
		},
		{
			id:      "004",
			content: "Just a friendly message from a friend.",
		},
	}

	tagged := tagMessages(testMessages, tagger)

	for _, msg := range tagged {
		fmt.Printf("ID: %s | Content: %q | Tags: %v\n", msg.id, msg.content, msg.tags)
	}
}

/*output:
ID: 001 | Content: "URGENT: Your account has been compromised!" | Tags: [Urgent]
ID: 002 | Content: "Don't miss our summer sale! Huge discounts!" | Tags: [Promo]
ID: 003 | Content: "Get your urgent sale items before they are gone!" | Tags: [Urgent Promo]
ID: 004 | Content: "Just a friendly message from a friend." | Tags: []
➜  slices git:(main) ✗ */
