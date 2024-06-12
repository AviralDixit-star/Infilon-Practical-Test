package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/AviralDixit-star/Infilon-Practical-Test/service"
	"github.com/gorilla/mux"
)

type PersonHandler struct {
	Service service.PersonService
}

func (p *PersonHandler) GetPersonDetailsByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	personID, err := strconv.Atoi(vars["person_id"])
	if err != nil {
		log.Fatal(err)
	}
	personDetails, err := p.Service.GetPersonByID(personID)
	if err != nil {
		log.Fatal(err)
		w.Header().Add("content-Type", "application/json")
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			panic(err)
		}
	}
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(personDetails)
	if err != nil {
		panic(err)
	}
}
