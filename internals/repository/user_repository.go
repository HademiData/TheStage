package repository

import "stage/internals/domain"

type UserRepository interface {
	Create(user domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
