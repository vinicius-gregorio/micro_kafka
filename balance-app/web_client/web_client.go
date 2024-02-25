package web_client

import (
	"encoding/json"
	"net/http"

	"github.com/vinicius-gregorio/balance-api/database"
)

type Server struct {
	balanceDB *database.BalanceDB
}

func NewServer(balanceDB *database.BalanceDB) *Server {
	return &Server{
		balanceDB: balanceDB,
	}
}

func (s *Server) Start() {
	http.HandleFunc("/balances/", s.getBalanceByID)
	http.ListenAndServe(":3003", nil)
}

func (s *Server) getBalanceByID(w http.ResponseWriter, r *http.Request) {
	// Extract account ID from URL path
	accountID := r.URL.Path[len("/balances/"):]

	// Retrieve balance from database
	balance, err := s.balanceDB.FindByID(accountID)
	if err != nil {
		http.Error(w, "Error retrieving balance", http.StatusInternalServerError)
		return
	}
	if balance == nil {
		http.Error(w, "Balance not found", http.StatusNotFound)
		return
	}

	// Marshal balance to JSON
	responseJSON, err := json.Marshal(balance)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
