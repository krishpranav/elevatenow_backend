/**
 * @file: enrollment.go
 * @author: Krisna Pranav
 * @brief: enrollment
 * @version 1.0
 * @date 2025-01-17
 *
 * @copyright Copyright (c) 2024-2025 Krisna Pranav
 *
 */

package models

import "time"

type Enrollment struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id"`
	CourseID  uint      `json:"course_id"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
