package models

type Entrep struct {
	Amount    float64 `json:"amount"`
	ClientID  string `json:"client_id"`
	FirstName string `json:"first_name"`
	Gender    string `json:"gender"`
	LastName  string `json:"last_name"`
	LoanID    string `json:"loan_id"`
}
