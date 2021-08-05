package Models

import (
	_ "fmt"

	"github.com/jinzhu/gorm"
)

type Hochbeet struct {
	gorm.Model
	Name                string                `sql:"size:30" gorm:"column:name;not null"`
	Wasserventile       []Wasserventil        `gorm:"column:wasserventil;not null;foreignKey:hochbeetId"`
	Feuchtigkeitssensor []Feuchtigkeitssensor `gorm:"column:feuchtigkeitssensor;not null;foreignKey:feuchtigkeitssensorId"`
}

// func GeneriereHochbeet(vent *Wasserventil) *Hochbeet {

// 	return &Hochbeet{Name: "HB_1", Wasserventil: vent}
// }

func GeneriereVentil() *Wasserventil {
	return &Wasserventil{Name: "Zulauf_1", Durchflussmenge: 3}
}
