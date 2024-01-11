package data

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	username string
	password string
}

func createUser(db *sql.DB, user User) (int64) {
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, user.username, user.password, createdAt)

	if err != nil {
		fmt.Printf(err.Error())
	}

	userID, err := result.LastInsertId()

	return userID
}

func SetupDB() {
	fmt.Println("Connecting to db")
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/GoWebDev")

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

	createUser(db, User{"Dzon", "1234"})
}
