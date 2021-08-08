package Models

import (
	"fmt"
	"time"
	//"gorm.io/gorm"
)

type Feuchtigkeitssensor struct {
	ID              int          `gorm:"column:id;autoIncrement;type:int"`
	Name            string       `sql:"size:30" gorm:"column:name;not null"`
	HochbeetID      int          `gorm:"column:hochbeetID;not null"`
	Wasserventil    Wasserventil `gorm:"column:wasserventil;not null"`
	MinFeuchtigkeit int          `gorm:"column:maxFeuchtigkeit;not null"`
	MaxFeuchtigkeit int          `gorm:"column:minFeuchtigkeit;not null"`
	Temperatur      int          `gorm:"column:temperatur"`
	ErstelltAm      time.Time    `gorm:"column:erstelltAm"`
	GeändertAm      time.Time    `gorm:"column:geaendertAm"`
	GelöschtAm      time.Time    `gorm:"column:geloeschtAm"`
}

func GetFeuchtigkeitssensorListe() {
	fmt.Println("Erzeuge Feuchtsensor")
}
