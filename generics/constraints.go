package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	cost := newItem.GetCost()
	if balance < cost {
		return nil, 0.0, errors.New("insufficient funds")

	} else {
		oldItems = append(oldItems, newItem)
		newBalance := balance - cost
		return oldItems, newBalance, nil
	}

}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}
func main() {
	sub := subscription{
		userEmail: "a@site.com",
		startDate: time.Now(),
		interval:  "monthly",
	}

	yearly := subscription{
		userEmail: "b@site.com",
		startDate: time.Now(),
		interval:  "yearly",
	}

	usage := oneTimeUsagePlan{
		userEmail:        "c@site.com",
		numEmailsAllowed: 100,
	}

	// test 1: enough balance
	items1, balance1, err1 := chargeForLineItem(sub, []subscription{}, 100)
	fmt.Println("test1:", items1, balance1, err1)

	// test 2: not enough balance
	items2, balance2, err2 := chargeForLineItem(yearly, []subscription{}, 100)
	fmt.Println("test2:", items2, balance2, err2)

	// test 3: exact balance
	items3, balance3, err3 := chargeForLineItem(sub, []subscription{}, 25)
	fmt.Println("test3:", items3, balance3, err3)

	// test 4: one-time usage plan
	items4, balance4, err4 := chargeForLineItem(usage, []oneTimeUsagePlan{}, 10)
	fmt.Println("test4:", items4, balance4, err4)

	// test 5: append onto existing slice
	old := []subscription{
		{
			userEmail: "old@site.com",
			startDate: time.Now(),
			interval:  "monthly",
		},
	}
	items5, balance5, err5 := chargeForLineItem(sub, old, 80)
	fmt.Println("test5:", items5, balance5, err5)
}

/*output:
test1: [{a@site.com {14016466754208329896 167501 0x1028d6720} monthly}] 75 <nil>
test2: [] 0 insufficient funds
test3: [{a@site.com {14016466754208329896 167501 0x1028d6720} monthly}] 0 <nil>
test4: [{c@site.com 100}] 7 <nil>
test5: [{old@site.com {14016466754208604896 442459 0x1028d6720} monthly} {a@site.com {14016466754208329896 167501 0x1028d6720} monthly}] 55 <nil>
➜  generics git:(m*/
