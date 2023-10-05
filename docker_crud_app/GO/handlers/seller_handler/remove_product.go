package sellerhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RemoveProduct(c *gin.Context) {
	var product models.Product
	productidStr := c.Query("productid")
	productid, err := strconv.ParseUint(productidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productid"})
		return
	}
	if err := config.DB.Where("ID = ?", productid).Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Delete product"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Product Deleted successfully"})
	return
}
