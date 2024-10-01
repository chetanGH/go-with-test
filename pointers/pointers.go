package pointers

import (
	"errors"
	"fmt"
)

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("Cannot Withdraw, Insufficient Funds!!")

func (w *Wallet) Deposit(amount BitCoin) {
	//amount will be deposited, doesn't return
	w.balance += amount
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
