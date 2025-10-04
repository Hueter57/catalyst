package service

import "github.com/hueter57/catalyst/backend/internal/graph/model"

type CategoryRepository interface {
	GetCategoryByID(id string) (*model.Category, error)
	GetCategoryList() ([]*model.Category, error)
	CreateCategory(input model.CreateCategoryInput) (*model.Category, error)
	UpdateCategory(input model.UpdateCategoryInput) (*model.Category, error)
	DeleteCategory(id string) (bool, error)
}
