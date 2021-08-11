package Services

import (
	database "GartenBewaesserung/Database"
	models "GartenBewaesserung/Models"
	"strconv"
	"time"

	//	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Create(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		fmt.Println("-- erstelle Funksteckdose --")
	}
	vars := mux.Vars(r)
	name := vars["Name"]
	kennung := vars["Kennung"]
	status, err := strconv.Atoi(vars["Status"])
	if err != nil {
		log.Println("Fehler bei status : ", err)
	}
	systemcode := vars["Systemcode"]
	dipcode := vars["DipCode"]

	pulslaenge, err := strconv.Atoi(vars["Pulslaenge"])
	if err != nil {
		log.Println("Fehler bei status : ", err)
	}
	erstelltAm := time.Now()

	db.Create(&models.Funksteckdose{Name: name, Kennung: kennung, Status: status, Systemcode: systemcode,
		DipCode: dipcode, Pulslaenge: pulslaenge, ErstelltAm: erstelltAm})
	fmt.Println(" ******* Steckdose wurde erstellt ")

}

func Get(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		fmt.Println("-- hole Funksteckdosen --")

		funksteckdosen := []models.Funksteckdose{}
		db.Find(&funksteckdosen)
		fmt.Println(&funksteckdosen)
	}
}

func GetList() []models.Funksteckdose {

	// GetList from DB
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	listFunksteckdose := []models.Funksteckdose{}
	defer db.Close()
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		fmt.Println("-- Daten holen (Funksteckdose) --")

		db.Find(&listFunksteckdose)

		fmt.Println("{}", listFunksteckdose)
		//	json.NewEncoder(w).Encode(listFunksteckdose)
	}
	// for index, element := range listFunksteckdose {
	// 	fmt.Println("nr: ", index, "Name: ", element.Kennung+" "+element.Systemcode)
	// 	// index is the index where we are
	// 	// element is the element from someSlice for where we are
	// }

	return listFunksteckdose
}
