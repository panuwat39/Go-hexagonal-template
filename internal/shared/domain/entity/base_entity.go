package entity

import "time"

type BaseEntity struct {
	createdAt time.Time
	updatedAt time.Time
	deletedAt *time.Time

	createdBy string
	updatedBy string
	deletedBy *string
}

func NewBaseEntity(now time.Time, actorID string) BaseEntity {
	return BaseEntity{
		createdAt: now,
		updatedAt: now,
		createdBy: actorID,
		updatedBy: actorID,
	}
}

func RestoreBaseEntity(
	createdAt time.Time,
	updatedAt time.Time,
	deletedAt *time.Time,
	createdBy string,
	updatedBy string,
	deletedBy *string,
) BaseEntity {
	if updatedAt.IsZero() {
		updatedAt = createdAt
	}

	return BaseEntity{
		createdAt: createdAt,
		updatedAt: updatedAt,
		deletedAt: deletedAt,
		createdBy: createdBy,
		updatedBy: updatedBy,
		deletedBy: deletedBy,
	}
}

func (e BaseEntity) CreatedAt() time.Time {
	return e.createdAt
}

func (e BaseEntity) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e BaseEntity) DeletedAt() *time.Time {
	return e.deletedAt
}

func (e BaseEntity) CreatedBy() string {
	return e.createdBy
}

func (e BaseEntity) UpdatedBy() string {
	return e.updatedBy
}

func (e BaseEntity) DeletedBy() *string {
	return e.deletedBy
}

func (e BaseEntity) IsDeleted() bool {
	return e.deletedAt != nil
}

func (e *BaseEntity) Touch(now time.Time, actorID string) {
	e.updatedAt = now
	e.updatedBy = actorID
}

func (e *BaseEntity) MarkDeleted(now time.Time, actorID string) {
	e.deletedAt = &now
	e.deletedBy = &actorID
	e.Touch(now, actorID)
}

func (e *BaseEntity) Restore(now time.Time, actorID string) {
	e.deletedAt = nil
	e.deletedBy = nil
	e.Touch(now, actorID)
}
