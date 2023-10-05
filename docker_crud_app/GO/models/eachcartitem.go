package models

import "gorm.io/gorm"

type EachCartItem struct {
	gorm.Model
	// Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	// ProductID uint    `json:"productid"`
	// Quantity  int     `json:"quantity"`
}
