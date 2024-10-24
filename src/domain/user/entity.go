package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewUser(id uuid.UUID, name, email, passwordHash string, isActive bool, createdAt, updatedAt time.Time) *User {
	return &User{
		ID:           id,
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		IsActive:     isActive,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
