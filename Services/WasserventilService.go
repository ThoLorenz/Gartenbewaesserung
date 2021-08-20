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

// Funksteckdose anlegen
func CreateWasserventil(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
		return
	} else {
		defer db.Close()
		model := models.Wasserventil{}
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
		fmt.Println(" ******* Wasserventil wurde erstellt ")
	}
}

func DeleteWasserventil(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
		return
	} else {
		defer db.Close()
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])
		steckdose := GetWasserventilOr404(db, id, w, r)
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
func GetWasserventil(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		defer db.Close()
		fmt.Println("-- hole Wasserventil --")
		wasserventile := []models.Wasserventil{}
		db.Find(&wasserventile)
		fmt.Println(&wasserventile)
	}
}

func GetListWasserventile(w http.ResponseWriter, r *http.Request) {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	listWasserventile := []models.Wasserventil{}
	defer db.Close()
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		fmt.Println("-- List Wasserventile --")
		db.Find(&listWasserventile)
		for index, element := range listWasserventile {
			fmt.Println("nr: ", index, "Name: ", element.Name+" "+element.Name)
			// index is the index where we are
			// element is the element from someSlice for where we are
		}
		//	json.NewEncoder(w).Encode(listFunksteckdose)
	}
}

func GetWasserventilOr404(db *gorm.DB, id int, w http.ResponseWriter, r *http.Request) *models.Wasserventil {
	wasserventil := models.Wasserventil{}
	if err := db.First(&wasserventil, id).Error; err != nil {
		helper.ResponseError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &wasserventil
}
