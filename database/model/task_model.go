package model

type Task struct {
	ID			byte	   `gorm:"type:integer;gorm:primaryKey;autoIncrement"`
	Description string     `gorm:"type:text;not null"`
	Priority    string     `gorm:"type:text;not null"`
	Completed   byte        `gorm:"type:numeric;"`
	CreatedAt   int64 	   `gorm:"type:integer;not null"`
	CompletedAt int64 	   `gorm:"type:integer;"`
	UserID 		byte		
	User 		User	   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}