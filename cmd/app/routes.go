package app

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"kiva/models"
	"log"
	"net/http"
)

func init(){
	router := httprouter.New()
	test()
	router.GET("/", SayHello)
	router.POST("/api/kiva/repayments", SendPostPaymentsHandler)
	router.POST("/api/kiva/loan_draft/individual", SendPostIndividualLoan)
	router.POST("/api/kiva/loan_draft/group", SendPostGroupLoan)
	port := fmt.Sprintf(":%d", 8888)
	log.Println(http.ListenAndServe(port, router))
}

func test()  {
	var a models.Repayment
	a.Amount = 1
	a.ClientID = "ad"
	a.LoanID = "ab"
	bytes, err := json.Marshal(a)
	if err != nil {
		log.Fatalln("ERROR IS = ", err)
	}
	fmt.Println(string(bytes))
}