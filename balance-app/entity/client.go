package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Accounts  []*Account
}

func NewClient(name string, email string) (*Client, error) {

	c := Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := c.validate()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Client) Update(name string, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now()
	return c.validate()
}

func (c *Client) AddAccount(acc *Account) error {
	if acc.Client.ID != c.ID {
		return errors.New("account does not belong to this client")
	}
	if acc == nil {
		return errors.New("account is required")
	}
	c.Accounts = append(c.Accounts, acc)
	return nil
}

func (c *Client) validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
