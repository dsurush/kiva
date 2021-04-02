package models

type Loan struct {
	BorrowerCount             int64  `json:"borrower_count"`
	CreateTime                int64  `json:"create_time"`
	DefaultedTime             int64  `json:"defaulted_time"`
	Delinquent                bool   `json:"delinquent"`
	EndedTime                 int64  `json:"ended_time"`
	ExpiredTime               int64  `json:"expired_time"`
	InternalClientID          string `json:"internal_client_id"`
	InternalLoanID            string `json:"internal_loan_id"`
	IssueFeedbackTime         int64  `json:"issue_feedback_time"`
	IssueReportedBy           string `json:"issue_reported_by"`
	KivaID                    string `json:"kiva_id"`
	LoanCurrency              string `json:"loan_currency"`
	LoanLocalPrice            string `json:"loan_local_price"`
	LoanPrice                 string `json:"loan_price"`
	Location                  string `json:"location"`
	Name                      string `json:"name"`
	Partner                   string `json:"partner"`
	PartnerID                 string `json:"partner_id"`
	PlannedExpirationTime     int64  `json:"planned_expiration_time"`
	PlannedInactiveExpireTime int64  `json:"planned_inactive_expire_time"`
	RefundedTime              int64  `json:"refunded_time"`
	Status                    string `json:"status"`
	StatusDetail              string `json:"status_detail"`
	Uuid                      string `json:"uuid"`
}

type LoanDraftRequest struct {
	ActivityID            int64      `json:"activity_id"`
	ClientWaiverSigned    bool       `json:"client_waiver_signed"`
	Currency              string     `json:"currency"`
	Description           string     `json:"description"`
	DescriptionLanguageID int64      `json:"description_language_id"`
	DisburseTime          string     `json:"disburse_time"`
	Entreps               []Entrep   `json:"entreps"`
	GroupName             string     `json:"group_name"`
	ImageURL              string     `json:"image_url"`
	InternalClientID      string     `json:"internal_client_id"`
	InternalLoanID        string     `json:"internal_loan_id"`
	Loanuse               string     `json:"loanuse"`
	Location              string     `json:"location"`
	NotPictured           []bool     `json:"not_pictured"`
	RepPersonID           int64      `json:"rep_person_id"`
	Schedule              []Schedule `json:"schedule"`
	ThemeTypeID           int64      `json:"theme_type_id"`
}

type LoanResponse struct {
	Data         []Loan `json:"data"`
	TotalRecords int64  `json:"total_records"`
}
