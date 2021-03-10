package models

type Repayments struct {
	PartnerToken string
	UserID       int64
	Repayments   []Repayment
}
