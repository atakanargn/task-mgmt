package handlers

import (
	"net/http"

	"github.com/atakanargn/task-mgmt/backend/database"
	"github.com/atakanargn/task-mgmt/backend/models"
	"github.com/gin-gonic/gin"
)

// Panoları ve içindeki görevleri listeler
func GetBoards(c *gin.Context) {
	var boards []models.Board
	// Preload("Tasks") ile her panonun görevlerini de birlikte çekiyoruz.
	database.DB.Model(&models.Board{}).Preload("Tasks").Order("created_at asc").Find(&boards)
	c.JSON(http.StatusOK, boards)
}

// Yeni bir pano oluşturur
func CreateBoard(c *gin.Context) {
	var board models.Board
	if err := c.ShouldBindJSON(&board); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&board)
	c.JSON(http.StatusCreated, board)
}
