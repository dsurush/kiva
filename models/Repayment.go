package models

type Repayment struct {
	LoanID   string  `json:"loan_id"`
	ClientID string  `json:"client_id"`
	Amount   float64 `json:"amount"`
}
