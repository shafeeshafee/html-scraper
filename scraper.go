package main

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func fetchDocument(url string) (*goquery.Document, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return goquery.NewDocumentFromReader(response.Body)
}
