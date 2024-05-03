package parser

import (
	"strings"
	"net/http"

	"golang.org/x/net/html"
)

type Parser interface {
	GetSource() (*html.Node, error)
}

type ResultPage struct {
	URL string
}

func (rp ResultPage) GetSource() (*html.Node, error) {
	resp, err := http.Get(rp.URL)
	if err != nil {
		return nil, err
	}

	// Use the html package to parse the response body from the request
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func ParseSource(n *html.Node) []string {
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
	return rows
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
