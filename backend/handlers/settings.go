package handlers

import (
	"net/http"
	"stock-tracker/backend/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetSettings(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var settings models.UserSettings

	if err := models.DB.Where("user_id = ?", userID).First(&settings).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Settings not found"})
		return
	}

	c.JSON(http.StatusOK, settings)
}

func UpdateSettings(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var input struct {
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		Email        string `json:"email"`
		Password     string `json:"password"`
		BaseCurrency string `json:"base_currency"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Update user
	var user models.User
	models.DB.First(&user, userID)

	if input.FirstName != "" {
		user.FirstName = input.FirstName
	}
	if input.LastName != "" {
		user.LastName = input.LastName
	}
	if input.Email != "" {
		user.Email = input.Email
	}
	if input.Password != "" {
		hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		user.PasswordHash = string(hash)
	}
	models.DB.Save(&user)

	// Update settings
	var settings models.UserSettings
	models.DB.Where("user_id = ?", userID).First(&settings)
	settings.BaseCurrency = input.BaseCurrency
	models.DB.Save(&settings)

	c.JSON(http.StatusOK, gin.H{"message": "Settings updated"})
}
