package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	t := Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := t.validate()
	if err != nil {
		return nil, err
	}

	err = t.Commit()
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (t *Transaction) Commit() error {
	err := t.AccountFrom.Debit(t.Amount)
	if err != nil {
		return err
	}
	err = t.AccountTo.Credit(t.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transaction) validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient funds")
	}

	return nil
}
