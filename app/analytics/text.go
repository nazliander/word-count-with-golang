package analytics

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// WordCount given a string counts the words and returns to
// a word-count map.
func WordCount(text string) (counts map[string]int) {
	var isWord = regexp.MustCompile(`\w*`)
	var alphaChars = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	var stopWords = retrieveStopwords("analytics/stopwords/english_stopwords.txt")

	tokenArray := isWord.FindAllString(strings.ToLower(text), -1)

	counts = make(map[string]int, len(tokenArray))
	for _, word := range tokenArray {
		if alphaChars(word) && notIn(stopWords, word) {
			counts[word]++
		}
	}
	return counts
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
