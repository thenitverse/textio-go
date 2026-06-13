package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	user, ok := users[name]
	if !ok {
		return false, errors.New("not found")

	}
	if user.scheduledForDeletion == false {
		return false, nil
	}
	if user.scheduledForDeletion == true {
		delete(users, name)
	}
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	// 1. Setup a map of users
	testUsers := map[string]user{
		"Arthur":   {name: "Arthur", scheduledForDeletion: true},
		"Trillian": {name: "Trillian", scheduledForDeletion: false},
	}

	// Case 1: User does not exist
	deleted, err := deleteIfNecessary(testUsers, "Ford")
	fmt.Printf("Non-existent user: deleted=%v, err=%v\n", deleted, err)

	// Case 2: User exists but not scheduled for deletion
	deleted, err = deleteIfNecessary(testUsers, "Trillian")
	fmt.Printf("Keep user: deleted=%v, err=%v, count=%d\n", deleted, err, len(testUsers))

	// Case 3: User exists and IS scheduled for deletion
	deleted, err = deleteIfNecessary(testUsers, "Arthur")
	fmt.Printf("Delete user: deleted=%v, err=%v, count=%d\n", deleted, err, len(testUsers))
}

/*output:
Non-existent user: deleted=false, err=not found
Keep user: deleted=false, err=<nil>, count=2
Delete user: deleted=true, err=<nil>, count=1
*/
