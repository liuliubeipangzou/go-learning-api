package repository

import (
	"errors"
	"sort"
	"sync"

	"go-learning-api/internal/model"
)

var ErrUserNotFound = errors.New("user not found")

type UserRepository interface {
	FindAll() []model.User
	FindByID(id int64) (model.User, error)
	Save(user model.User) model.User
	Update(user model.User) (model.User, error)
	Delete(id int64) error
}

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	nextID int64
	users  map[int64]model.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	repo := &InMemoryUserRepository{
		nextID: 3,
		users: map[int64]model.User{
			1: {ID: 1, Name: "Tom", Email: "tom@example.com"},
			2: {ID: 2, Name: "Jerry", Email: "jerry@example.com"},
		},
	}

	return repo
}

func (r *InMemoryUserRepository) FindAll() []model.User {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.User, 0, len(r.users))
	for _, user := range r.users {
		result = append(result, user)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func (r *InMemoryUserRepository) FindByID(id int64) (model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return model.User{}, ErrUserNotFound
	}

	return user, nil
}

func (r *InMemoryUserRepository) Save(user model.User) model.User {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	r.nextID++
	r.users[user.ID] = user

	return user
}

func (r *InMemoryUserRepository) Update(user model.User) (model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[user.ID]; !ok {
		return model.User{}, ErrUserNotFound
	}

	r.users[user.ID] = user

	return user, nil
}

func (r *InMemoryUserRepository) Delete(id int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return ErrUserNotFound
	}

	delete(r.users, id)

	return nil
}
