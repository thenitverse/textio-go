package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		tier     string
		expected string
	}
	runCases := []testCase{
		{"basic", "You get 1,000 texts per month for $30 per month."},
		{"premium", "You get 50,000 texts per month for $60 per month."},
	}
	submitCases := append(runCases, []testCase{
		{"enterprise", "You get unlimited texts per month for $100 per month."},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := getProductMessage(test.tier)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, test.tier, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.tier, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true

/*output:
=== RUN   Test
---------------------------------
Inputs:     (basic)
Expecting:  You get 1,000 texts per month for $30 per month.
Actual:     You get 1,000 texts per month for $30 per month.
Pass
---------------------------------
Inputs:     (premium)
Expecting:  You get 50,000 texts per month for $60 per month.
Actual:     You get 50,000 texts per month for $60 per month.
Pass
---------------------------------
Inputs:     (enterprise)
Expecting:  You get unlimited texts per month for $100 per month.
Actual:     You get unlimited texts per month for $100 per month.
Pass
---------------------------------
3 passed, 0 failed
--- PASS: Test (0.00s)
=== RUN   TestGetMonthlyPrice
--- PASS: TestGetMonthlyPrice (0.00s)
PASS
ok      functions       0.752s
*/
