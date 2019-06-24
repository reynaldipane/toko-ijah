package product

import (
	"time"
)

//Product is the definition of product table in database
type Product struct {
	ID        uint       `gorm:"primary_key" json:"id" valid:"-"`
	Sku       string     `gorm:"type:varchar(100);unique;not null" json:"sku" valid:"required"`
	Name      string     `gorm:"type:varchar(100);not null" json:"name" valid:"required"`
	Size      string     `gorm:"type:varchar(5)" json:"size" valid:"alpha,required"`
	Color     string     `gorm:"type:varchar(100)" json:"color" valid:"alpha,required"`
	CreatedAt time.Time  `json:"created_at" valid:"-"`
	UpdatedAt time.Time  `json:"updated_at" valid:"-"`
	DeletedAt *time.Time `json:"deleted_at" valid:"-"`
}
