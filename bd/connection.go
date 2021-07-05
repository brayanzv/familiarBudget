package bd

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var server = "DESKTOP-HRICUD2"
var port = 1433
var user = "sa"
var password = "dev"
var database = "bd"

func ConectionBD() *gorm.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)
	db, err := gorm.Open("mssql", connectionString)

	if err != nil {
		log.Fatal("Failed to create connection pool. Error: " + err.Error())
	}
	gorm.DefaultCallback.Create().Remove("mssql:set_identity_insert")

	if err != nil {
		log.Fatal(err.Error())
		return db
	}
	log.Println("Connection db")
	return db
}


