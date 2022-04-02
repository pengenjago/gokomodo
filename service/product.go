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
		products = append(products, val.ToResponse())
	}

	return products, total, err
}

func (o *ProductSvc) GetBydID(id string) dto.ProductRes {

	data := repository.GetProductRepo().GetById(id)

	return data.ToResponse()
}

func (o *ProductSvc) Create(param *dto.ProductReq, userAuth *dto.UserAuth) (dto.ProductRes, error) {
	var res dto.ProductRes

	data := repository.Product{
		ProductName: param.ProductName,
		Description: param.Description,
		Price:       param.Price,
		SellerId:    userAuth.Id,
	}

	err := repository.GetProductRepo().Create(&data)

	if err != nil {
		return res, err
	}

	return o.GetBydID(data.ID), err
}
