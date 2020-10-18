package main

import (
	"bufio"
	"fmt"
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
	defer resp.Body.Close()

	bookBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	bookText = string(bookBytes)
	// show the HTML code as a string %s
	return bookText
}

func wordCount(text string) (counts map[string]int) {
	var isWord = regexp.MustCompile(`\w*`)
	var alphaChars = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	var stopWords = retrieveStopwords("app/stopwords/english_stopwords.txt")

	tokenArray := isWord.FindAllString(strings.ToLower(text), -1)

	counts = make(map[string]int, len(tokenArray))
	for _, word := range tokenArray {
		if alphaChars(word) && notIn(stopWords, word) {
			counts[word]++
		}
	}
	return counts
}

func storeBookText(text string, path string) {
	file, _ := os.Create(path)
	l, _ := file.WriteString(text)
	fmt.Println(l, "File written successfully")
	_ = file.Close()
}

func retrieveStopwords(path string) (stopwords []string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		stopwords = append(stopwords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}

func notIn(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return false
		}
	}
	return true
}
