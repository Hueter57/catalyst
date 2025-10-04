package service

type Service struct {
	r Repository
}

type Repository interface {
	TaskRepository
	CategoryRepository
	UserRepository
}

func New(r Repository) *Service {
	return &Service{
		r: r,
	}
}
