package database

import (
	models "GartenBewaesserung/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Connection struct {
	Provider   string
	ConnString string
}

func InitDatabase() {
	conn := GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		log.Println("Database Connection konnte nicht hergestellt werden")
	} else {
		log.Println("Database Connection erfolgreich erstellt")
		AutoMigrateDB()
	}
}

func GetConnectionString() *Connection {
	// read file
	data, err := ioutil.ReadFile("./Config/DatabaseConnection.json")
	if err != nil {
		fmt.Print(err)
	}
	var conn Connection
	err = json.Unmarshal(data, &conn)
	if err != nil {
		fmt.Println("error:", err)
	}
	return &conn
}

func AutoMigrateDB() {
	conn := GetConnectionString()
	db, err := gorm.Open(conn.Provider, conn.ConnString)
	defer db.Close()
	if err != nil {
		fmt.Print(err)
	}
	log.Println("... erstelle Tabelle 'Funksteckdose")
	db.Debug().AutoMigrate(&models.Funksteckdose{})

	log.Println("... erstelle Tabelle 'Wasserventil")
	db.Debug().AutoMigrate(&models.Wasserventil{})

	log.Println("... erstelle Tabelle 'Feuchtigkeitssensor")
	db.Debug().AutoMigrate(&models.Feuchtigkeitssensor{})

	log.Println("... erstelle Tabelle 'Hochbeet")
	db.Debug().AutoMigrate(&models.Hochbeet{})
}
