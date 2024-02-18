package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"github.com/vinicius-gregorio/fc_walletcore/internal/entity"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	transactionDB *TransactionDB
}

func (s *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	db.Exec("CREATE TABLE transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount float, created_at date)")

	/// Creating clients
	c1, err := entity.NewClient("John Doe", "email@email.com")
	s.Nil(err)
	s.client1 = c1
	c2, err := entity.NewClient("Jane Doe2", "email2@email.com")
	s.Nil(err)
	s.client2 = c2

	/// Creating accounts
	accFrom := entity.NewAccount(s.client1)
	accFrom.Balance = 1000
	s.accountFrom = accFrom

	accTo := entity.NewAccount(s.client2)
	accTo.Balance = 50
	s.accountTo = accTo

	s.transactionDB = NewTransactionDB(db)

}

func (s *TransactionDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE transactions")

}

func TestTransactionDBTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (s *TransactionDBTestSuite) TestCreate() {
	tr, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.transactionDB.Create(tr)
	s.Nil(err)
}
