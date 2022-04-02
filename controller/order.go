package controller

import (
	"gokomodo/dto"
	"gokomodo/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
}

var orderController *OrderController
var orderSvc *service.OrderSvc

func InitOrderController() {
	GET("/my-orders", orderController.findAll, Allow(Buyer, Seller))
	POST("/order", orderController.create, Allow(Buyer))
	PUT("/confirm-order/:id", orderController.confirm, Allow(Seller))
}

func (o *OrderController) findAll(c echo.Context) error {
	pageNo, _ := strconv.Atoi(c.QueryParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	var query dto.OrderQuery
	query.PageNo = pageNo
	query.PageSize = pageSize

	userAuth := getUserAuth(c)
	if userAuth.Role == Seller {
		query.SellerId = userAuth.Id
	}
	if userAuth.Role == Buyer {
		query.BuyerId = userAuth.Id
	}

	data, total, err := orderSvc.FindAll(&query)

	if err == nil {
		return pagedRes(c, data, pageNo, pageSize, int(total))
	}

	return invalidRequest(c, err)
}

func (o *OrderController) create(c echo.Context) error {

	var data dto.OrderReq
	err := c.Bind(&data)

	if err == nil {
		res, err := orderSvc.Create(data, getUserAuth(c))

		if err == nil {
			return c.JSON(StatusOK, ResponseData{Status: StatusOK, Message: Success, Data: res})
		} else {
			return invalidRequest(c, err)
		}
	}

	return invalidRequest(c, err)
}

func (o *OrderController) confirm(c echo.Context) error {

	err := orderSvc.ConfirmOrder(c.Param("id"), getUserAuth(c))

	if err == nil {
		return c.JSON(StatusOK, Response{Status: StatusOK, Message: "Success Confirm Order"})
	}

	return invalidRequest(c, err)
}
