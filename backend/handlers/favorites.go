package handlers

import (
	"net/http"
	"stock-tracker/backend/models"

	"github.com/gin-gonic/gin"
)

func GetFavorites(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var favorites []models.FavoriteStock
	models.DB.Where("user_id = ?", userID).Find(&favorites)
	c.JSON(http.StatusOK, favorites)
}

func SetFavorites(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var input []struct {
		Symbol        string `json:"symbol"`
		DisplayName   string `json:"display_name"`
		Color         string `json:"color"`
		IsMarketIndex bool   `json:"is_market_index"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Delete existing favorites
	models.DB.Where("user_id = ?", userID).Delete(&models.FavoriteStock{})

	// Insert new favorites
	for _, fav := range input {
		stock := models.FavoriteStock{
			UserID:        userID,
			Symbol:        fav.Symbol,
			DisplayName:   fav.DisplayName,
			Color:         fav.Color,
			IsMarketIndex: fav.IsMarketIndex,
		}
		models.DB.Create(&stock)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Favorites updated"})
}
