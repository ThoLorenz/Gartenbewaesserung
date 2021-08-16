package Models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Durchlaufsensor struct {
	ID         int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name       string `sql:"size:30" gorm:"column:name;not null"`
	Pumpe      Pumpe
	PumpeID    uint
	Liter      decimal.Decimal `sql:"type:decimal(3,5)" gorm:"column:liter"`
	ErstelltAm time.Time       `gorm:"column:erstelltAm"`
	GeändertAm time.Time       `gorm:"column:geaendertAm"`
	GelöschtAm time.Time       `gorm:"column:geloeschtAm"`
}
