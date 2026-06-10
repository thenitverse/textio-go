package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}
	return d.priorityLevel

}
func (g groupMessage) importance() int {
	return g.priorityLevel
}
func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, n.importance()
	case groupMessage:
		return v.groupName, n.importance()
	case systemAlert:
		return v.alertCode, n.importance()
	default:
		return "", 0
	}

}

func main() {
	// Test Case 1: Urgent Direct Message
	dm := directMessage{
		senderUsername: "Alice",
		isUrgent:       true,
		priorityLevel:  10,
	}
	name, score := processNotification(dm)
	fmt.Printf("DM - Expected: Alice/50, Got: %s/%d\n", name, score)

	// Test Case 2: Group Message
	gm := groupMessage{
		groupName:     "Gophers",
		priorityLevel: 25,
	}
	name, score = processNotification(gm)
	fmt.Printf("Group - Expected: Gophers/25, Got: %s/%d\n", name, score)

	// Test Case 3: System Alert
	sa := systemAlert{
		alertCode: "ERR_404",
	}
	name, score = processNotification(sa)
	fmt.Printf("Alert - Expected: ERR_404/100, Got: %s/%d\n", name, score)
}

/*output:
DM - Expected: Alice/50, Got: Alice/50
Group - Expected: Gophers/25, Got: Gophers/25
Alert - Expected: ERR_404/100, Got: ERR_404/100
*/
