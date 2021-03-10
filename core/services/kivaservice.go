package services

import (
	"bytes"
	"encoding/json"
	"kiva/models"
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
	URL := `https://api.kivaws.org/v3/partner/63/repayments`
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

/*
public ResponseEntity<?> sendPostIndividualLoan(IndividualLoan individualLoan) {
        Loan loan = null;
        try {
            LOGGER.info("got params for request\n" + individualLoan);
            loan = (Loan)this.restTemplate.postForObject("https://api.kivaws.org/v3/partner/63/loan_draft", individualLoan, Loan.class, new Object[0]);
            LOGGER.info("got answer for response\n" + loan);
            return new ResponseEntity(loan, HttpStatus.OK);
        }
        catch (Exception e) {
            this.str = e.getMessage();
            if (this.str.contains("Could not read document: Unrecognized token 'Success': was expecting ('true', 'false' or 'null')")) {
                LOGGER.info("SUCCESS");
                return new ResponseEntity(loan, HttpStatus.OK);
            }
            LOGGER.error(e);
            return new ResponseEntity(loan, HttpStatus.INTERNAL_SERVER_ERROR);
        }
    }
 */

func (receiver *KivaService) SendPostIndividualLoan(Body models.IndividualLoan) (err error, ok bool) {
	client := &http.Client{}
	var newBody models.Repayments
	body, err := json.Marshal(Body)
	if err != nil {
		log.Printf("can't add to body, err is %err", err)
		return err, true
	}
	URL := `https://api.kivaws.org/v3/partner/63/loan_draft`
	req, err := http.NewRequest(
		"POST", URL, bytes.NewBuffer(body),
	)
	if err != nil {
		log.Println("can't create Request", err)
		return err, true
	}
	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request\n")
		return err, true
	}
	///json.NewDecoder(resp.Body).Decode()
	log.Println(`Status Code = `, resp.StatusCode)
	if resp.StatusCode == http.StatusOK {
		return nil, false
	}

	err = json.NewDecoder(resp.Body).Decode(newBody)
	if err != nil {

	}
	return nil, true
}