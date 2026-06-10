package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	em, ok := e.(email)
	if ok {
		return em.toAddress, em.cost()

	}
	sm, ok := e.(sms)
	if ok {
		return sm.toPhoneNumber, sm.cost()
	}
	return "", 0.0

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

func (em email) cost() float64 {
	if !em.isSubscribed {
		return float64(len(em.body)) * .05
	}
	return float64(len(em.body)) * .01
}

func (sm sms) cost() float64 {
	if !sm.isSubscribed {
		return float64(len(sm.body)) * .1
	}
	return float64(len(sm.body)) * .03
}

func (inv invalid) cost() float64 {
	return 0.0
}

func main() {
	// Test Case 1: An Email
	e1 := email{
		isSubscribed: true,
		body:         "Hello there!",
		toAddress:    "obiwan@tatooine.com",
	}
	addr1, cost1 := getExpenseReport(e1)
	fmt.Printf("Email -> To: %s, Cost: %.2f\n", addr1, cost1)

	// Test Case 2: An SMS
	s1 := sms{
		isSubscribed:  false,
		body:          "General Kenobi!",
		toPhoneNumber: "+1555000123",
	}
	addr2, cost2 := getExpenseReport(s1)
	fmt.Printf("SMS   -> To: %s, Cost: %.2f\n", addr2, cost2)

	// Test Case 3: An Invalid type
	inv := invalid{}
	addr3, cost3 := getExpenseReport(inv)
	fmt.Printf("Other -> To: '%s', Cost: %.2f\n", addr3, cost3)
}

/*output:
Email -> To: obiwan@tatooine.com, Cost: 0.12
SMS   -> To: +1555000123, Cost: 1.50
Other -> To: '', Cost: 0.00
*/
