package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepo struct {
}

type Order struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	BuyerId    string
	SellerId   string
	ProductId  string
	Quantity   int
	TotalPrice int
	Status     string
	Buyer      Buyer   `gorm:"foreignKey:BuyerId"`
	Seller     Seller  `gorm:"foreignKey:SellerId"`
	Product    Product `gorm:"foreignKey:ProductId"`
}

func (o *Order) BeforeCreate(db *gorm.DB) error {

	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	return nil
}
