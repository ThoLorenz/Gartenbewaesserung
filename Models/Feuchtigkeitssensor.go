package Models

import (
	"fmt"
	"time"
	//"gorm.io/gorm"
)

type Feuchtigkeitssensor struct {
	ID              int       `gorm:"column:id;autoIncrement;type:int"`
	Name            string    `sql:"size:30" gorm:"column:name;not null"`
	MinFeuchtigkeit int       `gorm:"column:maxFeuchtigkeit;type:int"`
	MaxFeuchtigkeit int       `gorm:"column:minFeuchtigkeit;type:int"`
	Temperatur      int       `gorm:"column:temperatur;type:int"`
	ErstelltAm      time.Time `gorm:"column:erstelltAm"`
	GeändertAm      time.Time `gorm:"column:geaendertAm"`
	GelöschtAm      time.Time `gorm:"column:geloeschtAm"`
	HochbeetID      int       `gorm:"column:hochbeetID;not null;type:int"`
	Wasserventil    Wasserventil
	WasserventilID  int `gorm:"columns:wasserventilId"`
}

func GetFeuchtigkeitssensorListe() {
	fmt.Println("Erzeuge Feuchtsensor")
}
