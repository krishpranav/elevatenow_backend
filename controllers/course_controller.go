/**
 * @file: course_controller.go
 * @author: Krisna Pranav
 * @brief: course controller
 * @version 1.0
 * @date 2025-01-17
 *
 * @copyright Copyright (c) 2024-2025 Krisna Pranav
 *
 */

package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krishpranav/elevatenow_backend/models"
)

// / Get Courses
func GetCourses(c *gin.Context) {
	courses := []models.Course{
		{ID: 1, Title: "Introduction to Golang", Description: "Learn the basics of Go.", Category: "Programming", Level: "Beginner", CreatedAt: time.Now()},
		{ID: 2, Title: "Advanced Communication Skills", Description: "Master the art of communication.", Category: "Soft Skills", Level: "Advanced", CreatedAt: time.Now()},
	}

	c.JSON(http.StatusOK, courses)
}
