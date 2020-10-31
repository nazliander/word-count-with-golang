package collector

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TextWordCountCollect uses a POST method for collecting the requested book
// Count the words, return and store those into a MongoDB.
func TextWordCountCollect(w http.ResponseWriter, r *http.Request) {

	var tr TextRequest

	switch r.Method {

	case "POST":

		err := json.NewDecoder(r.Body).Decode(&tr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		textData := tr.wordCountURLText()

		textDataResponse, _ := json.Marshal(textData)

		w.Header().Set("Content-Type", "application/json")
		w.Write(textDataResponse)

	default:
		fmt.Fprintf(w, "Only POST method is supported.")
	}
}
