package main

import (
	"fmt"
	"testing"
	"time"
)

func sendMessage(msg message) (string, int) {
	result := msg.getMessage()
	return result, 3 * len(result)
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func TestSendMessage(t *testing.T) {
	tests := []struct {
		name         string
		msg          message
		expectedText string
		expectedCost int
	}{
		{
			name: "Birthday Message Test",
			msg: birthdayMessage{
				birthdayTime:  time.Date(1994, time.March, 21, 0, 0, 0, 0, time.UTC),
				recipientName: "John Doe",
			},
			expectedText: "Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z",
			expectedCost: 168, // 56 characters * 3
		},
		{
			name: "Sending Report Test",
			msg: sendingReport{
				reportName:    "First Report",
				numberOfSends: 10,
			},
			expectedText: `Your "First Report" report is ready. You've sent 10 messages.`,
			expectedCost: 183, // 61 characters * 3
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			text, cost := sendMessage(tc.msg)
			if text != tc.expectedText {
				t.Errorf("expected text %q, got %q", tc.expectedText, text)
			}
			if cost != tc.expectedCost {
				t.Errorf("expected cost %d, got %d", tc.expectedCost, cost)
			}
		})
	}
}
