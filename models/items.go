package models

import (
	"time"
)

type Items struct {
	Item_Id uint `gorm:"primaryKey"`
	Item_Code int `json:"itemCode"`
	Description string `gorm:"type:varchar" json:"description"`
	Quantity uint `gorm:"not null" json:"quantity"`
	Order_Id uint 
	Orders Orders `gorm:"foreignKey:Order_Id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}