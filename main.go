package main

import (
	"fmt"
	"strings"

	"github.com/andrewdaoust/f1-result-scraper/parser"
	// "github.com/andrewdaoust/f1-result-scraper/result"
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

	getResult(raceURL)
	getResult(qualURL)

	scheduleUrl := "https://www.formula1.com/en/racing/2024/Miami.html"
	rw := getSchedule(scheduleUrl)
	fmt.Println(rw)

	schedule2Url := "https://www.formula1.com/en/racing/2024/Japan.html"
	rw2 := getSchedule(schedule2Url)
	fmt.Println(rw2)
}

func getResult(url string) {
	page := parser.Page{URL: url}

	source, err := page.GetSource()
	if err != nil {
		panic("Error getting results source")
	}

	result, err := parser.ParseResultSource(source)
	// result, err := result.ParseResult(parsedSource)
	if err != nil {
		panic("Error parsing results")
	}

	fmt.Println(result)
}

func getSchedule(url string) parser.RaceWeekend {
	page := parser.Page{URL: url}

	source, err := page.GetSource()
	if err != nil {
		panic("Error getting results source")
	}

	urlSegments := strings.Split(url, "/")
	location := strings.Replace(urlSegments[len(urlSegments)-1], ".html", "", -1)
	year := urlSegments[len(urlSegments)-2]

	rw := parser.ParseScheduleSource(source, location, year)
	return rw
}