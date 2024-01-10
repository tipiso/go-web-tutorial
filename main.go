package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Book")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Book")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create Book")
}

func RegisterBooksRouter(r *mux.Router) {
	// Specific path restriction
	bookrouter := r.PathPrefix("/books").Subrouter()
	bookrouter.HandleFunc("/", AllBooks).Methods("GET")
	bookrouter.HandleFunc("/{title}", GetBook).Methods("GET")

	// Example route handlers
	bookrouter.HandleFunc("/{title}", CreateBook).Methods("POST")
	// r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	// r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	// r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	bookrouter.HandleFunc("/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
}

func main() {
	// Configure the database connection (always check errors)
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

	query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

	_, err = db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connected")
	// r := mux.NewRouter()
	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// RegisterBooksRouter(r)

	// r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hoy, Info from server!!!")
	// })

	// http.ListenAndServe(":80", r)
}
