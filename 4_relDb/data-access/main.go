package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

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

	//call getAlbumsByArtist
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)
}

// albumsByArtist query
func albumsByArtist(name string) ([]Album, error) {
	var albums []Album //albums slice

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name) //get artist's rows
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		//Scan takes a list of pointers to Go values, where the column values will be written.
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	//After the loop, check for an error from the overall query, using rows.Err.
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}
