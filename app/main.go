package main

import (
	"net/http"

	"github.com/nazliander/word-count-with-golang/app/collector"
)

func main() {
	http.HandleFunc("/text-collect", collector.TextWordCountCollect)
	http.ListenAndServe(":7979", nil)
}
