package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	amount Bitcoin
}

type Stringer interface {
	String() string
}

var Err = errors.New("NO")

func (b Bitcoin) String() (s string) {
	s = fmt.Sprintf("%d BTC", b)
	return
}

func (w *Wallet) Deposit(n Bitcoin) {
	w.amount += n
}

func (w *Wallet) Withdraw(n Bitcoin) error {
	if n > w.amount {
		return Err
	}
	w.amount -= n
	return nil
}

func (w Wallet) Balance() (res Bitcoin) {
	res = w.amount
	return
}
