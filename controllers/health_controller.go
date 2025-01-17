/**
 * @file: health_controller.go
 * @author: Krisna Pranav
 * @brief: health controller
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
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "api working successfully!",
	})
}
