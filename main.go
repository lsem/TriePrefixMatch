package main

import (
	"./tptnmatch"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func FileNameForURL(fileURL string) string {
	u, err := url.Parse(fileURL)
	if err != nil {
		panic(err)
	}
	return url.PathEscape(u.Path)
}

func GetCacheFileName(fileURL string) string {
	return fmt.Sprintf("/tmp/TextPatternMatch/%s", FileNameForURL(fileURL))
}

func DownloadText(textFileURL string) ([]byte, error) {
	// Failure, download text
	resp, err := http.Get(textFileURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func TryReadFromCache(URI string) (string, error) {
	cacheData, err := ioutil.ReadFile(GetCacheFileName(URI))
	if err != nil {
		return "", err
	}
	return string(cacheData), nil
}

func WriteIntoCache(URI string, data []byte) error {
	// Try to put into cache
	if err := os.MkdirAll(filepath.Dir(GetCacheFileName(URI)), 0744); err != nil {
		return err
	}
	if err := ioutil.WriteFile(GetCacheFileName(URI), data, 0644); err != nil {
		return err
	}
	return nil
}

func GetTextByURI(URI string) (string, error) {
	fromCache, err := TryReadFromCache(URI)
	if err == nil {
		return fromCache, nil
	}

	body, err := DownloadText(URI)
	if err != nil {
		fmt.Printf("ERROR: download failed. Error: %v\n", err)
		return "", err
	}

	// (try to) Put into cache
	if err := WriteIntoCache(URI, body); err != nil {
		fmt.Printf("WARNING: Failed writing into cache. Error: %v\n", err)
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

func main() {
	text, err := GetTextByURI("https://www.gutenberg.org/cache/epub/1112/pg1112.txt")
	if err != nil {
		fmt.Errorf("ERROR: Failed downloading sample text, check connection. Error: %v", err)
		return
	}
	matchCount := 0
	tptnmatch.MatchTextAgainstTrie(text, tptnmatch.BuildTrie(GetPatterns()), func(s string) {
		matchCount++
	})
	if matchCount != 14670 {
		fmt.Errorf("Match Count Expected 14760, got: %d\n", matchCount)
		return
	}

	fmt.Println("Passed")
}
