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
}

// GET /product/{id}
func GetProductById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	c := colly.NewCollector()

	c.OnHTML(".kVrcCF", func(e *colly.HTMLElement) {
		tempItem := &Item{}
		tempItem.ItemName = e.ChildText("")
		
	})

	c.Visit("https://www.mercari.com/us/item/" + id)
}

// GET /user/{id}
func GetUserById(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	c := colly.NewCollector()

	c.OnHTML(".hMDyjy", func(e *colly.HTMLElement) {
		tempUser := &User{}
		tempUser.Name = e.ChildText("")
		tempUser.ItemsListed = e.ChildText("")
		tempUser.Reviews = e.ChildText("")
		tempUser.Sales = e.ChildText("")

	})


	c.Visit("https://www.mercari.com/u/" + id)
}