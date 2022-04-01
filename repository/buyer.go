package repository

import (
	"gokomodo/config"

	"gorm.io/gorm"
)

// Base struct Buyer Repo
type BuyerRepo struct {
}

// Model Buyer
type Buyer struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
	Address  string
}

func GetBuyerRepo() *BuyerRepo {

	return &BuyerRepo{}
}

// Get Buyer by Email
func (o *BuyerRepo) GetByEmail(email string) *Buyer {
	var buyer Buyer

	err := config.DbConn().Model(Buyer{}).Where("email = ?", email).Find(&buyer).Error

	if err != nil || buyer.ID == "" {
		return nil
	}

	return &buyer
}
