package productstock

import (
	"time"
)

//ProductStock is the definition of product stock table in database
type ProductStock struct {
	ID              uint       `gorm:"primary_key" json:"id" valid:"-"`
	Stock           int        `gorm:"not null" json:"stock" valid:"numeric,required"`
	TotalOrdered    int        `gorm:"not null" json:"total_ordered" valid:"numeric,required"`
	AverageBuyPrice float64    `gorm:"not null" json:"average_buy_price" valid:"float,required"`
	ProductID       uint       `json:"product_id" valid:"numeric,required" sql:"type:uint REFERENCES products(id)"`
	CreatedAt       time.Time  `json:"created_at" valid:"-"`
	UpdatedAt       time.Time  `json:"updated_at" valid:"-"`
	DeletedAt       *time.Time `json:"deleted_at" valid:"-"`
}
