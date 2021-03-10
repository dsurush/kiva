package models

type Entrep struct {
	ClientID  string `json:"client_id"`
	LoanID    string `json:"loan_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Amount    string `json:"amount"`
}
