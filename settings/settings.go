package settings

type RequestsURL struct {
	SendPostPayments string `json:"send_post_payments"`
	SendPostIndividualLoan string `json:"send_post_individual_loan"`
	SendPostGroupLoan string `json:"send_post_group_loan"`
	PortRun string `json:"port_run"`
}
