package resolver

import "github.com/hueter57/catalyst/backend/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service *service.Service
}

func NewResolver(s *service.Service) *Resolver {
	return &Resolver{
		service: s,
		// TODO Loadersを実装する
		// Loaders: graph.NewLoaders(s),
	}
}
