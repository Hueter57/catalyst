package service

import "github.com/hueter57/catalyst/backend/internal/graph/model"

type TaskRepository interface {
	GetTaskByID(id string) (*model.Task, error)
	GetTasksByFilter(filter *model.TaskFilterInput, sortBy model.TaskSortInput) ([]*model.Task, error)
	UpdateTask(input model.UpdateTaskInput) (*model.Task, error)
	CreateTask(input model.CreateTaskInput) (*model.Task, error)
	DeleteTask(id string) (bool, error)
}