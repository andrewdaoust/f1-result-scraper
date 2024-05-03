package parser

import (
	"strings"

	"golang.org/x/net/html"
)

type RaceWeekend struct {
	Location string
	RaceName string
	Circuit string
	FP1 bool
	FP2 bool
	FP3 bool
	SprintShootout bool
	Sprint bool
	Qualifying bool
	Race bool
}

func NewRaceWeekend() RaceWeekend {
	rw := RaceWeekend{"", "", "", false, false, false, false, false, false, false}
	return rw
}

func ParseScheduleSource(n *html.Node, location, year string) RaceWeekend {
	rw := NewRaceWeekend()
	rw.Location = location
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "f1-race-hub--timetable-listings" {
					parseSchedule(n, &rw)
				}
			}
		} else if n.Data == "h2" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "f1--s" {
					parseRaceName(n, &rw, year)
				}
			}
		} else if n.Data == "p" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "f1-uppercase misc--tag no-margin" {
					parseCircuit(n, &rw)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	rw.RaceName = strings.Replace(rw.RaceName, year, "", 1)
	return rw
}

func parseSchedule(n *html.Node, rw *RaceWeekend) {
	for _, a := range n.Attr {
		if a.Key == "class" && strings.Contains(a.Val, "row js-") { //
			switch a.Val {
			case "row js-practice-1":
				rw.FP1 = true
			case "row js-practice-2":
				rw.FP2 = true
			case "row js-practice-3":
				rw.FP3 = true
			case "row js-sprint-shootout":
				rw.SprintShootout = true
			case "row js-sprint":
				rw.Sprint = true
			case "row js-qualifying":
				rw.Qualifying = true
			case "row js-race":
				rw.Race = true
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseSchedule(c, rw)
	}
}

func parseRaceName(n *html.Node, rw *RaceWeekend, year string) {
	if n.Type == html.TextNode {
		name := cleanRaceName(n.Data, year)
		// fmt.Println(name)
		rw.RaceName = name
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseRaceName(c, rw, year)
	}
}

func cleanRaceName(name, year string) string {
	newName := strings.ToUpper(name)
	newName = strings.Replace(newName, "FORMULA 1", "", -1)
	newName = strings.Replace(newName, year, "", 1)
	newName = strings.Replace(newName, "  ", " ", -1)
	newName = strings.TrimSpace(newName)

	return newName
}

func parseCircuit(n *html.Node, rw *RaceWeekend) {
	if n.Type == html.TextNode {
		rw.Circuit = n.Data
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseCircuit(c, rw)
	}
}