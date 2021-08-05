package main

import (
	database "GartenBewaesserung/Database"
	_ "GartenBewaesserung/Models"
	"fmt"
)

func main() {
	fmt.Println("starten..")
	database.InitDatabase()

	//test
	//Models.GetFunksteckdosenListe()
}
