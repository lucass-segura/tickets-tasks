package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	FullName string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Age      uint   `gorm:"not null"`
	Role     uint   `gorm:"not null"`
	Sector   uint   `gorm:"not null"`
}
