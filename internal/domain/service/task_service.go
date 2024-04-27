package service

import (
	"github.com/lucass-segura/tasks-tickets/internal/domain/model"
	"github.com/lucass-segura/tasks-tickets/internal/domain/repository"
)

type TaskService struct {
	TaskRepository repository.TaskRepository
}

func (ts *TaskService) CreateTask(task *model.Task) error {
	return ts.TaskRepository.CreateUser(task)
}

func (ts *TaskService) UpdateUser(task *model.Task) error {
	return ts.TaskRepository.UpdateUser(task)
}
