module web-tut/main

go 1.19

require (
	github.com/gorilla/mux v1.8.1
	go-web-tut/data v0.0.0-00010101000000-000000000000
)

require github.com/go-sql-driver/mysql v1.7.1

replace go-web-tut/data => ./data
