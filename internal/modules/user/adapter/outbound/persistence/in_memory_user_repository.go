package persistence

import (
	"context"
	"strings"
	"sync"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/entity"
)

type InMemoryUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*entity.User
	emails map[string]string
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users:  make(map[string]*entity.User),
		emails: make(map[string]string),
	}
}

func (r *InMemoryUserRepository) FindByEmail(_ context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	userID, exists := r.emails[normalizedEmail]

	if !exists {
		return nil, nil
	}

	return r.users[userID], nil
}

func (r *InMemoryUserRepository) Save(_ context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID()] = user
	r.emails[user.Email()] = user.ID()

	return nil
}
