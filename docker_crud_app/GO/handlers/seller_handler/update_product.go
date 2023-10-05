package sellerhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Update_Product(c *gin.Context) {
	var product models.Product
	var updatedPro models.Product
	if err := c.ShouldBindJSON(&updatedPro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	productidStr := c.Query("productid")
	productid, err := strconv.ParseUint(productidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid productid"})
		return
	}
	if err := config.DB.First(&product, productid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	if updatedPro.Name != "" {
		product.Name = updatedPro.Name
	}

	if updatedPro.Stock != 0 {
		product.Stock = updatedPro.Stock
	}

	if updatedPro.Price != 0 {
		product.Price = updatedPro.Price
	}

	if updatedPro.Category != "" {
		product.Category = updatedPro.Category
	}

	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
