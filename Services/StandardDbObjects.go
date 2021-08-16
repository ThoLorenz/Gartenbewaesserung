package Services

import (
	database "GartenBewaesserung/Database"
	models "GartenBewaesserung/Models"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

func CreateStandardObjectsForDb() {
	conn := database.GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
		return
	} else {
		defer db.Close()

		funksteckdose_1_1 := models.Funksteckdose{
			Name:       "Funk_1",
			Kennung:    "B",
			Status:     0,
			Systemcode: "010110",
			DipCode:    "10100",
			Pulslaenge: 300,
			ErstelltAm: time.Now()}

		funksteckdose_1_2 := models.Funksteckdose{
			Name:       "Funk_2",
			Kennung:    "A",
			Status:     0,
			Systemcode: "01110",
			DipCode:    "11000",
			Pulslaenge: 300,
			ErstelltAm: time.Now()}

		funksteckdose_1_3 := models.Funksteckdose{
			Name:       "Funk_3",
			Kennung:    "B",
			Status:     0,
			Systemcode: "10011",
			ErstelltAm: time.Now(),
			DipCode:    "01011",
			Pulslaenge: 300}

		wasserventil_1_1 := models.Wasserventil{
			Name:            "Ventil_1",
			Status:          0,
			Dauer:           5,
			Durchflussmenge: 2,
			Funksteckdose:   funksteckdose_1_1,
			ErstelltAm:      time.Now()}

		wasserventil_1_2 := models.Wasserventil{
			Name:            "Ventil_2",
			Status:          0,
			Dauer:           5,
			Durchflussmenge: 2,
			Funksteckdose:   funksteckdose_1_2,
			ErstelltAm:      time.Now()}

		hochbeet_1 := models.Hochbeet{
			Name:       "Hochbeet_1",
			ErstelltAm: time.Now()}

		feuchtsensor_1_1 := models.Feuchtigkeitssensor{
			Name:            "Petersilie_Sensor",
			MinFeuchtigkeit: 40,
			MaxFeuchtigkeit: 60,
			Temperatur:      30,
			Wasserventil:    wasserventil_1_1,
			ErstelltAm:      time.Now()}

		feuchtsensor_1_2 := models.Feuchtigkeitssensor{
			Name:            "Broccoli_Sensor",
			MinFeuchtigkeit: 40,
			MaxFeuchtigkeit: 60,
			Wasserventil:    wasserventil_1_2,
			Temperatur:      30,
			ErstelltAm:      time.Now()}

		pumpe_1 := models.Pumpe{
			Name:             "Pumpe_1",
			Status:           0,
			IstPumpeFuerBeet: true,
			Funksteckdose:    funksteckdose_1_3,
			ErstelltAm:       time.Now(),
		}

		durchlaufsensor_1 := models.Durchlaufsensor{
			Name:       "Durchlaufsensor_!",
			Liter:      decimal.NewFromInt(3.0),
			Pumpe:      pumpe_1,
			ErstelltAm: time.Now()}

		fmt.Println("... Datenbank leeren ..")
		fmt.Println("... neue Standardobjekte erstellen")
		if err := db.Save(&durchlaufsensor_1).Error; err != nil {
			fmt.Println("********** Fehler beim DB speichern ", durchlaufsensor_1.Name)
			return
		}

		if err := db.Save(&hochbeet_1).Error; err != nil {
			fmt.Println("********** Fehler beim DB speichern ", hochbeet_1.Name)
			return
		}
		feuchtsensor_1_1.HochbeetID = uint(hochbeet_1.ID)
		if err := db.Save(&feuchtsensor_1_1).Error; err != nil {
			fmt.Println("********** Fehler beim DB speichern ", feuchtsensor_1_1.Name)
			return
		}
		feuchtsensor_1_2.HochbeetID = uint(hochbeet_1.ID)
		if err := db.Save(&feuchtsensor_1_2).Error; err != nil {
			fmt.Println("********** Fehler beim DB speichern ", feuchtsensor_1_2.Name)
			return
		}

		hochbeet_1.ListSensoren = []models.Feuchtigkeitssensor{feuchtsensor_1_1, feuchtsensor_1_2}

		models.DeleteSingleSensor(hochbeet_1, 2)

	}
}
