package configs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() {
	// Capture connection properties.
	cfg := mysql.Config{
		//User:   os.Getenv("DBUSER"),
		//Passwd: os.Getenv("DBPASS"),
		User:                 "root",
		Passwd:               "12345",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "futbol",
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
