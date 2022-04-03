package controller

import (
	"encoding/json"
	"gokomodo/dto"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateOrderAsBuyer(t *testing.T) {
	setupDSN()

	token := getToken(t, "budiman@gmail.com", "Buyer123!", Buyer)
	resp := request(t, "POST", "/order", token, []dto.OrderReq{
		{
			ProductId: "400ecf9e-e3e2-4e69-904c-b1b494c6105a",
			Quantity:  1,
		},
		{
			ProductId: "44a6c61a-e820-4eea-bfea-13adffcbdad8",
			Quantity:  1,
		},
	})

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)
}

func TestGetMyOrderAsBuyer(t *testing.T) {
	setupDSN()

	token := getToken(t, "budiman@gmail.com", "Buyer123!", Buyer)
	resp := request(t, "GET", "/my-orders", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)
}

func TestGetMyOrderAsSeller(t *testing.T) {
	setupDSN()

	token := getToken(t, "clothing@gmail.com", "Seller123!", Seller)
	resp := request(t, "GET", "/my-orders", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)
}

func TestConfirmOrderStatusPendingAsSeller(t *testing.T) {
	setupDSN()

	token := getToken(t, "deb@gmail.com", "Seller123!", Seller)
	resp := request(t, "PUT", "/confirm-order/6ceea581-94f0-4130-8435-250756224355", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success Confirm Order", dataResp.Message)
}

func TestConfirmOrderStatusAcceptedAsSeller(t *testing.T) {
	setupDSN()

	token := getToken(t, "deb@gmail.com", "Seller123!", Seller)
	resp := request(t, "PUT", "/confirm-order/bbb757d7-1bb6-46b7-b4f8-625f4d22966d", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusBadRequest, dataResp.Status)
	assert.Equal(t, "order already confirm", dataResp.Message)
}

func TestConfirmOrderAsBuyer(t *testing.T) {
	setupDSN()

	token := getToken(t, "budiman@gmail.com", "Buyer123!", Buyer)
	resp := request(t, "PUT", "/confirm-order/bbb757d7-1bb6-46b7-b4f8-625f4d22966d", token, nil)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusUnauthorized, dataResp.Status)
	assert.Equal(t, "access is not allowed", dataResp.Message)
}
