package main

import (
	_database "GartenBewaesserung/Database"
	_services "GartenBewaesserung/Services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	_database.InitDatabase()
	_services.CreateStandardObjectsForDb()
	AddHttpEndpoints()

	//	funksteckdoseService.Create()
	//Models.GetFunksteckdosenListe()
}

func AddHttpEndpoints() {
	fmt.Println(" ... erstelle API-Endpunkte ..")
	router := mux.NewRouter()
	router.HandleFunc("/funksteckdose/create", _services.Create).Methods("POST")
	router.HandleFunc("/funksteckdose", _services.Get).Methods("GET")
	router.HandleFunc("/funksteckdose/list", _services.GetList).Methods("GET")
	router.HandleFunc("/funksteckdose/delete/{id}", _services.Delete).Methods("DELETE")

	http.ListenAndServe(":8000", router)

}
