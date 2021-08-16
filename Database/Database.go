package database

import (
	models "GartenBewaesserung/Models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Connection struct {
	Provider   string
	ConnString string
}

func InitDatabase() {
	conn := GetConnectionString()
	dsn := conn.ConnString
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,

		//	NamingStrategy:  schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
		DefaultStringSize:        256,  // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		//	DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		//DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		//	SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

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
	//db.SingularTable(true)
	db.Migrator().DropTable(&models.Hochbeet{})
	db.Migrator().DropTable(&models.Feuchtigkeitssensor{})
	db.Migrator().DropTable(&models.Wasserventil{})
	db.Migrator().DropTable(&models.Durchlaufsensor{})
	db.Migrator().DropTable(&models.Pumpe{})
	db.Migrator().DropTable(&models.Funksteckdose{})
	fmt.Println("... DropTable beendet")

	db.AutoMigrate(&models.Funksteckdose{})
	db.AutoMigrate(&models.Durchlaufsensor{})
	db.AutoMigrate(&models.Wasserventil{})

	db.AutoMigrate(&models.Feuchtigkeitssensor{})
	db.AutoMigrate(&models.Hochbeet{})
	db.AutoMigrate(&models.Pumpe{})
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
