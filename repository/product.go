package repository

import (
	"errors"
	"gokomodo/config"
	"gokomodo/dto"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
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

func (o *Product) ToResponse() dto.ProductRes {

	return dto.ProductRes{
		Id:          o.ID,
		ProductName: o.ProductName,
		Description: o.Description,
		Price:       o.Price,
		SellerName:  o.Seller.Name,
	}
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

func (o *ProductRepo) GetById(id string) Product {

	var data Product

	config.DbConn().Model(Product{}).Joins("Seller").Where(`"products".id = ?`, id).First(&data)

	return data
}

func (o *ProductRepo) Create(data *Product) error {

	err := config.DbConn().Create(data).Error

	if err != nil {

		log.Error("Error Create Product :", err.Error())
		err = errors.New("something went wrong")
	}

	return err
}
