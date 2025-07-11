package handlers

import (
	"net/http"
	"stock-tracker/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateAlert(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var input struct {
		Symbol      string  `json:"symbol"`
		TargetPrice float64 `json:"target_price"`
		Direction   string  `json:"direction"` // "above" or "below"
	}

	if err := c.ShouldBindJSON(&input); err != nil || (input.Direction != "above" && input.Direction != "below") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	alert := models.StockAlert{
		UserID:      userID,
		Symbol:      input.Symbol,
		TargetPrice: input.TargetPrice,
		Direction:   input.Direction,
	}
	models.DB.Create(&alert)

	c.JSON(http.StatusOK, gin.H{"message": "Alert created"})
}

func GetAlerts(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var alerts []models.StockAlert
	models.DB.Where("user_id = ?", userID).Find(&alerts)
	c.JSON(http.StatusOK, alerts)
}

func DeleteAlert(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	result := models.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.StockAlert{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Alert not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alert deleted"})
}
