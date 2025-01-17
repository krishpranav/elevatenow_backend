/**
 * @file: main.go
 * @author: Krisna Pranav
 * @brief: main
 * @version 1.0
 * @date 2025-01-17
 *
 * @copyright Copyright (c) 2024-2025 Krisna Pranav
 *
 */

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", health)

	r.Run()
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
