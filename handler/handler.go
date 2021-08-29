package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ramrodo/tech-assessment-loan-startup/model"
	"github.com/ramrodo/tech-assessment-loan-startup/service"
	log "github.com/sirupsen/logrus"
)

// CreditAssignment - returns the Investment assigned
func CreditAssignment(w http.ResponseWriter, r *http.Request) {
	var investment model.Investment

	// This is a good way to protect against malicious attacks on your server limiting JSON size
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(fmt.Sprintf("Error at reading request body: %s", err))
	}

	if err := r.Body.Close(); err != nil {
		panic(fmt.Sprintf("Error at closing body: %s", err))
	}

	if err := json.Unmarshal(body, &investment); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // HTTP 422: unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(fmt.Sprintf("Error at unmarshalling body: %s", err))
		}
	}

	investmentService := service.NewCreditAssigner()
	credit_type_300, credit_type_500, credit_type_700, err := investmentService.Assign(investment.Amount)

	if err != nil {
		log.Error(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(400)
	}

	creditTypes := model.CreditTypes{
		CreditType300: credit_type_300,
		CreditType500: credit_type_500,
		CreditType700: credit_type_700,
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creditTypes); err != nil {
		panic(fmt.Sprintf("Error at unmarshalling body: %s", err))
	}

}
