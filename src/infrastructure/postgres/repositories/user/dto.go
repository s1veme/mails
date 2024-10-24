package user

import (
	"time"

	"mails/src/domain/user"

	"github.com/google/uuid"
)

type UserDTO struct {
	ID           uuid.UUID
	Name         string
	Email        string
	PasswordHash string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *UserDTO) ToDomain() *user.User {
	return &user.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromDomain(user *user.User) *UserDTO {
	return &UserDTO{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		IsActive:     user.IsActive,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
