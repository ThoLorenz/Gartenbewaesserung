package Models

import (
	_ "fmt"
	"time"
)

type Hochbeet struct {
	ID           int    `gorm:"column:id;autoIncrement;type:int"`
	Name         string `sql:"size:30" gorm:"column:name;not null"`
	ListSensoren []Feuchtigkeitssensor
	ErstelltAm   time.Time `gorm:"column:erstelltAm"`
	GeändertAm   time.Time `gorm:"column:geaendertAm"`
	GelöschtAm   time.Time `gorm:"column:geloeschtAm"`
}

// func GeneriereHochbeet(vent *Wasserventil) *Hochbeet {

// 	return &Hochbeet{Name: "HB_1", Wasserventil: vent}
// }

func GeneriereVentil() *Wasserventil {
	return &Wasserventil{Name: "Zulauf_1", Durchflussmenge: 3}
}
