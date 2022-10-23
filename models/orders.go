package models

import (
	"time"
)

type Orders struct {
	Order_Id uint `gorm:"primaryKey"`
	Customer_name string `gorm:"not null;type:varchar" json:"customerName"`
	OrderedAt time.Time `json:"orderedAt"`
}