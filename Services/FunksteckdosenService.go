package Services

import (
	database "GartenBewaesserung/Database"
	helper "GartenBewaesserung/Helper"
	models "GartenBewaesserung/Models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	mux "github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// test
// Funksteckdose anlegen
func CreateFunksteckdose(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
		return
	} else {
		defer db.Close()
		model := models.Funksteckdose{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			helper.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}
		model.ErstelltAm = time.Now()
		defer r.Body.Close()
		if err := db.Save(&model).Error; err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(" ******* Steckdose wurde erstellt ")
	}
}

func DeleteFunksteckdose(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
		return
	} else {
		defer db.Close()
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		steckdose := GetFunksteckdoseOr404(db, id, w, r)
		if steckdose == nil {
			return
		}
		if err := db.Delete(&steckdose).Error; err != nil {
			helper.ResponseError(w, http.StatusInternalServerError, err.Error())
			return
		}
		helper.ResponseJSON(w, http.StatusOK, nil)
	}
}

func GetFunksteckdose(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		defer db.Close()
		fmt.Println("-- hole Funksteckdosen --")
		funksteckdosen := []models.Funksteckdose{}
		db.Find(&funksteckdosen)
		fmt.Println(&funksteckdosen)
	}
}

func GetListFunksteckdose(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	listFunksteckdose := []models.Funksteckdose{}
	defer db.Close()
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		fmt.Println("-- List Funksteckdose --")
		db.Find(&listFunksteckdose)
		for index, element := range listFunksteckdose {
			fmt.Println("nr: ", index, "Name: ", element.Name+" "+element.Systemcode)
			// index is the index where we are
			// element is the element from someSlice for where we are
		}
		//	json.NewEncoder(w).Encode(listFunksteckdose)
	}

}
func GetFunksteckdoseOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *models.Funksteckdose {
	funksteckdose := models.Funksteckdose{}
	if err := db.First(&funksteckdose, id).Error; err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &funksteckdose
}
