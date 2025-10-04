package service

import "github.com/hueter57/catalyst/backend/internal/graph/model"

type UserRepository interface {
	GetUserByID(id string) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
}
