package models

type IndividualLoan struct {
	PartnerToken          string     `json:"partner_token"`
	UUID                  string     `json:"uuid"`
	DescriptionLanguageID int64      `json:"description_language_id"`
	ActivityID            string     `json:"activity_id"`
	ThemeTypeID           string     `json:"theme_type_id"`
	Location              string     `json:"location"`
	RepPersonID           string     `json:"rep_person_id"`
	ClientWaiverSigned    bool       `json:"client_waiver_signed"`
	Loanuse               string     `json:"loanuse"`
	Description           string     `json:"description"`
	Entreps               []Entrep   `json:"entreps"`
	Currency              string     `json:"currency"`
	DisburseTime          string     `json:"disburse_time"`
	Schedule              []Schedule `json:"schedule"`
	ImageUrl              string     `json:"image_url"`
}
