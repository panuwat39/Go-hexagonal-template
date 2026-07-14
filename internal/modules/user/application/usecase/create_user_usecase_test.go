package usecase_test

import (
	"context"
	"errors"
	"strings"
	"sync"
	"testing"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/command"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/usecase"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/entity"
)

func TestCreateUserUseCaseExecuteCreatesUser(t *testing.T) {
	userRepository := newFakeUserRepository()
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)

	output, err := createUserUseCase.Execute(context.Background(), command.CreateUserCommand{
		Email: "John@Example.com",
		Name:  "John",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if output.ID == "" {
		t.Fatal("expected user ID")
	}

	if output.Email != "john@example.com" {
		t.Fatalf("expected normalized email, got %s", output.Email)
	}

	if output.Name != "John" {
		t.Fatalf("expected name John, got %s", output.Name)
	}
}

func TestCreateUserUseCaseExecuteRejectsDuplicateEmail(t *testing.T) {
	userRepository := newFakeUserRepository()
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)

	_, err := createUserUseCase.Execute(context.Background(), command.CreateUserCommand{
		Email: "john@example.com",
		Name:  "John",
	})

	if err != nil {
		t.Fatalf("expected first create to succeed, got %v", err)
	}

	_, err = createUserUseCase.Execute(context.Background(), command.CreateUserCommand{
		Email: "JOHN@example.com",
		Name:  "John Duplicate",
	})

	if !errors.Is(err, usecase.ErrUserAlreadyExists) {
		t.Fatalf("expected ErrUserAlreadyExists, got %v", err)
	}
}

type fakeUserRepository struct {
	mu     sync.RWMutex
	users  map[string]*entity.User
	emails map[string]string
}

func newFakeUserRepository() *fakeUserRepository {
	return &fakeUserRepository{
		users:  make(map[string]*entity.User),
		emails: make(map[string]string),
	}
}

func (r *fakeUserRepository) FindByEmail(_ context.Context, email string) (*entity.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	userID, exists := r.emails[normalizedEmail]

	if !exists {
		return nil, nil
	}

	return r.users[userID], nil
}

func (r *fakeUserRepository) Save(_ context.Context, user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.users[user.ID()] = user
	r.emails[user.Email()] = user.ID()

	return nil
}
