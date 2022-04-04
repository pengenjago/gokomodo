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
	GET("/products", productController.findAll, Allow(Buyer, Seller))
	GET("/my-products", productController.myProduct, Allow(Seller))
	POST("/product", productController.create, Allow(Seller))
}

// Findall godoc
// @Tags Products
// @Summary Get Products
// @Description Get All Products
// @ID get-all-products
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param pageNo query string optional "Page No"
// @Param pageSize query string optional "Page Size"
// @Param search query string optional "Search Product"
// @Success 200 {object} ResponsePaged{data=[]dto.ProductRes}
// @Router /products [get]
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

// Findall godoc
// @Tags Products
// @Summary Get My Products
// @Description Get My Products as Seller
// @ID get-my-products
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param pageNo query string optional "Page No"
// @Param pageSize query string optional "Page Size"
// @Param search query string optional "Search Product"
// @Success 200 {object} ResponsePaged{data=[]dto.ProductRes}
// @Router /my-products [get]
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

// Create godoc
// @Tags Products
// @Summary Create Product
// @Description Create Product by Seller
// @ID create-product
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authorization"
// @Param data body dto.ProductReq true "request"
// @Success 200 {object} ResponseData{data=[]dto.ProductRes}
// @Router /product [post]
func (o *ProductController) create(c echo.Context) error {

	var data dto.ProductReq
	err := c.Bind(&data)

	if err == nil {
		res, err := productSvc.Create(&data, getUserAuth(c))

		if err == nil {
			return c.JSON(StatusOK, ResponseData{Status: StatusOK, Message: Success, Data: res})
		} else {
			return invalidRequest(c, err)
		}
	}

	return invalidRequest(c, err)
}
