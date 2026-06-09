package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)

}

func main() {
	// Test Case 1: Standard credentials
	user1 := authenticationInfo{
		username: "wagslane",
		password: "password123",
	}
	fmt.Println("Test 1:", user1.getBasicAuth())

	// Test Case 2: Empty password
	user2 := authenticationInfo{
		username: "admin",
		password: "",
	}
	fmt.Println("Test 2:", user2.getBasicAuth())

	// Test Case 3: Numerical strings
	user3 := authenticationInfo{
		username: "123",
		password: "456",
	}
	fmt.Println("Test 3:", user3.getBasicAuth())
}

/*output:
Test 1: Authorization: Basic wagslane:password123
Test 2: Authorization: Basic admin:
Test 3: Authorization: Basic 123:456
*/
