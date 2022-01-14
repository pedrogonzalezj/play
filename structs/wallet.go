package structs

import (
	"errors"

	"github.com/pedrogonzalezj/play/custom"
)

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance custom.Bitcoin
}

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount custom.Bitcoin) {
	w.balance += amount
}

// ErrInsufficientFunds means a wallet does not have enough Bitcoin to perform a withdraw.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw subtracts some Bitcoin from the wallet, returning an error if it cannot be performed.
func (w *Wallet) Withdraw(amount custom.Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance returns the number of Bitcoin a wallet has.
func (w *Wallet) Balance() custom.Bitcoin {
	return w.balance
}
