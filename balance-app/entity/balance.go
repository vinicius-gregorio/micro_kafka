package entity

import (
	"errors"
)

type Balance struct {
	AccountID string
	Amount    float64
}

func NewBalance(b *Balance) *Balance {
	if b == nil {
		return nil
	}
	newBalance := Balance{
		AccountID: b.AccountID,
		Amount:    b.Amount,
	}
	return &newBalance
}

func (a *Balance) Credit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero when crediting")
	}
	a.Amount += amount
	return nil
}
func (a *Balance) Debit(amount float64) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero when debiting")
	}
	if a.Amount < amount {
		return errors.New("insufficient funds")
	}
	a.Amount -= amount
	return nil
}
