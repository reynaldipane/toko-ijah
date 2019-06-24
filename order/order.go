package order

import (
	"time"
)

//Order is the definition of purchase table in database
type Order struct {
	ID             uint       `gorm:"primary_key" json:"id" valid:"-"`
	NumberSold     int        `gorm:"not null" json:"number_sold" valid:"numeric,required"`
	SellPrice      int        `gorm:"not null" json:"sell_price" valid:"numeric,required"`
	TotalPrice     int        `gorm:"not null" json:"total_price" valid:"numeric"`
	Notes          string     `gorm:"type:varchar(100)" json:"receipt_number" valid:"-"`
	ProductID      uint       `json:"product_id" valid:"numeric,required" sql:"type:uint REFERENCES products(id)"`
	ProductStockID uint       `json:"product_stock_id" valid:"numeric,required" sql:"type:uint REFERENCES product_stocks(id)"`
	CreatedAt      time.Time  `json:"created_at" valid:"-"`
	UpdatedAt      time.Time  `json:"updated_at" valid:"-"`
	DeletedAt      *time.Time `json:"deleted_at" valid:"-"`
}
