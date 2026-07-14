package repository

import (
	"context"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/entity"
)

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
