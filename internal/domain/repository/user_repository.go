package repository

import (
	"github.com/lucass-segura/tasks-tickets/internal/domain/model"
	"gorm.io/gorm"
)

// Operaciones que se pueden realizar con un usuario
type UserRepository interface {
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error
	DeleteUser(userID uint) error
	GetUserByID(userID uint) (*model.User, error)
	GetUsers() ([]*model.User, error)
}

// Implementación de UserRepository
type GormUserRepository struct {
	DB *gorm.DB
}

// Implementación de los métodos de UserRepository
func (ur *GormUserRepository) CreateUser(user *model.User) error {
	return ur.DB.Create(user).Error
}

func (ur *GormUserRepository) UpdateUser(user *model.User) error {
	return ur.DB.Save(user).Error
}

func (ur *GormUserRepository) DeleteUser(userID uint) error {
	return ur.DB.Delete(&model.User{}, userID).Error
}

func (ur *GormUserRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	if err := ur.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *GormUserRepository) GetUserByID(userID uint) (*model.User, error) {
	var user model.User
	if err := ur.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
