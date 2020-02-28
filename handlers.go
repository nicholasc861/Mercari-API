package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
)

// GET /
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Welcome to the Unofficial MercariAPI! For documentation, please consult [PLACEHOLDER]")
}

// GET /products/{keyword}
func GetProductsByKeyword(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	keyword := vars["keyword"]
	items := []Item{}

	c := colly.NewCollector()

	c.OnHTML(".kXmgUV", func(e *colly.HTMLElement) {
		data := e.ChildText("script")
		jsonData := data[strings.Index(data, "{"):len(data)]
		tempItem := &Item{}
		
		err := json.Unmarshal([]byte(jsonData), tempItem)
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, *tempItem)
	})

	url := UrlBuilderQuery(keyword)
	c.Visit(url)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(items); err != nil {
		panic(err)
	}
}

// GET /product/{id}
func GetProductById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	c := colly.NewCollector()
	tempItem := &Item{}

	c.OnHTML(".kVrcCF", func(e *colly.HTMLElement) {
		tempItem.ItemName = e.ChildText("")
		
	})

	urlBuilt := "https://www.mercari.com/us/item" + id
	c.Visit(urlBuilt)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(tempItem); err != nil {
		panic(err)
	}
}

// GET /user/{id}
func GetUserById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	c := colly.NewCollector()
	tempUser := &User{}

	c.OnHTML(".hMDyjy", func(e *colly.HTMLElement) {
		tempUser.Name = e.ChildText("")
		tempUser.ItemsListed = e.ChildText("")
		tempUser.Reviews = e.ChildText("")
		tempUser.Sales = e.ChildText("")

	})
	
	urlBuilt := "https://www.mercari.com/u/" + id
	c.Visit(urlBuilt)

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(res).Encode(tempUser); err != nil {
		panic(err)
	}
}