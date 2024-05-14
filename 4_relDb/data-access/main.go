package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var userName string
var password string

func main() {
	fmt.Print("Enter your username: ")
	fmt.Scan(&userName)

	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	//connection properties
	cfg := mysql.Config{
		User:                 userName,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "recordings",
		AllowNativePasswords: true, //THIS IS IMPORTANT
	}
	// fmt.Println("DBUSER:", os.Getenv("DBUSER"))
	// fmt.Println("DBPASS:", os.Getenv("DBPASS"))
	// fmt.Println("CFG " + cfg.FormatDSN())

	//db handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("connected <3")
}
