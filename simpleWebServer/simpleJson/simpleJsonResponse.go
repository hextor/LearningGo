package main

import (
	"log"
	"net/http"
	"encoding/json"
)

type Book struct { // simple struct
	Title string
	Author string
	Genre string
	Year int `json:"Released"` // changes name of Key
}

// linter doesn't like me
var testBooks = []Book{
	{Title: "my little vampire", Author: "sassy lassy", Genre: "young adult", Year: 2012}, // notice we still use original Key name
	{Title: "big dreamer", Author: "billy the kiddo", Genre: "inspirational", Year: 2018},
	{Title: "a lot to do about nothing", Author: "darwin jokavish", Genre: "non fiction comedy", Year: 2014}}

func showBooks(w http.ResponseWriter, r *http.Request) {
	// transforms data into a JSON encoding, forces UTF-8, escapes characters for html ex. '>' becomes '\u003c'
	data, err:= json.Marshal(testBooks) 

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s, err")
	}

	w.Header().Set("Content-Type", "application/json") // sets a response header, needed for json
	w.Write(data) // writes response
}

func main() {
	http.HandleFunc("/getBook", showBooks) // set the router
	log.Fatal(http.ListenAndServe(":8080", nil)) // set listening port
}
