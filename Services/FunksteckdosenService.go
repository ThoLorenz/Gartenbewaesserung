package Services

import (
	database "GartenBewaesserung/Database"
	models "GartenBewaesserung/Models"
	"encoding/json"
	_ "strconv"
	"time"
	_ "time"

	//	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/gorilla/mux"
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

	model := models.Funksteckdose{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&model); err != nil {
		fmt.Println(http.StatusBadRequest, err.Error())
		return
	}
	model.ErstelltAm = time.Now()
	defer r.Body.Close()
	if err := db.Save(&model).Error; err != nil {
		fmt.Println(w, http.StatusInternalServerError, err.Error())
		return
	}
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

func GetList(w http.ResponseWriter, r *http.Request) {

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

		for index, element := range listFunksteckdose {
			fmt.Println("nr: ", index, "Name: ", element.Kennung+" "+element.Systemcode)
			// index is the index where we are
			// element is the element from someSlice for where we are
		}
		//	json.NewEncoder(w).Encode(listFunksteckdose)
	}

}
