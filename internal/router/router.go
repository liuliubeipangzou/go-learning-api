package router

import (
	"net/http"

	"go-learning-api/internal/handler"
	"go-learning-api/internal/middleware"
)

func New(userHandler *handler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", userHandler.Health)
	mux.HandleFunc("GET /api/v1/users", userHandler.ListUsers)
	mux.HandleFunc("GET /api/v1/users/{id}", userHandler.GetUser)
	mux.HandleFunc("POST /api/v1/users", userHandler.CreateUser)
	mux.HandleFunc("PUT /api/v1/users/{id}", userHandler.UpdateUser)
	mux.HandleFunc("DELETE /api/v1/users/{id}", userHandler.DeleteUser)

	return middleware.Chain(
		mux,
		middleware.RequestLogger,
		middleware.Recoverer,
	)
}
