package models

type Repayment struct {
	Amount   float64 `json:"amount"`
	ClientID string  `json:"client_id"`
	LoanID   string  `json:"loan_id"`
}

type RepaymentRequest struct {
	Repayments []Repayment `json:"repayments"`
	UserID int64 `json:"user_id"`
}

type Response struct {
	Code string `json:"code"`
	Message string `json:"message"`
}