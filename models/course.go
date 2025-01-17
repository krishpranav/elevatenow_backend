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

package models

import "time"

type Course struct {
	ID          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Level       string    `json:"level"`
	CreatedAt   time.Time `json:"created_at"`
}
