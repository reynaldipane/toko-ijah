package purchasehistory

import (
	"time"
)

//PurchaseHistory is the definition of purchase history table in database
type PurchaseHistory struct {
	ID             uint       `gorm:"primary_key" json:"id" valid:"-"`
	NumberReceived int        `gorm:"not null" json:"number_received" valid:"numeric,required"`
	PurchaseID     uint       `json:"purchase_id" valid:"numeric,required" sql:"type:uint REFERENCES purchases(id)"`
	Notes          string     `gorm:"type:varchar(100)" json:"notes" valid:"-"`
	CreatedAt      time.Time  `json:"created_at" valid:"-"`
	UpdatedAt      time.Time  `json:"updated_at" valid:"-"`
	DeletedAt      *time.Time `json:"deleted_at" valid:"-"`
}
