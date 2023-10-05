package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	User      User    `gorm:"foreignKey:UserID" json:"user"`
	UserID    uint    `json:"userid"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	ProductID uint    `json:"productid"`
	// CartItems  []EachCartItem `gorm:"foreignKey:EachCartID" json:"cartitems"`
}
