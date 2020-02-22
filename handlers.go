package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"github.com/gorilla/mux"
	"github.com/gocolly/colly"
)

type Item struct {
	Context  string `json:"@context"`
	ItemType string	`json:"@type"`
	ItemName string `json:"name"`
	ItemPic  string	`json:"image"`
	ItemDesc string	`json:"description"`
	brand struct {
		BrandType string `json:"@type"`
		BrandName string `json:"name"`
	}
	Offers struct {
		OfferType string `json:"@type"`
		OfferURL  string `json:"url"`
		Currency  string `json:"priceCurrency"`
		Price     string `json:"price"`
		ItemCond  string `json:"itemCondition"`
		ItemAva   string `json:"availability"`
		Seller struct {
			SellerType string `json:"@type"`
			SellerName string `json:"name"`
		}
	}
}

// GET /
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Welcome to Unofficial MercariAPI! For documentation, please consult [PLACEHOLDER]")
}

// GET /products/{keyword}
func GetProductsByKeyword(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	keyword := vars["keyword"]
	
	c := colly.NewCollector(
		colly.Async(true),
	)

	c.OnHTML("div.Flex-ych44r-0.Space-cutht5-0.Container-sc-9aa7mx-0.hepKGV", func(e *colly.HTMLElement) {
		e.ForEach("div.Flex-ych44r-0.Space-cutht5-0.Container-sc-9aa7mx-0.withMetaInfo__FullContainer-sc-1j2k5ln-0.hyLExl", func(_ int, e *colly.HTMLElement) {
			data := e.ChildText("script")

			jsonData := data[strings.Index(data,"{"):len(data)]
			i := &Item{}
			err := json.Unmarshal([]byte(jsonData), i)
			if err != nil {
				log.Fatal(err)
			}

			
		})
	})


	url := urlBuilderQuery(keyword)
	c.Visit(url)	
}

// GET /product/{id}
func GetProductById(res http.ResponseWriter, req *http.Request) {

}

func urlBuilderQuery(searchTerm string) string {
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
