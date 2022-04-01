package service

import (
	"gokomodo/dto"
	"gokomodo/repository"
)

type ProductSvc struct {
}

func (o *ProductSvc) FindAll(param *dto.ProductQuery) ([]dto.ProductRes, int64, error) {
	data, total, err := repository.GetProductRepo().FindAll(param)

	var products []dto.ProductRes
	for _, val := range data {
		product := dto.ProductRes{
			Id:          val.ID,
			ProductName: val.ProductName,
			Description: val.Description,
			Price:       val.Price,
			SellerId:    val.Seller.ID,
			SellerName:  val.Seller.Name,
		}

		products = append(products, product)
	}

	return products, total, err
}
