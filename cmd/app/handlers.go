package app

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"kiva/core/services"
	"kiva/models"
	"log"
	"net/http"
	"strconv"
)

func SayHello(writer http.ResponseWriter, _ *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	//writer.WriteHeader(http.StatusBadRequest)
	_, err := writer.Write([]byte("say hello"))
	if err != nil {
		log.Print(err)
	}
}

/*
   @PostMapping(value={"repayments"})
   public ResponseEntity<?> repayments(HttpServletRequest request) {
       Repayments repayments = new Repayments();
       ArrayList<Repayment> repayment = new ArrayList<Repayment>();
       try {
           String partnerToken = request.getParameter("partner_token");
           Long userId = Long.valueOf(request.getParameter("user_id"));
           Integer rep_count = Integer.valueOf(request.getParameter("rep_count"));
           if (rep_count > 0) {
               for (int i = 0; i < rep_count; ++i) {
                   String loan_id = request.getParameter("loan_id_" + (i + 1));
                   String client_id = request.getParameter("client_id_" + (i + 1));
                   Double amount = Double.valueOf(request.getParameter("amount_" + (i + 1)));
                   repayment.add(new Repayment(loan_id, client_id, amount));
               }
           }
           repayments.setPartner_token(partnerToken);
           repayments.setUser_id(userId);
           repayments.setRepayments(repayment);
           LOGGER.info(repayments + " REP_COUNT = " + rep_count);
           return this.kivaService.sendPostRepayments(repayments);
       }
       catch (Exception e) {
           LOGGER.error(e);
           return new ResponseEntity(repayments, HttpStatus.OK);
       }
   }
*/

func SendPostPaymentsHandler(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var repayments models.Repayments
	service := services.NewKivaService()
	repayments.PartnerToken = request.Header.Get("partner_token")
	userIdString := request.Header.Get("user_id")
	userIdInt, err := strconv.Atoi(userIdString)
	repayments.UserID = int64(userIdInt)
	if err != nil {
		log.Println("can't convert")
		return
	}
	repCountString := request.Header.Get("rep_count")
	repCount, err := strconv.Atoi(repCountString)
	var repayment models.Repayment
	if repCount > 0 {
		for i := 1; i <= repCount; i++ {

			repayment.LoanID = request.Header.Get(fmt.Sprintf("loan_id_%d", i))
			repayment.ClientID = request.Header.Get(fmt.Sprintf("client_id_%d", i))
			AmountString := request.Header.Get(fmt.Sprintf("amount_%d", i))
			repayment.Amount, err = strconv.ParseFloat(AmountString, 64)
			if err != nil {
				log.Println("can't convert to Float")
				return
			}
			repayments.Repayments = append(repayments.Repayments, repayment)
		}
	}

	err, status, responseBody := service.SendPostPayments(repayments)
	if err != nil {
		log.Println("can't send request err is ", err)
	}
	if !status{
		log.Println("status not 200ok")
	}
	err = json.NewEncoder(writer).Encode(&responseBody)
	if err != nil {
		log.Println(err)
	}
	return
}

func SendPostIndividualLoan(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	return
}
