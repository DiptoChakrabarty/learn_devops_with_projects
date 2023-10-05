package models

import (
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	ProductID uint    `json:"productid"`
}
