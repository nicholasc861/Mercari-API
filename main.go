package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"strings"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"

)

func main() {
	fetch_all_items()
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

func fetch_all_items() {
	c := colly.NewCollector(
	//colly.Async(true),
	)

	c.OnHTML("div.Flex-ych44r-0.Space-cutht5-0.Container-sc-9aa7mx-0.hepKGV", func(e *colly.HTMLElement) {
		e.ForEach("div.Flex-ych44r-0.Space-cutht5-0.Container-sc-9aa7mx-0.withMetaInfo__FullContainer-sc-1j2k5ln-0.hyLExl", func(_ int, e *colly.HTMLElement) {
			dat := e.ChildText("script")

			jsonData := dat[strings.Index(dat, "{"):len(dat)]
			i := &Item{}
			err := json.Unmarshal([]byte(jsonData), i)

			if err != nil {
				log.Fatal(err)
			}

			name := i.ItemDesc

			fmt.Printf("Product Name: %s \n", name)
		})
	})

	//test := "Nike"
	//test1 := urlBuilderQuery(test).String()
	c.Visit("https://www.mercari.com/search/")

}

func urlBuilderQuery(searchTerm string) *url.URL {
	baseUrl, err := url.Parse("https://www.mercari.com/search/")
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
	}

	q := baseUrl.Query()
	q.Set("keyword", searchTerm)
	baseUrl.RawQuery = q.Encode()

	return baseUrl
}
