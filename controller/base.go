package controller

import (
	"fmt"
	"gokomodo/config"
	"gokomodo/dto"
	"gokomodo/repository"
	"gokomodo/util"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"
	echoSwagger "github.com/swaggo/echo-swagger"
)

var e *echo.Echo

const (
	StatusOK int = http.StatusOK

	StatusUnauthorized int = http.StatusUnauthorized

	StatusBadRequest int = http.StatusBadRequest

	Success string = "Success"

	InvalidReq string = "Invalid Request"

	// Role Seller
	Seller string = "Seller"

	// Role Buyer
	Buyer string = "Buyer"
)

// Start
func Start() {

	InitEcho()

	RouteController()

	Migrate()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", util.GetIntProperties("app.server.port"))))
}

func InitEcho() {

	e = echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{echo.OPTIONS, echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowedHeaders: []string{echo.HeaderContentType, echo.HeaderOrigin, echo.HeaderAccessControlAllowHeaders, echo.HeaderAccessControlAllowOrigin, echo.HeaderAccessControlAllowMethods},
		Debug:          true,
	})

	e.Use(middleware.Logger())
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.Use(AuthMiddleware)
}

// RouteController
func RouteController() {
	// Routes
	e.GET("/", index)

	InitAuthController()
	InitProductController()
	InitOrderController()
}

// Generate Table
func Migrate() {
	conn := config.DbConn()

	if conn != nil {

		conn.AutoMigrate(&repository.Buyer{})
		conn.AutoMigrate(&repository.Order{})
		conn.AutoMigrate(&repository.Seller{})
		conn.AutoMigrate(&repository.Product{})
	}
}

//AuthMiddleware is
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		authString := c.Request().Header.Get("Authorization")

		var err error
		if authString == "" {
			c.Request().Header.Set("AccessToken", "")
			c.Request().Header.Set("Access-Control-Allow-Origin", "*")
		} else {
			if strings.HasPrefix(authString, "Bearer ") {

				authString = authString[7:]

				token, errToken := authSvc.ValidateToken(authString)

				if errToken == nil {
					c.Request().Header.Set("Uid", token.Id)
					c.Request().Header.Set("Email", token.Email)
					c.Request().Header.Set("Role", token.Role)
					c.Request().Header.Set("AccessToken", authString)
					c.Request().Header.Set("Access-Control-Allow-Origin", "*")
				} else {
					return c.JSON(StatusBadRequest, Response{
						Status:  StatusBadRequest,
						Message: errToken.Error(),
					})
				}

			} else {
				return c.JSON(StatusBadRequest, Response{
					Status:  StatusBadRequest,
					Message: "not a bearer token authorization",
				})
			}
		}

		if err == nil {
			return next(c)
		}

		return c.JSON(StatusBadRequest, Response{
			Status:  StatusBadRequest,
			Message: InvalidReq,
		})
	}
}

//Allow Role
func Allow(roleList ...string) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {

		return func(c echo.Context) error {

			allowed := false

			role := c.Request().Header.Get("Role")

			if role != "" {

				for _, val := range roleList {
					if val == role {
						allowed = true
						break
					}
				}
			}

			if allowed {

				return next(c)
			}

			return c.JSON(StatusUnauthorized, Response{Status: StatusUnauthorized, Message: "access is not allowed"})
		}
	}
}

// Get User Auth
func getUserAuth(c echo.Context) *dto.UserAuth {
	return &dto.UserAuth{
		Id:    c.Request().Header.Get("Uid"),
		Role:  c.Request().Header.Get("Role"),
		Email: c.Request().Header.Get("Email"),
	}
}

func index(c echo.Context) error {
	return c.JSON(http.StatusOK, time.Now().Format("02/01/2006 15:04"))
}

//GET is
func GET(pattern string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {

	e.GET(pattern, handler, middlewares...)
}

//POST is
func POST(pattern string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {

	e.POST(pattern, handler, middlewares...)
}

//PUT is
func PUT(pattern string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {

	e.PUT(pattern, handler, middlewares...)
}

//DELETE is
func DELETE(pattern string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {

	e.DELETE(pattern, handler, middlewares...)
}

func invalidRequest(c echo.Context, err error) error {

	if err != nil {
		return c.JSON(StatusBadRequest, Response{Status: StatusBadRequest, Message: err.Error()})
	}

	return c.JSON(StatusBadRequest, Response{Status: StatusBadRequest, Message: InvalidReq})
}

func pagedRes(c echo.Context, data interface{}, pageNo int, pageSize int, totalRecord int) error {

	totalPage := 0

	if pageSize > 0 {

		totalPage = (totalRecord / pageSize)

		if totalRecord%pageSize > 0 {
			totalPage++
		}
	}

	if pageNo > totalPage {

		pageNo = totalPage
	}

	if data == nil {

		return c.JSON(StatusOK, ResponsePaged{Status: StatusOK, Message: Success, PageNo: pageNo, PageSize: pageSize, TotalRecord: totalRecord, PageTotal: totalPage})
	} else {

		return c.JSON(StatusOK, ResponsePaged{Status: StatusOK, Message: Success, PageNo: pageNo, PageSize: pageSize, TotalRecord: totalRecord, PageTotal: totalPage, Data: data})
	}
}

//Response is
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

//ResponseData is
type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//ResponsePaged is
type ResponsePaged struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	PageNo      int         `json:"pageNo"`
	PageSize    int         `json:"pageSize"`
	PageTotal   int         `json:"pageTotal"`
	TotalRecord int         `json:"totalRecord"`
}
