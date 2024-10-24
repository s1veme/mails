package user

import (
	"context"

	"mails/src/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *user.User) (*user.User, error)
	FindByEmail(ctx context.Context, email string) (*user.User, error)
}
