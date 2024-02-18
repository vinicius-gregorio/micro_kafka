package database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"github.com/vinicius-gregorio/fc_walletcore/internal/entity"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestSave() {
	c := &entity.Client{
		ID:    "123",
		Name:  "John Doe",
		Email: "email@email.com",
	}
	err := s.clientDB.Save(c)
	s.Nil(err)

	//TODO: add more tests
}
func (s *ClientDBTestSuite) TestGet() {
	c, _ := entity.NewClient("John Doe", "email@email.com")
	s.clientDB.Save(c)
	cDB, err := s.clientDB.Get(c.ID)
	s.Nil(err)
	s.Equal(c.ID, cDB.ID)
}
