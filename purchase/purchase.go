package purchase

import (
	"time"
)

//Purchase is the definition of purchase table in database
type Purchase struct {
	ID             uint       `gorm:"primary_key" json:"id" valid:"-"`
	NumberOrdered  int        `gorm:"not null" json:"number_ordered" valid:"numeric,required"`
	NumberReceived int        `gorm:"not null" json:"number_received" valid:"numeric,required"`
	BuyPrice       int        `gorm:"not null" json:"buy_price" valid:"numeric,required"`
	TotalPrice     int        `gorm:"not null" json:"total_price" valid:"numeric"`
	ReceiptNumber  string     `gorm:"type:varchar(100)" json:"receipt_number" valid:"-"`
	ProductID      uint       `json:"product_id" valid:"numeric,required" sql:"type:uint REFERENCES products(id)"`
	CreatedAt      time.Time  `json:"created_at" valid:"-"`
	UpdatedAt      time.Time  `json:"updated_at" valid:"-"`
	DeletedAt      *time.Time `json:"deleted_at" valid:"-"`
}
