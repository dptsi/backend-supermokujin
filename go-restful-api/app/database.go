package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

func DotEnv(key string) string {

	err := godotenv.Load(".env") // load from env file

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func NewDB() *sql.DB {
	host := DotEnv("SQLSRV_HOST")
	port, _ := strconv.Atoi(DotEnv("SQLSRV_PORT"))
	user := DotEnv("SQLSRV_USERNAME")
	password := DotEnv("SQLSRV_PASSWORD")
	database := DotEnv("SQLSRV_DATABASE")

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

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
