package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env") // load from env file

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ConnectDB() {

	var host = goDotEnvVariable("SQLSRV_HOST")
	var port = goDotEnvVariable("SQLSRV_PORT")
	var user = goDotEnvVariable("SQLSRV_USERNAME")
	var password = goDotEnvVariable("SQLSRV_PASSWORD")
	var database = goDotEnvVariable("SQLSRV_DATABASE")

	sqlsrv_port, _ := strconv.Atoi(port) // convert string port to int port

	conn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		host, user, password, sqlsrv_port, database)
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
