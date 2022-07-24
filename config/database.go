package database

import (
	"database/sql"
	"os"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() sql.DB {
	db_name := os.Getenv("DB_NAME")
	if db_name == "" {
		db_name = "golang"
	}
	db_host := os.Getenv("DB_HOST")
	if db_host == "" {
		db_host = "127.0.0.1"
	}
	db_port := os.Getenv("DB_PORT")
	if db_port == "" {
		db_port = "3306"
	}
	db_user := os.Getenv("DB_USER")
	if db_user == "" {
		db_user = "root"
	}
	db_password := os.Getenv("DB_PASSWORD")
	conn_string := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_password, db_host, db_port, db_name)
	db, err := sql.Open("mysql", conn_string)
	checkErr(err)
	return *db
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Panic WOY")
		panic(err)
	}
}
