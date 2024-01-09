package main

import (
	"fmt"
	"net/http"

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
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	RegisterBooksRouter(r)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hoy, Info from server!!!")
	})

	http.ListenAndServe(":80", r)
}
