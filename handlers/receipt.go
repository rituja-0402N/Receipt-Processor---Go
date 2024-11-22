package handlers

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/storage"
	"receipt-processor/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()
	points := utils.CalculatePoints(receipt)

	storage.SaveReceipt(id, receipt, points)
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")
	points, err := storage.GetPoints(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"points": points})
}
