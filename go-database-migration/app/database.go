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
	db, err := sql.Open("mssql", conn)

	// conn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s;",
	// 	user, password, host, port, database)
	// db, err := sql.Open("sqlserver", conn)

	// log.Println(conn)

	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third

	// up semua versi
	// migrate -database "sqlserver://tutorial:b1e22f71..@10.199.2.23:1433?database=TUTORIAL_MIG;" -path db/migrations up
	// up 1 versi
	// migrate -database "sqlserver://tutorial:b1e22f71..@10.199.2.23:1433?database=TUTORIAL_MIG;" -path db/migrations up 1
	// up 1 versi lagi
	// migrate -database "sqlserver://tutorial:b1e22f71..@10.199.2.23:1433?database=TUTORIAL_MIG;" -path db/migrations up 1
	// down semua versi
	// migrate -database "sqlserver://tutorial:b1e22f71..@10.199.2.23:1433?database=TUTORIAL_MIG;" -path db/migrations down
	// down 1 versi dari yang terakhir
	// migrate -database "sqlserver://tutorial:b1e22f71..@10.199.2.23:1433?database=TUTORIAL_MIG;" -path db/migrations down 1

	// membersihkan dirty migration ke versi sebelumnya (versi yang works)
	// migrate -database "sqlserver://tutorial:b1e22f71..@10.199.2.23:1433?database=TUTORIAL_MIG;" -path db/migrations force <versi yang works>
}
