/**
 * @file: course.go
 * @author: Krisna Pranav
 * @brief: course models
 * @version 1.0
 * @date 2025-01-17
 *
 * @copyright Copyright (c) 2024-2025 Krisna Pranav
 *
 */

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishpranav/elevatenow_backend/models"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
