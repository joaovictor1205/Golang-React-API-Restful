package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {

	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang Pointers", Author: "Mr. Pointers", Year: "2010"},
		Book{ID: 2, Title: "Go Routines", Author: "Mr. Routines", Year: "2011"},
		Book{ID: 3, Title: "Golang Routers", Author: "Mr. Routers", Year: "2012"},
		Book{ID: 4, Title: "Golang Concurrency", Author: "Mr. Concurrency", Year: "2013"},
		Book{ID: 5, Title: "Golang Paterns", Author: "Mr. Paterns", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stringID, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatal(err)
	}

	for _, book := range books {
		if book.ID == stringID {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	books = append(books,book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Update book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove an book")
}
