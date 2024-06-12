package app

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/AviralDixit-star/Infilon-Practical-Test/domain"
	"github.com/AviralDixit-star/Infilon-Practical-Test/service"
	"github.com/gorilla/mux"
)

func Start() {

	dbClient := getDbClient()

	//router set up
	router := mux.NewRouter()

	per := PersonHandler{Service: service.NewPersonService(domain.NewPersonRepository(dbClient))}
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))}

	router.HandleFunc("/person/{person_id}", per.GetPersonDetailsByID).Methods(http.MethodGet)

	//server setup
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err)
	}

}

func getDbClient() *sql.DB {
	db, err := sql.Open("mysql", "root:aviral9956@tcp(localhost:3306)/infilon")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
