package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	name := "John Doe"
	email := "email@email.com"
	c, err := NewClient(name, email)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, name, c.Name)
	assert.Equal(t, email, c.Email)
}

func TestCreateNewClientWhenInvalidArgs(t *testing.T) {
	name := ""
	email := ""
	c, err := NewClient(name, email)
	assert.Nil(t, c)
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")

}
func TestUpdateClient(t *testing.T) {
	name := "John Doe"
	email := "email@email.com"
	c, _ := NewClient(name, email)
	newName := "updated"
	newEmail := "updated@updated.com"
	err := c.Update(newName, newEmail)
	assert.Nil(t, err)
	assert.Equal(t, newName, c.Name)
	assert.Equal(t, newEmail, c.Email)
}

func TestUpdateClientWhenInvalidArgs(t *testing.T) {
	name := "John Doe"
	email := "email@email.com"
	c, _ := NewClient(name, email)
	newName := ""
	newEmail := ""
	err := c.Update(newName, newEmail)
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestAddAccount(t *testing.T) {
	c, _ := NewClient("John Doe", "email@email.com")
	account := NewAccount(c)
	err := c.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(c.Accounts))
	assert.Equal(t, account, c.Accounts[0])
}
