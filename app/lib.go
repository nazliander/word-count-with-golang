package main

import (
	"fmt"
	"gopkg.in/jdkato/prose.v2"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func collectBook(bookLink string) (bookText string) {
	resp, err := http.Get(bookLink)
	// handle the error if there is one
	if err != nil {
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	bookBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bookText = string(bookBytes)
	// show the HTML code as a string %s
	return bookText
}

func wordCount(text string) (counts map[string]int) {
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	lowerText := strings.ToLower(text)

	doc, err := prose.NewDocument(lowerText)
	if err != nil {
		panic(err)
	}
	var tokenArray []string
	for _, tok := range doc.Tokens() {
		if IsLetter(tok.Text) {
			tokenArray = append(tokenArray, tok.Text)
		}
	}

	counts = make(map[string]int, len(tokenArray))
	for _, word := range tokenArray {
		counts[word]++
	}
	return counts
}

func storeBookText(text string, path string) {
	file, _ := os.Create(path)
	l, _ := file.WriteString(text)
	fmt.Println(l, "File written successfully")
	_ = file.Close()
}
