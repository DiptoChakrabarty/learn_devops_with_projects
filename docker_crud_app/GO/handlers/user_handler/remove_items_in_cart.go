package userhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Remove_Item_In_Cart(c *gin.Context) {
	var cart models.Cart
	productidStr := c.Query("productid")
	productid, err := strconv.ParseUint(productidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productid"})
		return
	}
	useridStr := c.Query("userid")
	userid, err := strconv.ParseUint(useridStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userid"})
		return
	}
	if err := config.DB.Where("product_id = ?", productid).Where("user_id = ?", userid).Delete(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove product from cart"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product removed successfully from cart"})
	return
}
