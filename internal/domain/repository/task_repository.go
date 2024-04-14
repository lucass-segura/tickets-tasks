package repository

import (
	"github.com/lucass-segura/tasks-tickets/internal/domain/model"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(user *model.Task) error
	Update(user *model.Task) error
	Delete(user *model.Task) error
	FindByID(userID uint) (*model.Task, error)
}

type GormTaskRepository struct {
	DB *gorm.DB
}

func (ts *GormTaskRepository) Create(task *model.Task) error {
	return ts.DB.Create(task).Error
}

func (ts *GormTaskRepository) Update(task *model.Task) error {
	return ts.DB.Save(task).Error
}

func (ts *GormTaskRepository) Delete(task *model.Task) error {
	return ts.DB.Delete(task).Error
}

func (ts *GormTaskRepository) FindByID(taskID uint) (*model.Task, error) {
	var task model.Task
	if err := ts.DB.First(&task, taskID).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
