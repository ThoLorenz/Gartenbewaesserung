package Models

import (
	_ "fmt"
	"time"
)

type Pumpe struct {
	ID               int    `gorm:"column:id;primaryKey;autoIncrement"`
	Name             string `gorm:"column:name;unique;not null;type:varchar(40)"`
	Funksteckdose    Funksteckdose
	FunksteckdoseID  int       `gorm:"column:funksteckdoseId"`
	Status           int       `gorm:"column:status"`
	IstPumpeFuerBeet bool      `gorm:"column:istPumpeFuerBeet"`
	ErstelltAm       time.Time `gorm:"column:erstelltAm"`
	GeändertAm       time.Time `gorm:"column:geaendertAm"`
	GelöschtAm       time.Time `gorm:"column:geloeschtAm"`
}
