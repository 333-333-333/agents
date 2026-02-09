// internal/auth/infrastructure/repository/memory.go
package repository

import (
	"context"
	"sync"

	"api/auth/internal/auth/domain"
)

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*domain.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID.String()] = user
	return nil
}

func (r *InMemoryUserRepository) FindByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	u, ok := r.users[id.String()]
	if !ok {
		return nil, domain.ErrUserNotFound
	}
	return u, nil
}

func (r *InMemoryUserRepository) FindByEmail(ctx context.Context, email domain.Email) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, domain.ErrUserNotFound
}

func (r *InMemoryUserRepository) ExistsByEmail(ctx context.Context, email domain.Email) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, u := range r.users {
		if u.Email == email {
			return true, nil
		}
	}
	return false, nil
}
