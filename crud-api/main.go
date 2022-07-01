package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
	Year   int     `json:"year"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {
	// Init Router
	r := mux.NewRouter()
	// Mock Data
	books = append(books, Book{ID: "1", Isbn: "12345", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}, Year: 2022})
	books = append(books, Book{ID: "2", Isbn: "54321", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}, Year: 2021})

	// Route Handlers / Endpoints
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
