package Models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type Feuchtigkeitssensor struct {
	gorm.Model
	Name            string `sql:"size:30" gorm:"column:name;not null"`
	HochbeetID      uint   `gorm:"column:hochbeetID;not null"`
	MinFeuchtigkeit int    `gorm:"column:maxFeuchtigkeit;not null"`
	MaxFeuchtigkeit int    `gorm:"column:minFeuchtigkeit;not null"`
	Temperatur      int    `gorm:"column:temperatur"`
}

func GetFeuchtigkeitssensorListe() {
	fmt.Println("Erzeuge Feuchtsensor")
}
