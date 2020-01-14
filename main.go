package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

type Item struct {
	ItemName     string  `json:"name"`
	ItemDesc     string  `json:"description"`
	ItemPrice    float64 `json:"price"`
	ShippingCost float64 `json:"priceCurrency"`
	ItemCond     string  `json:"itemCondition"`
}

func main() {

}

func fetch_all_items(searchTerm string) {
	c := colly.NewCollector(
		colly.AllowedDomains("mercari.com"),
		colly.Async(true),
	)

	c.OnHTML("div.Flex-ych44r-0.Space-cutht5-0.Container-sc-9aa7mx-0.hepKGV", func(e *colly.HTMLElement) {
		e.ForEach("div.Flex-ych44r-0.Space-cutht5-0.Container-sc-9aa7mx-0.withMetaInfo__FullContainer-sc-1j2k5ln-0.hyLExl > script:first-of-type", func(_ int, e *colly.HTMLElement) {
			dat := e.Text

			jsonData := dat[strings.Index(dat, "{") : len(dat)-1]
			i := &Item{}
			err := json.Unmarshal([]byte(jsonData), i)
			if err != nil {
				log.Fatal(err)
			}

			name := i.ItemName

			fmt.Printf("Product Name: %s", name)
		})
	})

	//test := "Nike"
	//test1 := urlBuilderQuery(test).String()
	c.Visit("https://mercari.com/search/?keyword=lego")

}

func urlBuilderQuery(searchTerm string) *url.URL {
	baseUrl, err := url.Parse("https://mercari.com/search/")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
	}

	q := baseUrl.Query()
	q.Set("keyword", searchTerm)
	baseUrl.RawQuery = q.Encode()

	return baseUrl
}
