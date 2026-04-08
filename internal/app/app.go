package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-learning-api/internal/config"
	"go-learning-api/internal/handler"
	"go-learning-api/internal/repository"
	"go-learning-api/internal/router"
	"go-learning-api/internal/service"
)

type App struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

func New(cfg config.Config) *App {
	userRepository := repository.NewInMemoryUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	httpServer := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           router.New(userHandler),
		ReadHeaderTimeout: cfg.ReadTimeout,
		ReadTimeout:       cfg.ReadTimeout,
		WriteTimeout:      cfg.WriteTimeout,
		IdleTimeout:       cfg.IdleTimeout,
	}

	return &App{
		server:          httpServer,
		shutdownTimeout: cfg.ShutdownTimeout,
	}
}

func (a *App) Run() error {
	serverErr := make(chan error, 1)

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			serverErr <- err
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-stop:
		ctx, cancel := context.WithTimeout(context.Background(), a.shutdownTimeout)
		defer cancel()
		return a.server.Shutdown(ctx)
	}
}
