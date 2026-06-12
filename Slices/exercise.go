package main

func isValidPassword(password string) bool {
	length := len(password)
	hasUpper := false
	hasNumber := false
	for _, ch := range password {
		if ch >= 'A' && ch <= 'Z' {
			hasUpper = true

		}
		if ch >= '0' && ch <= '9' {
			hasNumber = true
		}

	}
	return length >= 5 && length <= 12 && hasNumber && hasUpper

}

func main() {
	// Valid: 7 chars, has uppercase 'P', has digit '1'
	println("Pass123:", isValidPassword("Pass123")) // Expected: true

	// Invalid: Too short (3 chars)
	println("pas:", isValidPassword("pas")) // Expected: false

	// Invalid: No digits
	println("Password:", isValidPassword("Password")) // Expected: false

	// Invalid: No uppercase letters
	println("123456:", isValidPassword("123456")) // Expected: false

	// Invalid: Too long (17 chars)
	println("VeryLongPassword1:", isValidPassword("VeryLongPassword1")) // Expected: false

	// Valid: Exactly 5 chars, has 'S' and '5'
	println("Short5:", isValidPassword("Short5")) // Expected: true
}

/*output:
Pass123: true
pas: false
Password: false
123456: false
VeryLongPassword1: false
Short5: true
*/
