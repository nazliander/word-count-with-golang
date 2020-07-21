package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	// Sherlock book link
	searchString string = "https://www.gutenberg.org/files/1661/1661-0.txt"
	// Path to store Sherlock Book
	pathToStoreBook string = "./book.txt"
)

func main() {
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Please visit: http://localhost:7979/book")
		bookText := collectBook(searchString)
		storeBookText(bookText, pathToStoreBook)
		bookWordCounts := wordCount(bookText)
		jsonCounts, _ := json.Marshal(bookWordCounts)
		fmt.Fprintln(w, string(jsonCounts))
	})

	http.ListenAndServe(":7979", nil)
}
