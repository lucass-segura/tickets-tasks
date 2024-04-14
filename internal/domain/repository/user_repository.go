package repository

import (
	"github.com/lucass-segura/tasks-tickets/internal/domain/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(userID uint) error
	FindByID(userID uint) (*model.User, error)
}

type GormUserRepository struct {
	DB *gorm.DB
}

func (ur *GormUserRepository) Create(user *model.User) error {
	return ur.DB.Create(user).Error
}

func (ur *GormUserRepository) Update(user *model.User) error {
	return ur.DB.Save(user).Error
}

func (ur *GormUserRepository) Delete(userID uint) error {
	return ur.DB.Delete(&model.User{}, userID).Error
}

func (ur *GormUserRepository) FindByID(userID uint) (*model.User, error) {
	var user model.User
	if err := ur.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
