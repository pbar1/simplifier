package thesaurus

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gomodule/redigo/redis"
)

// SearchResult is the structured response of a Big Huge Thesaurus search
type SearchResult struct {
	PartOfSpeech string
	Category     string
	Word         string
}

// SearchText queries Big Huge Thesaurus and returns a raw string
func SearchText(apiKey string, word string) string {
	url := "https://words.bighugelabs.com/api/2/" + apiKey + "/" + word + "/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return bodyString
}

// SearchTextCached queries Big Huge Thesaurus and returns a raw string, using a redis cache
func SearchTextCached(apiKey string, word string) string {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bodyString, err := redis.String(conn.Do("GET", word))
	if err != nil && err.Error() == "redigo: nil returned" {
		bodyString = SearchText(apiKey, word)
		conn.Do("SET", word, bodyString)
	} else if err != nil {
		log.Fatalln(err)
	}

	return bodyString
}

// Search queries Big Huge Thesaurus and returns a SearchResult slice
func Search(apiKey string, word string) []SearchResult {
	resultsString := SearchTextCached(apiKey, word)
	scanner := bufio.NewScanner(strings.NewReader(resultsString))
	var results []SearchResult

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		result := SearchResult{
			PartOfSpeech: line[0],
			Category:     line[1],
			Word:         line[2],
		}
		results = append(results, result)
	}

	return results
}
