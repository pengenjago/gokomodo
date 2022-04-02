package service

import (
	"errors"
	"fmt"
	"gokomodo/dto"
	"gokomodo/repository"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var sign []byte = []byte("TGOKOMODO2022")

// Base Auth Svc
type AuthSvc struct {
}

// Login
func (o *AuthSvc) Login(param dto.UserLogin) (*dto.LoginRes, error) {

	var token string
	var err error

	if strings.EqualFold(param.LoginAs, "Buyer") {
		buyer := repository.GetBuyerRepo().GetByEmail(param.Email)

		if buyer == nil {
			return nil, errors.New("username or password is incorrect")
		}

		checkPass := checkPassword(buyer.Password, param.Password)
		if checkPass != nil {
			return nil, errors.New("username or password is incorrect")
		}

		token, err = generateToken(buyer.ID, buyer.Name, buyer.Email, "Buyer")
	}

	if strings.EqualFold(param.LoginAs, "Seller") {
		seller := repository.GetSellerRepo().GetByEmail(param.Email)

		if seller == nil {
			return nil, errors.New("username or password is incorrect")
		}

		checkPass := checkPassword(seller.Password, param.Password)
		if checkPass != nil {
			return nil, errors.New("username or password is incorrect")
		}

		token, err = generateToken(seller.ID, seller.Name, seller.Email, "Seller")
	}

	if err == nil {
		return &dto.LoginRes{AccessToken: token}, nil
	}

	return nil, err
}

func (o *AuthSvc) ValidateToken(accessToken string) (*dto.UserAuth, error) {

	var ret *dto.UserAuth
	parser := new(jwt.Parser)

	token, err := parser.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return sign, nil
	})

	if err == nil {
		mapClaim := token.Claims.(jwt.MapClaims)

		ret = &dto.UserAuth{
			Role:  mapClaim["role"].(string),
			Id:    mapClaim["uid"].(string),
			Email: mapClaim["email"].(string),
		}
	} else {
		return nil, errors.New("invalid access token")
	}

	return ret, err
}

// Compare Hash Password and Parameter Password
func checkPassword(hashPassword, paramPassword string) error {
	check := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(paramPassword))

	return check
}

// Generate JWT
func generateToken(id, name, email, role string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   id,
		"name":  name,
		"email": email,
		"role":  role,
		"iat":   time.Now().Unix(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(sign)
}
