package services

import (
	"bytes"
	"encoding/json"
	"kiva/models"
	"kiva/settings"
	"log"
	"net/http"
)

type KivaService struct {

}

func NewKivaService() *KivaService {
	return &KivaService{}
}

func (receiver *KivaService) SendPostPayments(Body models.Repayments) (error, bool, models.Repayments) {
	client := &http.Client{}
	var newBody models.Repayments
	body, err := json.Marshal(Body)
	if err != nil {
		log.Printf("can't add to body, err is %err", err)
		return err, true, newBody
	}
//	URL := `https://api.kivaws.org/v3/partner/63/repayments`
	URL := settings.ReqURL.SendPostPayments
	req, err := http.NewRequest(
		"POST", URL, bytes.NewBuffer(body),
	)
	if err != nil {
		log.Println("cant create Request", err)
		return err, true, newBody
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request\n")
		return err, true, newBody
	}
	///json.NewDecoder(resp.Body).Decode()
	log.Println(`Status Code = `, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		return nil, false, newBody
	}

	err = json.NewDecoder(resp.Body).Decode(newBody)
	if err != nil {
		log.Println("invalid_json err is ", err)
	}
	return nil, true, newBody
}

func (receiver *KivaService) SendPostIndividualLoan(Body models.IndividualLoan) (error, bool, models.IndividualLoan) {
	client := &http.Client{}
	var newBody models.IndividualLoan
	body, err := json.Marshal(Body)
	if err != nil {
		log.Printf("can't add to body, err is %err", err)
		return err, true, newBody
	}
//	URL := `https://api.kivaws.org/v3/partner/63/loan_draft`
	URL := settings.ReqURL.SendPostIndividualLoan

	req, err := http.NewRequest(
		"POST", URL, bytes.NewBuffer(body),
	)
	if err != nil {
		log.Println("can't create Request", err)
		return err, true, newBody
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request\n")
		return err, true, newBody
	}
	///json.NewDecoder(resp.Body).Decode()
	log.Println(`Status Code = `, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		return nil, false, newBody
	}

	err = json.NewDecoder(resp.Body).Decode(newBody)
	if err != nil {
		log.Println("invalid_json err is ", err)
	}
	return nil, true, newBody
}
func (receiver *KivaService) SendPostGroupLoan(Body models.GroupLoan) (error, bool, models.GroupLoan) {
	client := &http.Client{}
	var newBody models.GroupLoan
	body, err := json.Marshal(Body)
	if err != nil {
		log.Printf("can't add to body, err is %err", err)
		return err, true, newBody
	}
//	URL := `http://localhost:7070/kiva/ws/v1/api.dev.kivaws.org/v3/partner/63/loan_draft`
	URL := settings.ReqURL.SendPostGroupLoan
	req, err := http.NewRequest(
		"POST", URL, bytes.NewBuffer(body),
	)
	if err != nil {
		log.Println("can't create Request", err)
		return err, true, newBody
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request\n")
		return err, true, newBody
	}
	///json.NewDecoder(resp.Body).Decode()
	log.Println(`Status Code = `, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		return nil, false, newBody
	}

	err = json.NewDecoder(resp.Body).Decode(newBody)
	if err != nil {
		log.Println("invalid_json err is ", err)
	}
	return nil, true, newBody
}


/*
 public ResponseEntity<?> sendPostGroupLoan(GroupLoan groupLoan) {
        Loan loan = (Loan)this.restTemplate.postForObject("http://localhost:7070/kiva/ws/v1/api.dev.kivaws.org/v3/partner/63/loan_draft", groupLoan, Loan.class, new Object[0]);
        return new ResponseEntity(loan, HttpStatus.OK);
    }
 */