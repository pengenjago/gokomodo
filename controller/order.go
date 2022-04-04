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

// Findall godoc
// @Tags Orders
// @Summary Get My Orders
// @Description Get My Orders
// @ID get-my-orders
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param pageNo query string optional "Page No"
// @Param pageSize query string optional "Page Size"
// @Success 200 {object} ResponsePaged{data=[]dto.OrderRes}
// @Router /my-orders [get]
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

// Create godoc
// @Tags Orders
// @Summary Create Orders
// @Description Create Orders by Buyer
// @ID create-orders
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param data body dto.OrderReq true "request"
// @Success 200 {object} ResponseData{data=[]dto.OrderRes}
// @Router /order [post]
func (o *OrderController) create(c echo.Context) error {

	var data []dto.OrderReq
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

// Update godoc
// @Tags Orders
// @Summary Cofirm Orders
// @Description Confirm Orders by Seller
// @ID confirm-orders
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param data body dto.OrderReq true "request"
// @Param id path string true "id orders"
// @Success 200 {object} Response{}
// @Router /confirm-order/{id} [put]
func (o *OrderController) confirm(c echo.Context) error {

	err := orderSvc.ConfirmOrder(c.Param("id"), getUserAuth(c))

	if err == nil {
		return c.JSON(StatusOK, Response{Status: StatusOK, Message: "Success Confirm Order"})
	}

	return invalidRequest(c, err)
}
