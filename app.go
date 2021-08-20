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
	AddFunksteckdoseEndpoints(router)
	AddWasserventilEndpoints(router)
	http.ListenAndServe(":8000", router)

}

func AddFunksteckdoseEndpoints(router *mux.Router) {
	router.HandleFunc("/funksteckdose/create", _services.CreateFunksteckdose).Methods("POST")
	router.HandleFunc("/funksteckdose", _services.GetFunksteckdose).Methods("GET")
	router.HandleFunc("/funksteckdose/list", _services.GetListFunksteckdose).Methods("GET")
	router.HandleFunc("/funksteckdose/delete/{id}", _services.DeleteFunksteckdose).Methods("DELETE")
}

func AddWasserventilEndpoints(router *mux.Router) {
	router.HandleFunc("/wasserventil/create", _services.CreateWasserventil).Methods("POST")
	router.HandleFunc("/wasserventil", _services.GetWasserventil).Methods("GET")
	router.HandleFunc("/wasserventil/list", _services.GetListWasserventile).Methods("GET")
	router.HandleFunc("/wasserventil/delete/{id}", _services.DeleteWasserventil).Methods("DELETE")
}
