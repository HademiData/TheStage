package service

import (
	"stage/internals/domain"
	"stage/internals/repository"
	"errors"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) Register(email, password, role string) error {

	user := domain.User{
		ID:       email,
		Email:    email,
		Password: password,
		Role:     role,
	}

	return a.repo.Create(user)
}

func (a *AuthService) Login(email, password string) error {

	user, err := a.repo.FindByEmail(email)

	if err != nil {
		return err
	}
	if user.Password != password {
		return errors.New("invalid credentials")
	}

	return nil
}
