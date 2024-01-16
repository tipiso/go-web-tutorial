package main

import (
	"database/sql"
	"go-web-tut/data"
	"go-web-tut/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func RegisterUsersRouter(r *mux.Router, db *sql.DB) {
	// Specific path restriction
	usersrouter := r.PathPrefix("/users").Subrouter()
	usersrouter.HandleFunc("/{userID}", handlers.GetUserHandler)
	usersrouter.HandleFunc("/delete/{userID}", handlers.DeleteUserHandler)
	usersrouter.HandleFunc("", handlers.GetUsersHandler)
	usersrouter.HandleFunc("/create", handlers.CreateUserHandler)
}

func main() {
	db := data.SetupDB()

	r := mux.NewRouter()
	
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	RegisterUsersRouter(r, db)

	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hoy, Info from server!!!")
	// })

	http.ListenAndServe(":80", r)
}
