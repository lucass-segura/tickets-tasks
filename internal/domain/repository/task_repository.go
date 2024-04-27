package repository

import (
	"github.com/lucass-segura/tasks-tickets/internal/domain/model"
	"gorm.io/gorm"
)

// Operaciones que se pueden realizar con una tarea
type TaskRepository interface {
	CreateUser(user *model.Task) error
	UpdateUser(user *model.Task) error
	DeleteUser(user *model.Task) error
	GetUserByID(userID uint) (*model.Task, error)
}

// Implementación de TaskRepository
type GormTaskRepository struct {
	DB *gorm.DB
}

// Implementación de los métodos de TaskRepository
func (ts *GormTaskRepository) CreateUser(task *model.Task) error {
	return ts.DB.Create(task).Error
}

func (ts *GormTaskRepository) UpdateUser(task *model.Task) error {
	return ts.DB.Save(task).Error
}

func (ts *GormTaskRepository) DeleteUser(task *model.Task) error {
	return ts.DB.Delete(task).Error
}

func (ts *GormTaskRepository) GetTasks() ([]*model.Task, error) {
	var tasks []*model.Task
	if err := ts.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ts *GormTaskRepository) GetUserByID(taskID uint) (*model.Task, error) {
	var task model.Task
	if err := ts.DB.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
