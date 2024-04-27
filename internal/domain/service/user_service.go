package service

import (
	"github.com/lucass-segura/tasks-tickets/internal/domain/model"
	"github.com/lucass-segura/tasks-tickets/internal/domain/repository"
)

// Define las operaciones que se pueden realizar con un usuario
type UserService struct {
	UserRepository repository.UserRepository
}

// Funciones que se pueden realizar con un usuario
func (u *UserService) CreateUser(user *model.User) error {
	return u.UserRepository.CreateUser(user)
}

func (u *UserService) GetAllUsers() ([]*model.User, error) {
	return u.UserRepository.GetUsers()
}

func (u *UserService) GetUserByID(id int) (*model.User, error) {
	return u.UserRepository.GetUserByID(uint(id))
}

func (u *UserService) UpdateUser(user *model.User) error {
	return u.UserRepository.UpdateUser(user)
}

func (u *UserService) DeleteUser(user *model.User) error {
	return u.UserRepository.DeleteUser(user.ID)
}
