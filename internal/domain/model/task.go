package model

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Images      string `gorm:"type:varchar(250)[]"`
	Status      uint   `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
}
