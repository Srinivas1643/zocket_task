package models

import "time"

// Task represents a task in the task management system
type Task struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	AssignedTo  int       `json:"assigned_to"`  // User ID of the assignee
	Status      string    `json:"status" gorm:"default:'pending'"` // pending, in-progress, completed
	Priority    string    `json:"priority" gorm:"default:'medium'"` // low, medium, high
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
