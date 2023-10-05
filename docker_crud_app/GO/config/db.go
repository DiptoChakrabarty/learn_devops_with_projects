package config

import (
	"amazon-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// dsn := "host=amazon-shopping-backend-db-1 user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := "postgresql://postgres:1234@localhost:5432/postgres?sslmode=disable"
	// dsn := "host=db user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Seller{}, &models.Cart{}, &models.Inventory{}, &models.EachCartItem{}, &models.Orders{}, &models.Product{}, &models.User{})
	DB = db
}
