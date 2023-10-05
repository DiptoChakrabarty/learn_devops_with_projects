package sellerhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Get_Product(c *gin.Context) {
	var products []models.Product
	selleridStr := c.Query("sellerid")
	sellerid, err := strconv.ParseUint(selleridStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sellerid"})
		return
	}
	query := config.DB
	query = query.Where("seller_id = ?", sellerid)
	query.Preload("Seller").Find(&products)
	c.JSON(http.StatusCreated, &products)
}
