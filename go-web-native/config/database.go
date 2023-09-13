package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var host = "10.199.2.23"
var port = 1433
var user = "sa"
var password = "b1e22f71."
var database = "TUTORIYAL"

var DB *sql.DB

func ConnectDB() {

	conn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		host, user, password, port, database)
	// conn := fmt.Sprintf("sqlserver://%s:%s@%s?port=%d&database=%s;",
	// 	user, password, host, port, database)

	// log.Println(conn)

	db, err := sql.Open("mssql", conn)

	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	} else {
		log.Println("Database connected!")
	}

	DB = db

}
