// Package account provides banck account mechanism
package account

import (
	"sync"
)

// BankOps defines bank operations
type BankOps interface {
	Close() (int64, bool)
	Balance() (int64, bool)
	Deposit(int64) (int64, bool)
}

// Account implements bank operations
type Account struct {
	amount int64
	closed bool
	sync.Mutex
}

// Open a account with positive init, else nil
func Open(init int64) *Account {
	if init < 0 {
		return nil
	}

	return &Account{amount: init, closed: false}
}

// Close account
func (acc *Account) Close() (int64, bool) {
	acc.Lock()
	defer acc.Unlock()
	if acc.closed {
		return 0, false
	}
	acc.closed = true
	return acc.amount, true
}

// Balance account's amount
func (acc *Account) Balance() (int64, bool) {
	return acc.amount, !acc.closed
}

// Deposit account
func (acc *Account) Deposit(depo int64) (int64, bool) {
	acc.Lock()
	defer acc.Unlock()
	if acc.closed || acc.amount+depo < 0 {
		return depo, false
	}

	acc.amount += depo
	return acc.amount, true
}
