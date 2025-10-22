package repository

import "hcm-be/internal/domain"

type UserRepository interface {
	FindAll() ([]domain.User, error)
	FindByID(id string) (*domain.User, error)
	Create(u domain.User) error
}
