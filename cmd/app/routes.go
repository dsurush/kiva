package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func init(){
	router := httprouter.New()
	//test()
	router.GET("/", SayHello)
	router.POST("/api/kiva/repayments", SendPostPaymentsHandler)
	router.POST("/api/kiva/loan_draft/individual", SendPostIndividualLoan)
	router.POST("/api/kiva/loan_draft/group", SendPostGroupLoan)
	port := fmt.Sprintf(":%d", 8888)
	log.Println(http.ListenAndServe(port, router))
}

func test()  {

}