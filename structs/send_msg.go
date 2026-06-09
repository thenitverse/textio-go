package main

import "fmt"

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}
	return "", false
}

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

func main() {
	tests := []struct {
		name           string
		membershipType string
		message        string
		expectMsg      string
		expectOK       bool
	}{
		{"Alice", "standard", "hi", "hi", true},
		{"Bob", "premium", "hello world", "hello world", true},
		{"Cara", "standard", "", "", true},
		{"Dane", "standard", string(make([]byte, 100)), string(make([]byte, 100)), true},
		{"Eli", "standard", string(make([]byte, 101)), "", false},
		{"Faye", "premium", string(make([]byte, 1000)), string(make([]byte, 1000)), true},
		{"Gus", "premium", string(make([]byte, 1001)), "", false},
	}

	for _, tc := range tests {
		user := newUser(tc.name, tc.membershipType)
		gotMsg, gotOK := user.SendMessage(tc.message, len(tc.message))

		fmt.Println("----")
		fmt.Println("user:", tc.name)
		fmt.Println("membership:", tc.membershipType)
		fmt.Println("expected:", tc.expectMsg, tc.expectOK)
		fmt.Println("got:", gotMsg, gotOK)
	}
}

/*output:
----
user: Alice
membership: standard
expected: hi true
got: hi true
----
user: Bob
membership: premium
expected: hello world true
got: hello world true
----
user: Cara
membership: standard
expected:  true
got:  true
----
user: Dane
membership: standard
expected:  true
got:  true
----
user: Eli
membership: standard
expected:  false
got:  false
----
user: Faye
membership: premium
expected:  true
got:  true
----
user: Gus
membership: premium
expected:  false
got:  false
*/
