package Services

import (
	models "GartenBewaesserung/Models"
	"fmt"
)

func GetList() []models.Funksteckdose {
	list := make([]models.Funksteckdose, 10)

	// GetList from DB

	for index, element := range list {
		fmt.Println("nr: ", index, "Name: ", element.Name+" "+element.Kennung)
		// index is the index where we are
		// element is the element from someSlice for where we are
	}

	return list
}
