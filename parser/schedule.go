package parser

import (
	// "fmt"
	"strings"

	"golang.org/x/net/html"
)

type RaceWeekend struct {
	FP1 bool
	FP2 bool
	FP3 bool
	SprintShootout bool
	Sprint bool
	Qualifying bool
	Race bool
}

func NewRaceWeekend() RaceWeekend {
	rw := RaceWeekend{false, false, false, false, false, false, false}
	return rw
}

func ParseScheduleSource(n *html.Node) RaceWeekend {
	rw := NewRaceWeekend()
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "f1-race-hub--timetable-listings" {
					parseSchedule(n, &rw)
					return
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
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

// func parseDataCell(n *html.Node, s *string) {
// 	if n.Type == html.TextNode {
// 		data := strings.Replace(n.Data, "\n", "", -1)
// 		data = strings.Trim(data, " ")

// 		if data != "" {
// 			*s += data + " "
// 		}
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		parseDataCell(c, s)
// 	}
// }

// func parseRow(n *html.Node, r *string) {
// 	if n.Type == html.ElementNode && n.Data == "td" {
// 		cell := ""
// 		parseDataCell(n, &cell)
// 		cell = strings.Trim(cell, " ")
// 		*r += "," + cell
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		parseRow(c, r)
// 	}
// }

// func parseTable(n *html.Node, rows *[]string) {
// 	if n.Type == html.ElementNode && n.Data == "tr" {
// 		row := ""
// 		parseRow(n, &row)
// 		row = strings.Trim(row, " ,")
// 		if row != "" {
// 			*rows = append(*rows, row)
// 		}
// 	}

// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		parseTable(c, rows)
// 	}
// }
