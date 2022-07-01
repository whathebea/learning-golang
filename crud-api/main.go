package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
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

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through books and find one with the id from the params
	for index, book := range books {
		if book.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through books and find one with the id from the params
	for index, book := range books {
		// If the id from the params matches the id in the book, update the book
		if book.ID == params["id"] {
			var bookUpdate Book
			_ = json.NewDecoder(r.Body).Decode(&bookUpdate)
			books = append(books[:index], books[index+1:]...)
			books = append(books, bookUpdate)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Get a single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through books and find one with the id from the params
	for _, book := range books {
		// If the id from the params matches the id in the book, return the book
		if book.ID == params["id"] {
			// Return the book
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	// Add ID
	book.ID = strconv.Itoa(rand.Intn(10000000))
	// Add book to books array
	books = append(books, book)
	// Return the book
	json.NewEncoder(w).Encode(book)
}

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

	// Start Server
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
