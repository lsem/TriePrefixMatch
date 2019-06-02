package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
)

func GetSampleText() (string, error) {
	fileName := "https://www.gutenberg.org/cache/epub/1112/pg1112.txt"

	var body []byte

	// Try load text from FS cache
	cacheFileName := fmt.Sprintf("/tmp/TriePrefixMatch/%s", url.PathEscape(fileName))
	cacheData, err := ioutil.ReadFile(cacheFileName)
	if err != nil {
		// Failure, download text
		resp, err := http.Get(fileName)
		if err != nil {
			return "", err
		}

		defer resp.Body.Close()
		body, err = ioutil.ReadAll(resp.Body)

		if err != nil {
			return "", err
		}

		// Try to put into cache
		if err := os.MkdirAll(filepath.Dir(cacheFileName), 0744); err != nil {
			fmt.Printf("WARNING: cannot crate directory for cache. Error: %v\n", err)
		} else {
			if err := ioutil.WriteFile(cacheFileName, body, 0644); err != nil {
				fmt.Printf("WARNING: failed writing downloaded text data into cache. Error: %v\n", err)
			}
		}
	} else {
		// Success, take cached text
		body = cacheData
	}

	return string(body), nil
}

func GetPatterns() []string {
	return []string{
		"the", "be", "to", "of", "and", "a", "in", "that", "have", "I", "it", "for", "not", "on", "with", "he", "as",
		"you", "do", "at", "this", "but", "his", "by", "from", "they", "we", "say", "her", "she", "or", "an", "will",
		"my", "one", "all", "would", "there", "their", "what", "so", "up", "out", "if", "about", "who", "get", "which",
		"go", "me", "when", "make", "can", "like", "time", "no", "just", "him", "know", "take", "people", "into",
		"year", "your", "good", "some", "could", "them", "see", "other", "than", "then", "now", "look", "only", "come",
		"its", "over", "think", "also", "back", "after", "use", "two", "how", "our", "work", "first", "well", "way",
		"even", "new", "want", "because", "any", "these", "give", "day", "most", "us"}
}

func TestMatchingMostPopularWordsInRomeoAndJuliet(t *testing.T) {
	text, err := GetSampleText()
	if err != nil {
		_ = fmt.Errorf("ERROR: Failed downloading sample text, check connection. Error: %v", err)
		os.Exit(1)
	}
	matchCount := 0
	TrieMatching(text, BuildTrie(GetPatterns()), func(s string) {
		matchCount++
	})
	if matchCount != 14670 {
		t.Errorf("Match Count Expected 14760, got: %d\n", matchCount)
	}
}
