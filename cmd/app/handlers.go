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
/*
public ResponseEntity<?> loanIndividual(HttpServletRequest request) {
        IndividualLoan loan = new IndividualLoan();
        ArrayList<Schedule> schedules = new ArrayList<Schedule>();
        ArrayList<Entrep> entreps = new ArrayList<Entrep>();
        String partnerToken = request.getParameter("partner_token");
        String uuid = request.getParameter("uuid");
        Integer description_language_id = Integer.valueOf(request.getParameter("description_language_id"));
        String activity_id = request.getParameter("activity_id");
        String theme_type_id = request.getParameter("theme_type_id");
        String location = request.getParameter("location");

        String rep_person_id = request.getParameter("rep_person_id");
        Boolean client_waiver_signed = Boolean.valueOf(request.getParameter("client_waiver_signed"));
        String loanuse = request.getParameter("loanuse");
        String description = request.getParameter("description");
        String currency = request.getParameter("currency");
        String disburse_time = request.getParameter("disburse_time");
        String image_url = request.getParameter("image_url");
        String client_id = request.getParameter("client_id");
        String loan_id = request.getParameter("loan_id");
        String first_name = request.getParameter("first_name");
        String last_name = request.getParameter("last_name");
        String gender = request.getParameter("gender");
        Double amount = Double.valueOf(request.getParameter("amount"));
        entreps.add(new Entrep(client_id, loan_id, first_name, last_name, gender, amount));
        Integer schedule_count = Integer.valueOf(request.getParameter("schedule_count"));
        if (schedule_count > 0) {
            for (int i = 0; i < schedule_count; ++i) {
                String date = request.getParameter("date_" + (i + 1));
                Double principal = Double.valueOf(request.getParameter("principal_" + (i + 1)));
                Double interest = Double.valueOf(request.getParameter("interest_" + (i + 1)));
                schedules.add(new Schedule(date, principal, interest));
            }
        }
        loan.setActivity_id(activity_id);
        loan.setClient_waiver_signed(client_waiver_signed);
        loan.setCurrency(currency);
        loan.setDescription(description);
        loan.setDescription_language_id(description_language_id);
        loan.setDisburse_time(disburse_time);
        loan.setImage_url(image_url);
        loan.setLoanuse(loanuse);
        loan.setLocation(location);
        loan.setPartner_token(partnerToken);
        loan.setUuid(uuid);
        loan.setTheme_type_id(theme_type_id);
        loan.setRep_person_id(rep_person_id);
        loan.setSchedule(schedules);
        loan.setEntreps(entreps);
        LOGGER.info(loan);
        return this.kivaService.sendPostIndividualLoan(loan);
    }
 */

func SendPostIndividualLoan(writer http.ResponseWriter, request *http.Request, pr httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var loan models.IndividualLoan
	var schedules []models.Schedule
	var entreps []models.Entrep
	service := services.NewKivaService()

	loan.PartnerToken = request.Header.Get("partner_token")
	loan.UUID = request.Header.Get("uuid")
	DescriptionLanguageID, err := strconv.Atoi(request.Header.Get("description_language_id"))
	if err != nil {
		log.Println("can't convert Description Language ID err is ", err)
		return
	}
	loan.DescriptionLanguageID = int64(DescriptionLanguageID)

	loan.ActivityID = request.Header.Get("activity_id")
	loan.ThemeTypeID = request.Header.Get("theme_type_id")
	loan.Location = request.Header.Get("location")
	loan.RepPersonID = request.Header.Get("rep_person_id")
	loan.ClientWaiverSigned, err = strconv.ParseBool(request.Header.Get("client_waiver_signed"))
	if err != nil {
		log.Println("can't conver to boolean client_waiver_signed err is ", err)
		return
	}
	loan.Loanuse = request.Header.Get("loanuse")
	loan.Description = request.Header.Get("description")
	loan.Currency = request.Header.Get("currency")
	loan.DisburseTime = request.Header.Get("disburse_time")
	loan.ImageUrl = request.Header.Get("image_url")
	ClientID := request.Header.Get("client_id")
	LoanID := request.Header.Get("loan_id")
	FirstName := request.Header.Get("first_name")
	LastName := request.Header.Get("last_name")
	Gender := request.Header.Get("gender")
	Amount, err := strconv.ParseFloat(request.Header.Get("amount"), 64)
	if err != nil {
		log.Println("can't parse amount err is ", err)
		return
	}
	newEntrep := models.Entrep{
		ClientID:  ClientID,
		LoanID:    LoanID,
		FirstName: FirstName,
		LastName:  LastName,
		Gender:    Gender,
		Amount:    Amount,
	}
	entreps = append(entreps, newEntrep)
	ScheduleCount, err := strconv.Atoi(request.Header.Get("schedule_count"))
	if err != nil {
		log.Println("can't convert ScheduleCount err is ", err)
		return
	}
	if ScheduleCount > 0{
		for i := 0; i < ScheduleCount; i++{
			var Schedule models.Schedule
			Schedule.Date = request.Header.Get(fmt.Sprintf("date_%d", i + 1))
			Schedule.Principal, err = strconv.ParseFloat(request.Header.Get(fmt.Sprintf("principal_%d", i + 1)), 64)
			Schedule.Interest, err = strconv.ParseFloat(request.Header.Get(fmt.Sprintf("interest_%d", i + 1)), 64)
			schedules = append(schedules, Schedule)
		}
	}
	err, ok, response := service.SendPostIndividualLoan(loan)
	/*
	        return this.kivaService.sendPostIndividualLoan(loan);
	 */
	return
}
