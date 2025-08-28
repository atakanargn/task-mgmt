package handlers

import (
	"net/http"

	"github.com/atakanargn/task-mgmt/backend/database"
	"github.com/atakanargn/task-mgmt/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	database.DB.Order("created_at desc").Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Gelen JSON'da board_id olmalı.
	if task.BoardID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "board_id is required"})
		return
	}
	database.DB.Create(&task)
	c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var task models.Task
	if err := database.DB.First(&task, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input struct {
		Title     *string    `json:"title"`
		Content   *string    `json:"content"`
		Completed *bool      `json:"completed"`
		BoardID   *uuid.UUID `json:"board_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Title != nil {
		task.Title = *input.Title
	}
	if input.Content != nil {
		task.Content = *input.Content
	}
	if input.Completed != nil {
		task.Completed = *input.Completed
	}
	if input.BoardID != nil {
		task.BoardID = *input.BoardID
	}

	database.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

func DeleteTask(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var task models.Task
	if err := database.DB.First(&task, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	database.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
