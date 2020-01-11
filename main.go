package main

import (
	"fmt"
	"net/url"
	"github.com/gocolly/colly"
)

type Item struct {
	ItemName     string
	ItemDesc     string
	ItemPrice    float64
	ShippingCost float64
	ItemCond     string
}

func main() {
	
}

func fetch_all_items(searchTerm string){
	c := colly.NewCollector(
		colly.AllowedDomains("mercari.com"),
		colly.Async(true),
	)

	c.OnHTML("div.Flex-ych44r-0 Space-cutht5-0 Container-sc-9aa7mx-0 Grid2__CellWrapper-mpt2p4-1 fKtviO", func(e *colly.HTMLElement) {
		e.ForEach("div.Flex__Box-ych44r-1 Grid2__Col-mpt2p4-0 kXKTPb",  func(_ int, e *colly.HTMLElement) {
			i := Item{}
			i.ItemName = e.ChildText("a.Text__LinkText-sc-1e98qiv-0-a Link__StyledAnchor-dkjuk2-0 htCDzf Link__StyledPlainLink-dkjuk2-2 hykQP")
			fmt.Printf("Product Name: %s", i.ItemName)
		})
	})
	test := "Nike"
	test1 := urlBuilderQuery(test).String()
	c.Visit(test1)

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
