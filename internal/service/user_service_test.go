package service

import (
	"errors"
	"testing"

	"go-learning-api/internal/repository"
)

func TestCreateUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)

	user, err := svc.CreateUser("Alice", "alice@example.com")
	if err != nil {
		t.Fatalf("CreateUser returned unexpected error: %v", err)
	}

	if user.ID == 0 {
		t.Fatalf("expected generated user id, got %d", user.ID)
	}

	if user.Name != "Alice" {
		t.Fatalf("expected name Alice, got %s", user.Name)
	}
}

func TestCreateUserWithEmptyFields(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)

	_, err := svc.CreateUser("", "")
	if err == nil {
		t.Fatal("expected validation error, got nil")
	}

	if !errors.Is(err, ErrInvalidUserInput) {
		t.Fatalf("expected ErrInvalidUserInput, got %v", err)
	}
}

func TestCreateUserWithInvalidEmail(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)

	_, err := svc.CreateUser("Alice", "invalid-email")
	if !errors.Is(err, ErrInvalidUserInput) {
		t.Fatalf("expected ErrInvalidUserInput, got %v", err)
	}
}

func TestUpdateUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)

	user, err := svc.UpdateUser(1, "Tom Updated", "tom.updated@example.com")
	if err != nil {
		t.Fatalf("UpdateUser returned unexpected error: %v", err)
	}

	if user.Name != "Tom Updated" {
		t.Fatalf("expected updated name, got %s", user.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := NewUserService(repo)

	if err := svc.DeleteUser(1); err != nil {
		t.Fatalf("DeleteUser returned unexpected error: %v", err)
	}

	_, err := svc.GetUser(1)
	if !errors.Is(err, repository.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}
