package model

type User struct {
	ID byte `gorm:"primaryKey"`
	Name string `gorm:"type:varchar;not null;size:100"`
	Email *string 
	Tasks []Task
}