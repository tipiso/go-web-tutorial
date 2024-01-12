package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *sql.DB {
	fmt.Println("Connecting to db")
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/GoWebDev?parseTime=true")

	if err != nil {
		fmt.Printf(err.Error())
	}
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var queryArr [1]string

	queryArr[0] = `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	for i := 0; i < len(queryArr); i++ {
		_, err = db.Exec(queryArr[i])
	}

	if err != nil {
		log.Fatal(err)
	}

	log.Print("Connected")

	return db
}
