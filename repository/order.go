package repository

import (
	"errors"
	"gokomodo/config"
	"gokomodo/dto"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
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

func GetOrderRepo() *OrderRepo {

	return &OrderRepo{}
}

func (o *Order) BeforeCreate(db *gorm.DB) error {

	if o.ID == "" {
		o.ID = uuid.New().String()
	}

	return nil
}

func (o *Order) ToResponse() dto.OrderRes {

	return dto.OrderRes{
		Id:                 o.ID,
		BuyerName:          o.Buyer.Name,
		SellerName:         o.Seller.Name,
		DestinationAddress: o.Buyer.Address,
		PickUpAddress:      o.Seller.Address,
		ProductName:        o.Product.ProductName,
		Quantity:           o.Quantity,
		Price:              o.Product.Price,
		TotalPrice:         o.TotalPrice,
		Status:             o.Status,
	}
}

func (o *OrderRepo) GetPreQuery() *gorm.DB {

	return config.DbConn().Joins("Buyer").Joins("Seller").Joins("Product")
}

func (o *OrderRepo) GetById(id string) Order {

	var data Order

	o.GetPreQuery().Where(`"orders".id = ?`, id).First(&data)

	return data
}

func (o *OrderRepo) GetOrderSeller(orderId, sellerId string) Order {

	var data Order

	o.GetPreQuery().Where(`"orders".id = ? and "Seller".id = ?`, orderId, sellerId).First(&data)

	return data
}

func (o *OrderRepo) FindAll(query *dto.OrderQuery) ([]Order, int64, error) {

	var data []Order

	tx := o.GetPreQuery()

	if query.SellerId != "" {
		tx = tx.Where(`"Seller".id = ?`, query.SellerId)
	}

	if query.BuyerId != "" {
		tx = tx.Where(`"Buyer".id = ?`, query.BuyerId)
	}

	total := tx.Find(&data).RowsAffected

	if query.PageSize > 0 {
		tx = tx.Offset((query.PageNo - 1) * query.PageSize).Limit(query.PageSize)
	}

	err := tx.Find(&data).Error

	return data, total, err
}

func (o *OrderRepo) Create(data *Order) error {

	err := config.DbConn().Create(data).Error

	if err != nil {

		log.Error("Error Create Order :", err.Error())
		err = errors.New("something went wrong")
	}

	return err
}

func (o *OrderRepo) Update(data *Order) error {

	err := config.DbConn().Save(data).Error

	if err != nil {
		log.Error("Error Update Order :", err.Error())
		err = errors.New("something went wrong")
	}

	return err
}
