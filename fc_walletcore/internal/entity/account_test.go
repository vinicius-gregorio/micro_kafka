package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "email@email.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, float64(0), account.Balance)

}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	c, _ := NewClient("John Doe", "email@email.com")

	account := NewAccount(c)
	err := account.Credit(100)
	assert.Nil(t, err)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	c, _ := NewClient("John Doe", "email@email.com")

	account := NewAccount(c)
	account.Credit(100)
	err := account.Debit(50)
	assert.Nil(t, err)
	assert.Equal(t, float64(50), account.Balance)
}

//	TODO: add more tests
