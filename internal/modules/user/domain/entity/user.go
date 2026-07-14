package entity

import (
	"errors"
	"strings"
	"time"

	sharedentity "github.com/panuwat39/go-hexagonal-template/internal/shared/domain/entity"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
	ErrInvalidName  = errors.New("name must be at least 2 characters")
)

type User struct {
	sharedentity.BaseEntity

	id    string
	email string
	name  string
}

func NewUser(id string, email string, name string, now time.Time, actorID string) (*User, error) {
	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	normalizedName := strings.TrimSpace(name)

	if !strings.Contains(normalizedEmail, "@") {
		return nil, ErrInvalidEmail
	}

	if len(normalizedName) < 2 {
		return nil, ErrInvalidName
	}

	return &User{
		BaseEntity: sharedentity.NewBaseEntity(now, normalizeActorID(actorID)),
		id:         id,
		email:      normalizedEmail,
		name:       normalizedName,
	}, nil
}

func RestoreUser(
	id string,
	email string,
	name string,
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
	createdBy string,
	updatedBy string,
	deletedBy *string,
) *User {
	return &User{
		BaseEntity: sharedentity.RestoreBaseEntity(
			createdAt,
			updatedAt,
			deletedAt,
			createdBy,
			updatedBy,
			deletedBy,
		),
		id:    id,
		email: strings.ToLower(strings.TrimSpace(email)),
		name:  strings.TrimSpace(name),
	}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Rename(name string, now time.Time, actorID string) error {
	normalizedName := strings.TrimSpace(name)

	if len(normalizedName) < 2 {
		return ErrInvalidName
	}

	u.name = normalizedName
	u.Touch(now, normalizeActorID(actorID))

	return nil
}

func (u *User) Delete(now time.Time, actorID string) {
	u.MarkDeleted(now, normalizeActorID(actorID))
}

func normalizeActorID(actorID string) string {
	normalizedActorID := strings.TrimSpace(actorID)

	if normalizedActorID == "" {
		return "system"
	}

	return normalizedActorID
}
