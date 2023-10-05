package sellerauth

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func Register(c *gin.Context) {
	var seller models.Seller
	if err := c.ShouldBindJSON(&seller); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	hashedPassword, err := HashPassword(seller.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	seller.Password = hashedPassword
	if err := config.DB.Create(&seller).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create seller"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "seller created successfully"})
}
