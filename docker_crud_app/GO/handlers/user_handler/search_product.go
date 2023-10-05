package userhandler

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type searching struct {
	SQ string `json:"sq"`
}

func Search_Product(c *gin.Context) {
	var searchQuery searching
	if err := c.ShouldBindJSON(&searchQuery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	var products []models.Product
	query := config.DB
	if searchQuery.SQ != "" {
		query = query.Where("name = ?", searchQuery.SQ)
	}
	query.Find(&products)
	c.JSON(http.StatusCreated, products)
}
