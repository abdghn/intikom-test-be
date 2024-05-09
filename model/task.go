package model

import (
	"time"
)

type Task struct {
	ID          int       `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	Status      string    `json:"status" gorm:"type:varchar(50);default:'pending'"`
	UserID      uint      `json:"user_id" gorm:"user_id"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      int    `json:"user_id"`
}
