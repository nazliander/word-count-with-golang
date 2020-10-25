package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type BookData struct {
	URL        string         `json:"URL"`
	WordCounts map[string]int `json:"BookWordCounts"`
}

type BookRequest struct {
	URL string `json:"URL"`
}

func bookAnalytics(w http.ResponseWriter, r *http.Request) {

	var br BookRequest

	switch r.Method {

	case "POST":

		err := json.NewDecoder(r.Body).Decode(&br)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		bookURL := br.URL

		bookText := collectBook(bookURL)
		bookWordCounts := wordCount(bookText)

		bookData := BookData{
			URL:        bookURL,
			WordCounts: bookWordCounts,
		}

		bookDataResponse, _ := json.Marshal(bookData)

		w.Header().Set("Content-Type", "application/json")
		w.Write(bookDataResponse)

	default:
		fmt.Fprintf(w, "Only POST method is supported.")
	}
}

func main() {
	http.HandleFunc("/book-collect", bookAnalytics)
	http.ListenAndServe(":7979", nil)
}
