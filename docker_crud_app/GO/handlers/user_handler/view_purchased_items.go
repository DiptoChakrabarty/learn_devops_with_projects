package userhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func View_Purchased_Items(c *gin.Context) {
	var order []models.Orders
	useridStr := c.Query("userid")
	userid, err := strconv.ParseUint(useridStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid userid"})
		return
	}
	if err := config.DB.Where("user_id = ?", userid).Preload("Product").Find(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get purchased product"})
		return
	}
	c.JSON(http.StatusCreated, order)
	return
}
