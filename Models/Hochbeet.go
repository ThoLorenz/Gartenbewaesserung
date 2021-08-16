package Models

import (
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

func Create(name string) *Hochbeet {
	return &Hochbeet{Name: name, ErstelltAm: time.Now()}
}

// func AddSingleSensor(beet Hochbeet, sensor Feuchtigkeitssensor) *Hochbeet {
// 	beet.ListSensoren = append(beet.ListSensoren, sensor)
// 	return &beet
// }

// https://stackoverflow.com/a/18203895
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func DeleteSingleSensorById(beet Hochbeet, SensorId int) {
	index := SliceIndex(len(beet.ListSensoren), func(i int) bool { return beet.ListSensoren[i].ID == SensorId })
	// Remove the element at index i from a.
	beet.ListSensoren[index] = beet.ListSensoren[len(beet.ListSensoren)-1] // Copy last element to index i.
	beet.ListSensoren[len(beet.ListSensoren)-1] = Feuchtigkeitssensor{}    // Erase last element (write zero value).
	beet.ListSensoren = beet.ListSensoren[:len(beet.ListSensoren)-1]       // Truncate slice.
}

// func GeneriereHochbeet(vent *Wasserventil) *Hochbeet {

// 	return &Hochbeet{Name: "HB_1", Wasserventil: vent}
// }

func GeneriereVentil() *Wasserventil {
	return &Wasserventil{Name: "Zulauf_1", Durchflussmenge: 3}
}
