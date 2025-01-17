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
	"github.com/gin-gonic/gin"
	"github.com/krishpranav/elevatenow_backend/controllers"
)

func main() {
	app := gin.Default()

	/// POST
	app.POST("/register", controllers.RegisterUser)

	/// GET
	app.GET("/healthcheck", controllers.CheckHealth)
	app.GET("/courses", controllers.GetCourses)

	app.Run()
}
