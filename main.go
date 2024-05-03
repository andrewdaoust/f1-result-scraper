package main

import (
	"fmt"

	"github.com/andrewdaoust/f1-result-scraper/parser"
	"github.com/andrewdaoust/f1-result-scraper/result"
)

func main() {
	// URL to make the HTTP request to
	// url := "https://www.formula1.com/en/results.html/2024/races/1229/bahrain/race-result.html"
	// url := "https://www.formula1.com/en/results.html/2024/races/1229/bahrain/qualifying.html"
	// url := "https://www.formula1.com/en/results.html/2024/races/1230/saudi-arabia/race-result.html"
	// url := "https://www.formula1.com/en/results.html/2024/races/1231/australia/race-result.html"
	// url := "https://www.formula1.com/en/results.html/2024/races/1233/china/sprint-results.html"

	urlStub := "https://www.formula1.com/en/results.html/2024/races/1229/bahrain"

	raceURL := fmt.Sprintf("%v/race-result.html", urlStub)
	qualURL := fmt.Sprintf("%v/qualifying.html", urlStub)

	racePage := parser.ResultPage{URL: raceURL}
	qualPage := parser.ResultPage{URL: qualURL}

	raceSource, raceErr := racePage.GetSource()
	if raceErr != nil {
		panic("Error getting race results")
	}

	qualSource, qualErr := qualPage.GetSource()
	if qualErr != nil {
		panic("Error getting qualifying results")
	}

	parsedRaceResults := parser.ParseSource(raceSource)
	parsedQualResults := parser.ParseSource(qualSource)

	raceResults, err := result.ParseResult(parsedRaceResults)
	if err != nil {
		panic("Error parsing race results")
	}
	qualResults, err := result.ParseResult(parsedQualResults)
	if err != nil {
		panic("Error parsing race results")
	}

	fmt.Println("Qualifying:", qualResults)
	fmt.Println("Race:", raceResults)
}
