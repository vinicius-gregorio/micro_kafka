package database

import (
	"database/sql"
	"fmt"

	"github.com/vinicius-gregorio/balance-api/entity"
)

type BalanceDB struct {
	DB *sql.DB
}

func NewBalanceDB(db *sql.DB) *BalanceDB {
	return &BalanceDB{
		DB: db,
	}
}

func (b *BalanceDB) FindByID(id string) (*entity.Balance, error) {
	var balance entity.Balance

	// Prepare the SQL query to select balance by account ID
	query := "SELECT account_id, amount FROM balances WHERE account_id = ?"

	// Query the database
	row := b.DB.QueryRow(query, id)
	err := row.Scan(&balance.AccountID, &balance.Amount)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no rows were returned, return nil for the balance and nil error
			return nil, nil
		}
		// If there was an error other than no rows, return the error
		return nil, err
	}

	// If a balance was found, return it
	return &balance, nil
}

func (b *BalanceDB) Save(balance *entity.Balance) error {
	// Check if a record with the same account_id already exists
	existingBalance, err := b.FindByID(balance.AccountID)
	if err != nil {
		fmt.Println("Error fetching existing balance:", err)
		return err
	}

	// If a record with the same account_id exists, update the amount
	if existingBalance != nil {
		existingBalance.Amount = balance.Amount
		return b.UpdateBalance(existingBalance)
	}

	// If no record with the same account_id exists, insert a new record
	stmt, err := b.DB.Prepare("INSERT INTO balances (account_id, amount) VALUES (?, ?)")
	if err != nil {
		fmt.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.AccountID, balance.Amount)
	if err != nil {
		fmt.Println("Error executing statement:", err)
		return err
	}
	return nil
}

func (b *BalanceDB) UpdateBalance(balance *entity.Balance) error {
	stmt, err := b.DB.Prepare("UPDATE balances SET amount = ? WHERE account_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(balance.Amount, balance.AccountID)
	if err != nil {
		return err
	}
	return nil
}
