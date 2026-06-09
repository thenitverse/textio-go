package main

import (
	"testing"
)

type User struct {
	Membership
	Name string
}
type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	myUser := User{
		Name: name,
	}
	myUser.Type = membershipType
	if membershipType == "premium" {
		myUser.MessageCharLimit = 1000
	} else {
		myUser.MessageCharLimit = 100
	}
	return myUser
}

// main_test.go

func TestNewUser(t *testing.T) {
	tests := []struct {
		name           string
		membershipType string
		wantLimit      int
	}{
		{
			name:           "Standard User",
			membershipType: "standard",
			wantLimit:      100,
		},
		{
			name:           "Premium User",
			membershipType: "premium",
			wantLimit:      1000,
		},
		{
			name:           "Empty Membership",
			membershipType: "",
			wantLimit:      100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newUser(tt.name, tt.membershipType)

			if got.Name != tt.name {
				t.Errorf("newUser().Name = %v, want %v", got.Name, tt.name)
			}

			if got.Type != tt.membershipType {
				t.Errorf("newUser().Type = %v, want %v", got.Type, tt.membershipType)
			}

			if got.MessageCharLimit != tt.wantLimit {
				t.Errorf("newUser().MessageCharLimit = %v, want %v", got.MessageCharLimit, tt.wantLimit)
			}
		})
	}
}
