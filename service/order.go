package service

import (
	"errors"
	"gokomodo/dto"
	"gokomodo/repository"
)

type OrderSvc struct {
}

func (o *OrderSvc) FindAll(query *dto.OrderQuery) ([]dto.OrderRes, int64, error) {

	data, total, err := repository.GetOrderRepo().FindAll(query)

	var dtos []dto.OrderRes
	for _, val := range data {
		dtos = append(dtos, val.ToResponse())
	}

	return dtos, total, err
}

const (
	Pending string = "Pending"

	Accepted string = "Accepted"
)

func (o *OrderSvc) GetBydID(id string) dto.OrderRes {

	data := repository.GetOrderRepo().GetById(id)

	return data.ToResponse()
}

func (o *OrderSvc) Create(param dto.OrderReq, userAuth *dto.UserAuth) (dto.OrderRes, error) {

	var res dto.OrderRes

	if param.ProductId == "" {
		return res, errors.New("product cannot be empty")
	}

	product := repository.GetProductRepo().GetById(param.ProductId)

	if product.ID == "" {
		return res, errors.New("product not found")
	}

	data := repository.Order{
		BuyerId:    userAuth.Id,
		SellerId:   product.SellerId,
		ProductId:  product.ID,
		Quantity:   param.Quantity,
		Status:     Pending,
		TotalPrice: param.Quantity * product.Price,
	}

	err := repository.GetOrderRepo().Create(&data)

	if err != nil {
		return res, err
	}

	return o.GetBydID(data.ID), err
}

func (o *OrderSvc) ConfirmOrder(orderId string, userAuth *dto.UserAuth) error {

	order := repository.GetOrderRepo().GetOrderSeller(orderId, userAuth.Id)

	if order.ID == "" {
		return errors.New("order not found")
	}

	if order.Status != Pending {
		return errors.New("only status pending can be confirm")
	}

	order.Status = Accepted

	err := repository.GetOrderRepo().Update(&order)

	return err
}
