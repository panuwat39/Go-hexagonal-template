package usecase

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/command"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/entity"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/repository"
)

var ErrUserAlreadyExists = errors.New("user already exists")

type CreateUserOutput struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserUseCase struct {
	userRepository repository.UserRepository
}

func NewCreateUserUseCase(userRepository repository.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, cmd command.CreateUserCommand) (*CreateUserOutput, error) {
	existingUser, err := uc.userRepository.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	id, err := newUserID()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	actorID := normalizeActorID(cmd.ActorID)

	user, err := entity.NewUser(id, cmd.Email, cmd.Name, now, actorID)
	if err != nil {
		return nil, err
	}

	if err := uc.userRepository.Save(ctx, user); err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		ID:        user.ID(),
		Email:     user.Email(),
		Name:      user.Name(),
		CreatedAt: user.CreatedAt(),
		UpdatedAt: user.UpdatedAt(),
	}, nil
}

func newUserID() (string, error) {
	var b [16]byte

	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}

	return hex.EncodeToString(b[:]), nil
}

func normalizeActorID(actorID string) string {
	normalizedActorID := strings.TrimSpace(actorID)

	if normalizedActorID == "" {
		return "system"
	}

	return normalizedActorID
}
