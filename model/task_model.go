package model

import (
	"time"
)

type Task struct {
	ID			byte	   `gorm:"primaryKey;autoIncrement"`
	Description string     `gorm:"type:text;not null"`
	Priority    string     `gorm:"type:char;size:6"`
	Completed   bool       
	CreatedAt   time.Time  
	CompletedAt *time.Time 
	UserID 		byte
	User 		User
}