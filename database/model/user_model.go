package model

type User struct {
	ID byte `gorm:"primaryKey"`
	Name string `gorm:"type:text;not null"`
	Email *string 
	Tasks []Task
}