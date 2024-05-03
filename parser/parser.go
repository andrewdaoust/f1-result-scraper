package parser

import (
	"net/http"

	"golang.org/x/net/html"
)

type Parser interface {
	GetSource() (*html.Node, error)
}

type Page struct {
	URL string
}

func (p Page) GetSource() (*html.Node, error) {
	resp, err := http.Get(p.URL)
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
