package database

import (
	models "GartenBewaesserung/Models"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDatabase() {
	db, err := gorm.Open("mysql",
		"snooker147admin:admin2020lorenz@tcp(127.0.0.1:3306)/gartenbewaesserung?charset=utf8&parseTime=True")
	db.SingularTable(true)
	//defer db.Close()
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		log.Println("Database Connection erfolgreich erstellt")

		log.Println("... erstelle Tabelle 'Funksteckdose")
		db.Debug().AutoMigrate(&models.Funksteckdose{})

		log.Println("... erstelle Tabelle 'Wasserventil")
		db.Debug().AutoMigrate(&models.Wasserventil{})

		log.Println("... erstelle Tabelle 'Feuchtigkeitssensor")
		db.Debug().AutoMigrate(&models.Feuchtigkeitssensor{})

		log.Println("... erstelle Tabelle 'Hochbeet")
		db.Debug().AutoMigrate(&models.Hochbeet{})
	}
}
