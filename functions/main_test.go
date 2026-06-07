package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		costPerSend  int
		numLastMonth int
		numThisMonth int
		expected     int
	}

	runCases := []testCase{
		{2, 89, 102, 26},
		{2, 98, 104, 12},
	}

	submitCases := append(runCases, []testCase{
		{3, 50, 40, -30},
		{3, 60, 60, 0},
	}...)

	testCases := runCases
	if withSubmit {
		testCases = submitCases
	}
	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for _, test := range testCases {
		output := monthlyBillIncrease(test.costPerSend, test.numLastMonth, test.numThisMonth)
		_ = getBillForMonth(0, 0)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Fail
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.costPerSend, test.numLastMonth, test.numThisMonth, test.expected, output)
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
Inputs:     (2, 89, 102)
Expecting:  26
Actual:     26
Pass
---------------------------------
Inputs:     (2, 98, 104)
Expecting:  12
Actual:     12
Pass
---------------------------------
Inputs:     (3, 50, 40)
Expecting:  -30
Actual:     -30
Pass
---------------------------------
Inputs:     (3, 60, 60)
Expecting:  0
Actual:     0
Pass
---------------------------------
4 passed, 0 failed
--- PASS: Test (0.00s)
=== RUN   TestGetMonthlyPrice
--- PASS: TestGetMonthlyPrice (0.00s)
PASS
ok      functions       0.854s
*/
