package parser

import (
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type Results map[string]int

func ParseResultSource(n *html.Node) (Results, error) {
	rows := []string{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Data == "table" {
			for _, a := range n.Attr {
				if a.Key == "class" && a.Val == "resultsarchive-table" {
					parseTable(n, &rows)
					return
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)
	return makeResults(rows)
}

func parseDataCell(n *html.Node, s *string) {
	if n.Type == html.TextNode {
		data := strings.Replace(n.Data, "\n", "", -1)
		data = strings.Trim(data, " ")

		if data != "" {
			*s += data + " "
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseDataCell(c, s)
	}
}

func parseRow(n *html.Node, r *string) {
	if n.Type == html.ElementNode && n.Data == "td" {
		cell := ""
		parseDataCell(n, &cell)
		cell = strings.Trim(cell, " ")
		*r += "," + cell
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseRow(c, r)
	}
}

func parseTable(n *html.Node, rows *[]string) {
	if n.Type == html.ElementNode && n.Data == "tr" {
		row := ""
		parseRow(n, &row)
		row = strings.Trim(row, " ,")
		if row != "" {
			*rows = append(*rows, row)
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parseTable(c, rows)
	}
}

func makeResults(rows []string) (Results, error) {
	results := Results{}
	for _, row := range rows {
		splitRow := strings.Split(row, ",")
		
		var position int
		var err error
		if splitRow[0] == "NC" {
			position = 0
		} else {
			position, err = strconv.Atoi(splitRow[0])
			if err != nil {
				return nil, err
			}
		}
		name := splitRow[2]
		shortName := strings.Split(name, " ")[2]

		results[shortName] = position
	}

	return results, nil
}