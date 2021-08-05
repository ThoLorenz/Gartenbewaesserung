package Models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Funksteckdose struct {
	gorm.Model
	Name           string `sql:"size:30" gorm:"column:name;not null"`
	Kennung        string `sql:"size:5" gorm:"column:kennung;not null"`
	Status         int    `gorm:"column:status;not null"`
	Systemcode     string `sql:"size:5" gorm:"column:systemcode;not null"`
	DipCode        string `sql:"size:5" gorm:"column:dipStatus;not null"`
	Pulslaenge     int    `gorm:"column:pulslaenge"`
	WasserventilID uint   // die ID kommt über das gorm.Model
}

//var funkList []Funksteckdose

func GetFunksteckdosenListe() {

	// read file
	data, err := ioutil.ReadFile("./Config/Funksteckdosen.json")
	if err != nil {
		fmt.Print(err)
	}

	funkList := make([]Funksteckdose, 0)
	err = json.Unmarshal([]byte(data), &funkList)
	if err == nil {
		for index, element := range funkList {
			fmt.Println("nr: ", index, "Name: ", element.Name)
			// index is the index where we are
			// element is the element from someSlice for where we are
		}
	} else {
		fmt.Println("leer")
	}
}
