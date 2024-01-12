package data

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UserDTO struct {
	Username string
	Password string
}

type user struct {
	id        int
    username  string
    password  string
    createdAt time.Time
}

func GetUsers(db *sql.DB) ([]user) {
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	defer rows.Close()

	if err != nil {
		fmt.Printf(err.Error())
	}

	var users []user
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)

		if err != nil {
			fmt.Printf(err.Error())
		}

		users = append(users, u)
	}

	return users
}

func CreateUser(db *sql.DB, user UserDTO) int64 {
	createdAt := time.Now()

	// Inserts our data into the users table and returns with the result and a possible error.
	// The result contains information about the last inserted id (which was auto-generated for us) and the count of rows this query affected.
	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, user.Username, user.Password, createdAt)

	if err != nil {
		fmt.Printf(err.Error())
	}

	userID, err := result.LastInsertId()

	return userID
}
