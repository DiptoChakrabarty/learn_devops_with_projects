package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	OrderDate time.Time `json:"orderdate"`
	Status    string    `json:"status"`
	UserID    uint      `json:"userid"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	SellerID  uint      `json:"sellerid"`
	Seller    Seller    `gorm:"foreignKey:SellerID" json:"seller"`
	ProductID uint      `json:"productid"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
}
