package main

import (
	"log"
	"os"
	"net/http"
	"database/sql"

	"github.com/lib/pq"
	"github.com/subosito/gotenv"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book
var db *sql.DB

func init(){
	gotenv.Load()
}

func logFatal(err error){
	if err != nil{
		log.Fatal(err)
	}
}

func main() {

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	log.Println(pgUrl)

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
}

func getBook(w http.ResponseWriter, r *http.Request) {
}

func addBook(w http.ResponseWriter, r *http.Request) {
}

func updateBook(w http.ResponseWriter, r *http.Request) {
}

func removeBook(w http.ResponseWriter, r *http.Request) {

}
