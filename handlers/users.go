package handlers

import (
	"database/sql"
	"fmt"
	"go-web-tut/data"
	"net/http"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
 
var db *sql.DB = data.SetupDB()

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	user := data.GetUser(db, userID)
	fmt.Println("Endpoint hit: get user", user)
}

func DeleteUserHandler (w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	user := data.DeleteUser(db, userID)
	fmt.Println("Endpoint hit: delete user", user)
}

func GetUsersHandler (w http.ResponseWriter, r *http.Request) {
	users := data.GetUsers(db)
	fmt.Println("Endpoint hit: get users", users)
}

func CreateUserHandler (w http.ResponseWriter, r *http.Request) {
	var user data.UserDTO
	json.NewDecoder(r.Body).Decode(&user)

	users := data.CreateUser(db, user)
	fmt.Println("Endpoint hit: create user", users)
}