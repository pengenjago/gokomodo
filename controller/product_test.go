package controller

import (
	"encoding/json"
	"gokomodo/dto"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	setupDSN()

	token := getToken(t, "budiman@gmail.com", "Buyer123!", Buyer)
	resp := request(t, "GET", "/products", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)

}

func TestGetMyProductAsBuyer(t *testing.T) {
	setupDSN()

	token := getToken(t, "budiman@gmail.com", "Buyer123!", Buyer)
	resp := request(t, "GET", "/my-products", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusUnauthorized, dataResp.Status)
	assert.Equal(t, "access is not allowed", dataResp.Message)

}

func TestGetMyProductAsSeller(t *testing.T) {
	setupDSN()

	token := getToken(t, "clothing@gmail.com", "Seller123!", Seller)
	resp := request(t, "GET", "/my-products", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)
}

func TestCreateProductAsSeller(t *testing.T) {
	setupDSN()

	token := getToken(t, "clothing@gmail.com", "Seller123!", Seller)
	resp := request(t, "POST", "/product", token, dto.ProductReq{
		ProductName: "Erigo Beach",
		Description: "Kaos Erigo Beach",
		Price:       80000,
	})

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)
}
