package user

import (
	"context"
	"log/slog"

	"mails/src/domain/user"
)

type UserService struct {
	repository UserRepository
	logger     *slog.Logger
}

func NewUserService(repository UserRepository, logger *slog.Logger) *UserService {
	return &UserService{repository: repository, logger: logger}
}

func (s *UserService) Create(ctx context.Context, user *user.User) (*user.User, error) {
	return s.repository.Create(ctx, user)
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	return s.repository.FindByEmail(ctx, email)
}
