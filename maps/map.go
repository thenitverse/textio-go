package main

import (
	"errors"
	"fmt" // Imported for printing results
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	n := make(map[string]user)
	for i, name := range names {
		n[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}
	return n, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	// Test Case 1: Matching slice sizes (Success case)
	names1 := []string{"Eren", "Armin", "Mikasa"}
	phones1 := []int{14355550987, 98765550987, 18265554567}

	fmt.Println("--- Test Case 1 ---")
	result1, err1 := getUserMap(names1, phones1)
	if err1 != nil {
		fmt.Printf("Unexpected error: %v\n", err1)
	} else {
		fmt.Printf("Returned map: %+v\n", result1)
	}

	// Test Case 2: Mismatched slice sizes (Error case)
	names2 := []string{"Eren", "Armin"}
	phones2 := []int{14355550987, 98765550987, 18265554567}

	fmt.Println("\n--- Test Case 2 ---")
	result2, err2 := getUserMap(names2, phones2)
	if err2 != nil {
		fmt.Printf("Expected error received: %v\n", err2)
	} else {
		fmt.Printf("Unexpected success! Map: %+v\n", result2)
	}
}

/*output:
--- Test Case 1 ---
Returned map: map[Armin:{name:Armin phoneNumber:98765550987} Eren:{name:Eren phoneNumber:14355550987} Mikasa:{name:Mikasa phoneNumber:18265554567}]

--- Test Case 2 ---
Expected error received: invalid sizes
➜  maps git:(main)
*/
