package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"stocktracker/models"

	"github.com/gin-gonic/gin"
)

func GetStockData(c *gin.Context) {
	symbols := c.QueryArray("symbol")
	if len(symbols) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing symbols"})
		return
	}

	args := append([]string{"stock_fetcher.py"}, symbols...)
	out, err := exec.Command("python3", args...).Output()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get stock data"})
		return
	}

	var data interface{}
	if err := json.Unmarshal(out, &data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse stock data"})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetStockHistory(c *gin.Context) {
	symbol := c.Param("symbol")
	var history models.StockHistory
	result := models.DB.Where("symbol = ?", symbol).First(&history)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No history found"})
		return
	}

	var jsonData interface{}
	json.Unmarshal([]byte(history.JSONData), &jsonData)

	c.JSON(http.StatusOK, gin.H{
		"data":         jsonData,
		"last_updated": history.LastUpdated,
	})
}
