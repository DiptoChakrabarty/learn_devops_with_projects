package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name     string `json:"name"`
	Stock    uint    `json:"stock"`
	Price    uint    `json:"price"`
	SellerID uint   `json:"sellerid"` // Define the foreign key here
	Seller   Seller `gorm:"foreignKey:SellerID" json:"seller"`
	Category string `json:"category"`
}
