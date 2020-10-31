package collector

import (
	"io/ioutil"
	"net/http"

	"github.com/nazliander/word-count-with-golang/app/analytics"
)

// wordCountURLText uses a TextRequest object to attach
// a word count map on a TextData struct.
func (tr TextRequest) wordCountURLText() TextData {

	linkText := collectURLText(tr.URL)
	wordCounts := analytics.WordCount(linkText)

	return TextData{
		URL:        tr.URL,
		WordCounts: wordCounts,
	}
}

// collectURLText uses a link object to get text information.
func collectURLText(link string) (linkText string) {

	resp, err := http.Get(link)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	linkText = string(responseBytes)

	return
}
