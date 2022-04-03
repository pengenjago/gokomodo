package controller

import (
	"bytes"
	"encoding/json"
	"gokomodo/dto"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	setupDSN()

	loginAs := "seller"
	resp := request(t, "POST", "/auth/login?as="+loginAs, "", dto.UserLogin{
		Email:    "clothing@gmail.com",
		Password: "Seller123!",
	})

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	assert.Equal(t, http.StatusOK, dataResp.Status)
	assert.Equal(t, "Success", dataResp.Message)

}

func request(t *testing.T, method, path, auth string, body interface{}) *httptest.ResponseRecorder {

	InitEcho()
	RouteController()

	var bodyByte []byte
	var err error
	if body != nil {
		bodyByte, err = json.Marshal(body)
		if err != nil {
			assert.FailNow(t, "failed to serialize request body: "+err.Error())
		}
	}
	req := httptest.NewRequest(method, path, bytes.NewBuffer(bodyByte))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	if auth != "" {
		req.Header.Add("Authorization", "Bearer "+auth)
	}

	rsp := httptest.NewRecorder()

	e.ServeHTTP(rsp, req)

	return rsp
}

func getToken(t *testing.T, email, password, loginAs string) string {
	resp := request(t, "POST", "/auth/login?as="+loginAs, "", dto.UserLogin{
		Email:    email,
		Password: password,
	})

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var dataResp ResponseData
	json.Unmarshal(bodyBytes, &dataResp)

	if dataResp.Status == 401 || dataResp.Status == 400 {
		assert.FailNow(t, dataResp.Message)
	}

	return dataResp.Data.(map[string]interface{})["accessToken"].(string)
}

// Setup DSN First
func setupDSN() {
	dsn := "user=postgres password=admin host=localhost dbname=t_gokomodo port=5433 sslmode=disable TimeZone=Asia/Jakarta"

	os.Setenv("app.db.dsn", dsn)
}
