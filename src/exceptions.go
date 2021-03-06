package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
import "log"
import "os"
import "fmt"

func main() {

	file, err := os.OpenFile("logPibes.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    	if err != nil {
		log.Fatal(err)
		panic(err.Error("Unable to open log file"))
    	}

        fmt.Println("Hello world")
	log.Println("Action registered into log")

	log.SetOutput(file)
	log.Println("Action sucessfully executed")

    // Open up our database connection.
    // The database is called testDb
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

    if err != nil {
        panic(err.Error("Connection to the DB failed"))
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

}
