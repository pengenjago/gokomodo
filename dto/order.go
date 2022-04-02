package dto

type OrderReq struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type OrderRes struct {
	Id                 string `json:"id"`
	BuyerName          string `json:"buyerName"`
	SellerName         string `json:"sellerName"`
	PickUpAddress      string `json:"pickupAddress"`
	DestinationAddress string `json:"destinationAddress"`
	ProductName        string `json:"productName"`
	Quantity           int    `json:"quantity"`
	Price              int    `json:"price"`
	TotalPrice         int    `json:"totalPrice"`
	Status             string `json:"status"`
}

type OrderQuery struct {
	SellerId string
	BuyerId  string
	PageNo   int
	PageSize int
}
