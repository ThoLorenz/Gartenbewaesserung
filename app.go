package main

import (
	database "GartenBewaesserung/Database"
	_ "GartenBewaesserung/Models"
	funksteckdoseService "GartenBewaesserung/Services"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("starten..")
	database.InitDatabase()
	router := mux.NewRouter()
	router.HandleFunc("/funksteckdose/create", funksteckdoseService.Create).Methods("POST")
	router.HandleFunc("/funksteckdose", funksteckdoseService.Get).Methods("GET")
	router.HandleFunc("/funksteckdose/list", funksteckdoseService.GetList).Methods("GET")
	http.ListenAndServe(":8000", router)
	//	funksteckdoseService.Create()
	//Models.GetFunksteckdosenListe()
}
