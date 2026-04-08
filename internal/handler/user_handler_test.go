package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-learning-api/internal/handler"
	"go-learning-api/internal/repository"
	"go-learning-api/internal/router"
	"go-learning-api/internal/service"
)

func TestCreateUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(`{"name":"Alice","email":"alice@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.New(userHandler).ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, rec.Code)
	}
}

func TestUpdateUserNotFound(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/users/99", strings.NewReader(`{"name":"Bob","email":"bob@example.com"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.New(userHandler).ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	repo := repository.NewInMemoryUserRepository()
	svc := service.NewUserService(repo)
	userHandler := handler.NewUserHandler(svc)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/users/1", nil)
	rec := httptest.NewRecorder()

	router.New(userHandler).ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("expected status %d, got %d", http.StatusNoContent, rec.Code)
	}
}
