package Models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Funksteckdose struct {
	ID             int    `gorm:"column:id;autoIncrement;type:int"`
	Name           string `sql:"size:30" gorm:"column:name;not null"`
	Kennung        string `sql:"size:5" gorm:"column:kennung;not null"`
	Status         int    `gorm:"column:status;not null"`
	Systemcode     string `sql:"size:5" gorm:"column:systemcode;not null"`
	DipCode        string `sql:"size:5" gorm:"column:dipStatus;not null"`
	Pulslaenge     int    `gorm:"column:pulslaenge"`
	WasserventilID int    `gorm:"column:wasserventilId"`
	Wasserventil   Wasserventil
	PumpeID        int
	Pumpe          Pumpe
	ErstelltAm     time.Time `gorm:"column:erstelltAm"`
	GeändertAm     time.Time `gorm:"column:geaendertAm"`
	GelöschtAm     time.Time `gorm:"column:geloeschtAm"`
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
