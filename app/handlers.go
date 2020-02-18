package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type Item struct {
	Context  string
	ItemType string
	ItemName string
	ItemPic  string
	ItemDesc string

	brand struct {
		BrandType string
		BrandName string
	}

	Offers struct {
		OfferType string
		OfferURL  string
		Currency  string
		Price     string
		ItemCond  string
		ItemAva   string

		Seller struct {
			SellerType string
			SellerName string
		}
	}
}

// GET /
func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Welcome to Unofficial MercariAPI! For documentation, please consult [PLACEHOLDER]")
}

// GET /products/{keyword}
func GetProductsByKeyword(res http.ResponseWriter, req *http.Request) {

}

// GET /product/{id}
func GetProductById(res http.ResponseWriter, req *http.Request) {

}