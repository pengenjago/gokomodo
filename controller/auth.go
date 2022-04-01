package controller

import (
	"gokomodo/dto"
	"gokomodo/service"

	"github.com/labstack/echo/v4"
)

// Base Auth Controller
type AuthController struct {
}

var authController *AuthController
var authSvc *service.AuthSvc

func InitAuthController() {
	POST("/auth/login", authController.login)
}

func (o *AuthController) login(c echo.Context) error {
	var data dto.UserLogin
	err := c.Bind(&data)

	data.LoginAs = c.QueryParam("as")

	if err == nil {
		data, err := authSvc.Login(data)

		if err == nil {
			return c.JSON(StatusOK, ResponseData{Status: StatusOK, Message: Success, Data: data})
		} else {
			return c.JSON(StatusUnauthorized, Response{Status: StatusUnauthorized, Message: err.Error()})
		}
	}

	return invalidRequest(c, err)
}
