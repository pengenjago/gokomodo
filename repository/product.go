package repository

import (
	"gokomodo/config"
	"gokomodo/dto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepo struct {
}

type Product struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	ProductName string
	Description string
	Price       int
	SellerId    string
	Seller      Seller `gorm:"foreignKey:SellerId"`
}

func GetProductRepo() *ProductRepo {

	return &ProductRepo{}
}

func (o *Product) BeforeCreate(db *gorm.DB) error {

	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	return nil
}

func (o *ProductRepo) FindAll(query *dto.ProductQuery) ([]Product, int64, error) {
	var data []Product

	tx := config.DbConn().Joins("Seller")

	if query.Search != "" {
		search := "%" + query.Search + "%"
		tx = tx.Where("product_name ilike ?", search)
	}

	if query.SellerId != "" {
		tx = tx.Where(`"products".seller_id = ?`, query.SellerId)
	}

	total := tx.Find(&data).RowsAffected

	if query.PageSize > 0 {
		tx = tx.Offset((query.PageNo - 1) * query.PageSize).Limit(query.PageSize)
	}

	err := tx.Find(&data).Error

	return data, total, err
}
