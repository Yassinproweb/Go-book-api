package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var books []Book

type Book struct {
	Title  string  `json:"title"`
	ISBN   string  `json:"isbn"`
	ID     string  `json:"id"`
	Author *Author `json:"author"`
}

type Author struct {
	Name  string `json:"name"`
	Field string `json:"field"`
}

func getbooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func addbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updatebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deletebook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	rtr := mux.NewRouter()

	books = append(books, Book{Title: "Fundamentals Of Golang", ISBN: "547662001", ID: "go001", Author: &Author{Name: "Katungi Yassin", Field: "Web Dev"}})
	books = append(books, Book{Title: "Fundamentals Of Elixir", ISBN: "747662001", ID: "ex002", Author: &Author{Name: "Katungi Yassin", Field: "Concurrency"}})
	books = append(books, Book{Title: "Fundamentals Of Haskell", ISBN: "647662009", ID: "hs003", Author: &Author{Name: "Katungi Yassin", Field: "Beginners"}})

	rtr.HandleFunc("/books", getbooks).Methods("GET")
	rtr.HandleFunc("/books/{id}", getbook).Methods("GET")
	rtr.HandleFunc("/books", addbook).Methods("POST")
	rtr.HandleFunc("/books/{id}", updatebook).Methods("PUT")
	rtr.HandleFunc("/books/{id}", deletebook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4007", rtr))
}
