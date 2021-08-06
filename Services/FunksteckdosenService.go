package Services

import (
	database "GartenBewaesserung/Database"
	models "GartenBewaesserung/Models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetList() []models.Funksteckdose {
	list := make([]models.Funksteckdose, 10)

	// GetList from DB
	conn := database.GetConnectionString()
	_, err := gorm.Open(conn.Provider, conn.ConnString)
	//defer db.Close()
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		fmt.Println("hier sollen die Daten geholt werden")
	}

	for index, element := range list {
		fmt.Println("nr: ", index, "Name: ", element.Name+" "+element.Kennung)
		// index is the index where we are
		// element is the element from someSlice for where we are
	}

	return list
}
