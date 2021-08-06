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
		AutoMigrateDB(db, err)
	}
}

func AutoMigrateDB(db *gorm.DB, err error) {
	if err != nil {
		fmt.Print(err)
	}
	db.Debug().AutoMigrate(&models.Funksteckdose{})
	db.Debug().AutoMigrate(&models.Wasserventil{})
	db.Debug().AutoMigrate(&models.Feuchtigkeitssensor{})
	db.Debug().AutoMigrate(&models.Hochbeet{})
	fmt.Println("... AutoMigration beendet")
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
