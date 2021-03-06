package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func openLog() {

	file, err := os.OpenFile("logPibes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		panic(err.Error("Unable to open log file"))
	}
}

func dbConnection() {

	// Open up our database connection.
	// The database is called testDb
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error("Connection to the DB failed"))
	}
}

func dbClose(driver.Conn db) {

	if db == nil {
		panic(err.Error("BD is not connected!"))
	} else {
		defer db.Close()
	}
}
