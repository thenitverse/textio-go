package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(c *customer, t transaction) error {
	if t.transactionType == transactionWithdrawal {
		if c.balance < t.amount {
			return errors.New("insufficient funds")
		}
		c.balance -= t.amount
	} else if t.transactionType == transactionDeposit {
		c.balance += t.amount
	} else {
		return errors.New("unknown transaction type")
	}
	return nil
}

func main() {
	// Test 1: Deposit
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit}
	err := updateBalance(&alice, deposit)
	fmt.Printf("Test 1 - Deposit: balance=%.2f, err=%v\n", alice.balance, err)
	// Expected: balance=150.00, err=<nil>

	// Test 2: Withdrawal with sufficient funds
	bob := customer{id: 2, balance: 200.0}
	withdrawal := transaction{customerID: 2, amount: 100.0, transactionType: transactionWithdrawal}
	err = updateBalance(&bob, withdrawal)
	fmt.Printf("Test 2 - Withdrawal: balance=%.2f, err=%v\n", bob.balance, err)
	// Expected: balance=100.00, err=<nil>

	// Test 3: Withdrawal with insufficient funds
	carol := customer{id: 3, balance: 50.0}
	overdraft := transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal}
	err = updateBalance(&carol, overdraft)
	fmt.Printf("Test 3 - Insufficient funds: balance=%.2f, err=%v\n", carol.balance, err)
	// Expected: balance=50.00, err=insufficient funds

	// Test 4: Unknown transaction type
	dave := customer{id: 4, balance: 100.0}
	unknown := transaction{customerID: 4, amount: 50.0, transactionType: "transfer"}
	err = updateBalance(&dave, unknown)
	fmt.Printf("Test 4 - Unknown type: balance=%.2f, err=%v\n", dave.balance, err)
	// Expected: balance=100.00, err=unknown transaction type
}

/*output:
Test 1 - Deposit: balance=150.00, err=<nil>
Test 2 - Withdrawal: balance=100.00, err=<nil>
Test 3 - Insufficient funds: balance=50.00, err=insufficient funds
Test 4 - Unknown type: balance=100.00, err=unknown transaction type
*/
