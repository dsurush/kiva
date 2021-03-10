package models

type GroupLoan struct {
	PartnerToken          string     `json:"partner_token"`
	UUID                  string     `json:"uuid"`
	DescriptionLanguageID int64      `json:"description_language_id"`
	ThemeTypeID           string     `json:"theme_type_id"`
	ActivityID            string     `json:"activity_id"`
	Location              string     `json:"location"`
	RepPersonID           string     `json:"rep_person_id"`
	InternaleClientID     string     `json:"internale_client_id"`
	InternalLoanID        string     `json:"internal_loan_id"`
	GroupName             string     `json:"group_name"`
	ClientWaiverSigned    bool       `json:"client_waiver_signed"`
	Loanuse               string     `json:"loanuse"`
	Description           string     `json:"description"`
	Entreps               []Entrep   `json:"entreps"`
	NotPictured           []bool     `json:"not_pictured"`
	Currency              string     `json:"currency"`
	DisburseTime          string     `json:"disburse_time"`
	Schedule              []Schedule `json:"schedule"`
	ImageUrl              string     `json:"image_url"`
}
