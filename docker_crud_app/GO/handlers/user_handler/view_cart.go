package userhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func View_Item_In_Cart(c *gin.Context) {
	var cart []models.Cart
	useridStr := c.Query("userid")
	userid, err := strconv.ParseUint(useridStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userid"})
		return
	}
	if err := config.DB.Where("user_id = ?", userid).Preload("Product").Find(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product from cart"})
		return
	}
	c.JSON(http.StatusCreated, cart)
	return
}
