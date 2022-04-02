package repository

import (
	"gokomodo/config"

	"gorm.io/gorm"
)

// Base Seller Repo
type SellerRepo struct {
}

// Model Seller
type Seller struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Email    string
	Name     string
	Password string
	Address  string
}

func GetSellerRepo() *SellerRepo {

	return &SellerRepo{}
}

// Get Seller by Email
func (o *SellerRepo) GetByEmail(email string) *Seller {
	var seller Seller

	err := config.DbConn().Model(Seller{}).Where("email = ?", email).Find(&seller).Error

	if err != nil || seller.ID == "" {
		return nil
	}

	return &seller
}
