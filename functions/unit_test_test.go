package main

import "testing"

func TestGetMonthlyPrice(t *testing.T) {
	// Test the basic tier
	res := getMonthlyPrice("basic")
	expected := 10000
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}

	// Test the enterprise tier
	res = getMonthlyPrice("enterprise")
	expected = 50000
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}

	// Test an invalid tier
	res = getMonthlyPrice("invalid")
	expected = 0
	if res != expected {
		t.Errorf("Expected %d, but got %d", expected, res)
	}
}

/*output:
=== RUN   TestGetMonthlyPrice
--- PASS: TestGetMonthlyPrice (0.00s)
PASS
ok      functions       0.793s
*/
