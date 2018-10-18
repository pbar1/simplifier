package thesaurus

import (
	"bufio"
	"log"
	"net/http"
	"strings"
)

// ThesaurusResult defines the output a Thesaurus search
type ThesaurusResult struct {
	PartOfSpeech string
	Category     string
	Word         string
}

// Thesaurus returns a list of synonyms in JSON format using the
// Big Huge Thesaurus API
func Thesaurus(apiKey string, word string) []ThesaurusResult {
	url := "https://words.bighugelabs.com/api/2/" + apiKey + "/" + word + "/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	var results []ThesaurusResult
	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		result := ThesaurusResult{
			PartOfSpeech: line[0],
			Category:     line[1],
			Word:         line[2],
		}
		results = append(results, result)
	}

	return results
}
