package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	c1, _ := NewClient("John Doe", "email@email.com")
	a1 := NewAccount(c1)

	c2, _ := NewClient("John Doe2", "email2@email.com")
	a2 := NewAccount(c2)

	a1.Credit(1000)
	a2.Credit(50)

	transact, err := NewTransaction(a1, a2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transact)
	assert.Equal(t, float64(150), a2.Balance)
	assert.Equal(t, float64(900), a1.Balance)
}

func TestCreateTransactionWithInsuficientFunds(t *testing.T) {
	c1, _ := NewClient("John Doe", "email@email.com")
	a1 := NewAccount(c1)

	c2, _ := NewClient("John Doe2", "email2@email.com")
	a2 := NewAccount(c2)

	a1.Credit(1000)
	a2.Credit(50)
	transact, err := NewTransaction(a1, a2, 2000)
	assert.NotNil(t, err)
	assert.Nil(t, transact)
	assert.Error(t, err, "insufficient funds")
	assert.Equal(t, float64(50), a2.Balance)
	assert.Equal(t, float64(1000), a1.Balance)

}
