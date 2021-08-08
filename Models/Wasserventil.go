package Models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//	"gorm.io/gorm"
)

type Wasserventil struct {
	ID                    int    `gorm:"column:id;autoIncrement;type:int"`
	Name                  string `sql:"size:30" gorm:"column:name"`
	Status                int    `gorm:"column:status;not null default 0"`
	Dauer                 int    `gorm:"column:dauer;not null default 1"`
	Durchflussmenge       int    `gorm:"column:durchflussmenge"`
	FeuchtigkeitssensorID int    `gorm:"column:feuchtigkeitssensorId"`
	//	Funksteckdose         Funksteckdose `gorm:"column:funksteckdose;not null"`
	ErstelltAm time.Time `gorm:"column:erstelltAm"`
	GeändertAm time.Time `gorm:"column:geaendertAm"`
	GelöschtAm time.Time `gorm:"column:geloeschtAm"`
	// Hochbeet Hochbeet `gorm:"column:hochbeet;not null"`
	// HochbeetID      int
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
