package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	e1 := email{
		isSubscribed: true,
		body:         "Hello!",
		toAddress:    "wizard@boot.dev",
	}

	e2 := sms{
		isSubscribed:  false,
		body:          "Your potion is ready.",
		toPhoneNumber: "555-0123",
	}

	e3 := invalid{}

	// Testing email
	addr1, cost1 := getExpenseReport(e1)
	fmt.Printf("Email -> To: %s, Cost: %f\n", addr1, cost1)

	// Testing sms
	addr2, cost2 := getExpenseReport(e2)
	fmt.Printf("SMS -> To: %s, Cost: %f\n", addr2, cost2)

	// Testing invalid/default
	addr3, cost3 := getExpenseReport(e3)
	fmt.Printf("Default -> To: '%s', Cost: %f\n", addr3, cost3)
}

/*output:
Email -> To: wizard@boot.dev, Cost: 0.060000
SMS -> To: 555-0123, Cost: 2.100000
Default -> To: '', Cost: 0.000000
*/
