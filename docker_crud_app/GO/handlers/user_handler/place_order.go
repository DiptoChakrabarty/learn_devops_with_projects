package userhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Place_Order(c *gin.Context) {
	var cart []models.Cart
	useridStr := c.Query("userid")
	userid, err := strconv.ParseUint(useridStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userid"})
		return
	}
	if err := config.DB.Where("user_id = ?", userid).Preload("Product").Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get cart for the user"})
		return
	}
	fmt.Print(cart)
	for _, item := range cart {
		fmt.Print("1")
		productiditer := item.ProductID
		fmt.Print(productiditer)
		var placeorder models.Orders
		var product models.Product
		if err := config.DB.Where("id = ?", productiditer).First(&product).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product from cart"})
			return
		}
		if product.Stock <= 0 {
			continue
		}
		placeorder = models.Orders{
			Status:    "Order Placed",
			OrderDate: time.Now(),
			UserID:    uint(userid),
			SellerID:  product.SellerID,
			ProductID: product.ID,
		}
		if err := config.DB.Create(&placeorder).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to place order"})
			return
		}
		product.Stock = product.Stock - 1
	}
	if err := config.DB.Where("user_id = ?", userid).Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove product from cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Orders placed successfully"})
}
