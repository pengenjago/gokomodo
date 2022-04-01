package dto

type ProductRes struct {
	Id          string `json:"id"`
	ProductName string `json:"productName"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	SellerId    string `json:"sellerId"`
	SellerName  string `json:"sellerName"`
}

type ProductQuery struct {
	Search   string
	SellerId string
	PageNo   int
	PageSize int
}
