package controller

import (
	"gokomodo/dto"
	"gokomodo/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
}

var productController *ProductController
var productSvc *service.ProductSvc

func InitProductController() {
	GET("/products", productController.findAll, Allow("Buyer", "Seller"))
	GET("/my-products", productController.myProduct, Allow("Seller"))
}

func (o *ProductController) findAll(c echo.Context) error {
	pageNo, _ := strconv.Atoi(c.QueryParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	data, total, err := productSvc.FindAll(&dto.ProductQuery{
		Search:   c.QueryParam("search"),
		PageNo:   pageNo,
		PageSize: pageSize,
	})

	if err == nil {
		return pagedRes(c, data, pageNo, pageSize, int(total))
	}

	return invalidRequest(c, err)
}

func (o *ProductController) myProduct(c echo.Context) error {

	pageNo, _ := strconv.Atoi(c.QueryParam("pageNo"))
	pageSize, _ := strconv.Atoi(c.QueryParam("pageSize"))

	data, total, err := productSvc.FindAll(&dto.ProductQuery{
		Search:   c.QueryParam("search"),
		PageNo:   pageNo,
		PageSize: pageSize,
		SellerId: c.Request().Header.Get("Uid"),
	})

	if err == nil {
		return pagedRes(c, data, pageNo, pageSize, int(total))
	}

	return invalidRequest(c, err)
}
