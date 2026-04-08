package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"go-learning-api/internal/repository"
	"go-learning-api/internal/response"
	"go-learning-api/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

type createUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type updateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Health(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, _ *http.Request) {
	users := h.userService.ListUsers()
	response.JSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.userService.GetUser(id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, "user not found")
			return
		}

		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest

	if err := decodeJSON(r, &req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.userService.CreateUser(req.Name, req.Email)
	if err != nil {
		if errors.Is(err, service.ErrInvalidUserInput) {
			response.Error(w, http.StatusBadRequest, "name is required and email must be valid")
			return
		}

		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	var req updateUserRequest
	if err := decodeJSON(r, &req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.userService.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrInvalidUserInput):
			response.Error(w, http.StatusBadRequest, "name is required and email must be valid")
			return
		case errors.Is(err, repository.ErrUserNotFound):
			response.Error(w, http.StatusNotFound, "user not found")
			return
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
			return
		}
	}

	response.JSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := h.userService.DeleteUser(id); err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			response.Error(w, http.StatusNotFound, "user not found")
			return
		}

		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func decodeJSON(r *http.Request, dst any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(dst)
}
