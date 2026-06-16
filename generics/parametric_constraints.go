package main

import (
	"fmt"
)

type biller[C customer] interface {
	Charge(C) bill
	Name() string
}

type userBiller struct {
	Plan string
}

func (ub userBiller) Charge(u user) bill {
	amount := 50.0
	if ub.Plan == "pro" {
		amount = 100.0
	}
	return bill{
		Customer: u,
		Amount:   amount,
	}
}

func (sb userBiller) Name() string {
	return fmt.Sprintf("%s user biller", sb.Plan)
}

type orgBiller struct {
	Plan string
}

func (ob orgBiller) Name() string {
	return fmt.Sprintf("%s org biller", ob.Plan)
}

func (ob orgBiller) Charge(o org) bill {
	amount := 2000.0
	if ob.Plan == "pro" {
		amount = 3000.0
	}
	return bill{
		Customer: o,
		Amount:   amount,
	}
}

type customer interface {
	GetBillingEmail() string
}

type bill struct {
	Customer customer
	Amount   float64
}

type user struct {
	UserEmail string
}

func (u user) GetBillingEmail() string {
	return u.UserEmail
}

type org struct {
	Admin user
	Name  string
}

func (o org) GetBillingEmail() string {
	return o.Admin.GetBillingEmail()
}

func main() {
	// Test userBiller with basic plan
	ub := userBiller{Plan: "basic"}
	u := user{UserEmail: "alice@example.com"}
	b := ub.Charge(u)
	fmt.Println(ub.Name())                    // "basic user biller"
	fmt.Println(b.Amount)                     // 50
	fmt.Println(b.Customer.GetBillingEmail()) // "alice@example.com"

	// Test userBiller with pro plan
	ubPro := userBiller{Plan: "pro"}
	bPro := ubPro.Charge(u)
	fmt.Println(ubPro.Name()) // "pro user biller"
	fmt.Println(bPro.Amount)  // 100

	// Test orgBiller with basic plan
	ob := orgBiller{Plan: "basic"}
	o := org{Admin: user{UserEmail: "admin@corp.com"}, Name: "Corp"}
	bo := ob.Charge(o)
	fmt.Println(ob.Name())                     // "basic org biller"
	fmt.Println(bo.Amount)                     // 2000
	fmt.Println(bo.Customer.GetBillingEmail()) // "admin@corp.com"

	// Test orgBiller with pro plan
	obPro := orgBiller{Plan: "pro"}
	boPro := obPro.Charge(o)
	fmt.Println(obPro.Name()) // "pro org biller"
	fmt.Println(boPro.Amount) // 3000

	// Test that biller[user] and biller[org] are satisfied
	var _ biller[user] = userBiller{Plan: "basic"}
	var _ biller[org] = orgBiller{Plan: "pro"}
	fmt.Println("Interface constraints satisfied!")
}

/*output:
basic user biller
50
alice@example.com
pro user biller
100
basic org biller
2000
admin@corp.com
pro org biller
3000
Interface constraints satisfied!*/
