package repository

import "github.com/hueter57/catalyst/backend/internal/graph/model"

func (r *Repository) GetTaskByID(id string) (*model.Task, error) {
	// TODO: implement
	// nolint nilnil
	return nil, nil
}

func (r *Repository) GetTasksByFilter(
	filter *model.TaskFilterInput,
	sortBy model.TaskSortInput,
) ([]*model.Task, error) {
	// TODO: implement
	// nolint nilnil
	return nil, nil
}

func (r *Repository) UpdateTask(input model.UpdateTaskInput) (*model.Task, error) {
	// TODO: implement
	// nolint nilnil
	return nil, nil
}

func (r *Repository) CreateTask(input model.CreateTaskInput) (*model.Task, error) {
	// TODO: implement
	// nolint nilnil
	return nil, nil
}

func (r *Repository) DeleteTask(id string) (bool, error) {
	// TODO: implement
	return false, nil
}
