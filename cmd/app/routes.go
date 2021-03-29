package app

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"kiva/settings"
	"log"
	"net/http"
)

func InitRoutes(){
	router := httprouter.New()
	fmt.Println("init routes")
	//test()
	router.GET("/", SayHello)
	router.POST("/api/kiva/repayments", SendPostPaymentsHandler)
	router.POST("/api/kiva/loan_draft/individual", SendPostIndividualLoan)
	router.POST("/api/kiva/loan_draft/group", SendPostGroupLoan)

	port := fmt.Sprintf(":%s", settings.ReqURL.PortRun)
	log.Println(http.ListenAndServe(port, router))
}

func test()  {

}