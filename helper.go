package main

import (
	"net/url"
	"fmt"
)

func UrlBuilderQuery(searchTerm string) string {
	baseUrl, err := url.Parse("https://www.mercari.com/search/")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
	}

	q := baseUrl.Query()
	q.Set("keyword", searchTerm)
	baseUrl.RawQuery = q.Encode()
	urlAsString := baseUrl.String()
	return urlAsString
}


