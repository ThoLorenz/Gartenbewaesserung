package Models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jinzhu/gorm"
)

type Wasserventil struct {
	gorm.Model
	Name            string        `sql:"size:30" gorm:"column:name;not null"`
	Status          int           `gorm:"column:status;not null"`
	Dauer           int           `gorm:"column:dauer;not null"`
	Durchflussmenge int           `gorm:"column:durchflussmenge"`
	Funksteckdose   Funksteckdose `gorm:"column:funksteckdose;not null"`
	HochbeetID      uint          `gorm:"column:hochbeetId;not null"`
}

// Auslesen der Config.Zuleitungsventil.json
func SetVentil() {
	// read file
	data, err := ioutil.ReadFile("./Config/Zuleitungsventil.json")
	if err != nil {
		fmt.Print(err)
	}

	var Zuleitungsventil Wasserventil
	// unmarshall it
	err = json.Unmarshal(data, &Zuleitungsventil)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("Name: ", Zuleitungsventil.Name)
	fmt.Println("Status: ", Zuleitungsventil.Status)
	fmt.Println("Dauer: ", Zuleitungsventil.Dauer)

}
