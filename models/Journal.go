package models

type Journal struct {
	Body             string `json:"body"`
	ImageURL         string `json:"image_url"`
	InternalClientID string `json:"internal_client_id"`
	InternalLoanID   string `json:"internal_loan_id"`
	Subject          string `json:"subject"`
}

type JournalRequest struct {
	Journals []Journal `json:"journals"`
}

type JournalResponse struct {
	ArchiveID string `json:"archive_id"`
	ConfirmURL string `json:"confirm_url"`
}