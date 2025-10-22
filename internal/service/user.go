package service

import (
	"hcm-be/internal/domain"
	"hcm-be/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) List() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Get(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) Create(u domain.User) (*domain.User, error) {
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return &u, nil
}
