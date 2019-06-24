package report

import (
	"time"

	"github.com/reynaldipane/toko-ijah/product"
	productstock "github.com/reynaldipane/toko-ijah/product_stock"
)

type productStock struct {
	ID              uint            `gorm:"primary_key" json:"id" valid:"-"`
	Stock           int             `gorm:"not null" json:"stock" valid:"numeric,required"`
	TotalOrdered    int             `gorm:"not null" json:"total_ordered" valid:"numeric,required"`
	AverageBuyPrice float64         `gorm:"not null" json:"average_buy_price" valid:"float,required"`
	ProductID       uint            `json:"product_id" valid:"numeric,required" sql:"type:uint REFERENCES products(id)"`
	Product         product.Product `json:"product"`
	CreatedAt       time.Time       `json:"created_at" valid:"-"`
	UpdatedAt       time.Time       `json:"updated_at" valid:"-"`
	DeletedAt       *time.Time      `json:"deleted_at" valid:"-"`
}

type order struct {
	ID             uint                      `gorm:"primary_key" json:"id" valid:"-"`
	NumberSold     int                       `gorm:"not null" json:"number_sold" valid:"numeric,required"`
	SellPrice      int                       `gorm:"not null" json:"sell_price" valid:"numeric,required"`
	TotalPrice     int                       `gorm:"not null" json:"total_price" valid:"numeric"`
	Notes          string                    `gorm:"type:varchar(100)" json:"receipt_number" valid:"-"`
	ProductID      uint                      `json:"product_id" valid:"numeric,required" sql:"type:uint REFERENCES products(id)"`
	ProductStockID uint                      `json:"product_stock_id" valid:"numeric,required" sql:"type:uint REFERENCES product_stocks(id)"`
	Product        product.Product           `json:"product"`
	ProductStock   productstock.ProductStock `json:"product_stock"`
	CreatedAt      time.Time                 `json:"created_at" valid:"-"`
	UpdatedAt      time.Time                 `json:"updated_at" valid:"-"`
	DeletedAt      *time.Time                `json:"deleted_at" valid:"-"`
}

type productValuesReport struct {
	Sku             string
	Name            string
	Stock           int
	AverageBuyPrice float64
	TotalPrice      float64
}

type salesDetailReport struct {
	OrderID    uint
	OrderTime  time.Time
	Sku        string
	Name       string
	NumberSold int
	SellPrice  int
	TotalPrice float64
	BuyPrice   float64
	Profit     float64
}

type salesReport struct {
	TotalTurnOver    float64
	TotalGrossProfit float64
	TotalSoldUnit    int
	TotalStock       int
}
