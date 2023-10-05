package userauth

import (
	"amazon-backend/config"
	"amazon-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type onlyEmailandPass struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user onlyEmailandPass
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	var dataFromDB models.User
	if err := config.DB.Where("email = ?", user.Email).First(&dataFromDB).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if !CheckPasswordHash(user.Password, dataFromDB.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to create Token"})
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("LoginAuthorization", tokenString, 60*60*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
