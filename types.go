package main

type Item struct {
	Context  string `json:"@context"`
	ItemType string `json:"@type"`
	ItemName string `json:"name"`
	ItemPic  string `json:"image"`
	ItemDesc string `json:"description"`
	Brand    struct {
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
		Seller    struct {
			SellerType string `json:"@type"`
			SellerName string `json:"name"`
		}
	}
}

type User struct {
	Name string
	ItemsListed string
	Sales string
	Reviews string
	ProfileDesc string
}