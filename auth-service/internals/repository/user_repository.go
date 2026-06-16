package repository

import "auth-service/internals/domain"

type UserRepository interface {
	Create(user domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
