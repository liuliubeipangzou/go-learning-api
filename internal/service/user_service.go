package service

import (
	"errors"
	"net/mail"
	"strings"

	"go-learning-api/internal/model"
	"go-learning-api/internal/repository"
)

var ErrInvalidUserInput = errors.New("invalid user input")

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers() []model.User {
	return s.repo.FindAll()
}

func (s *UserService) GetUser(id int64) (model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) CreateUser(name, email string) (model.User, error) {
	name, email, err := normalizeUserInput(name, email)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Name:  name,
		Email: email,
	}

	return s.repo.Save(user), nil
}

func (s *UserService) UpdateUser(id int64, name, email string) (model.User, error) {
	name, email, err := normalizeUserInput(name, email)
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		ID:    id,
		Name:  name,
		Email: email,
	}

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int64) error {
	return s.repo.Delete(id)
}

func normalizeUserInput(name, email string) (string, string, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name == "" || email == "" {
		return "", "", ErrInvalidUserInput
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return "", "", ErrInvalidUserInput
	}

	return name, email, nil
}
