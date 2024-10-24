package user

import (
	"context"
	"errors"

	"mails/src/domain/user"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const tableName = "users"

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *user.User) (*user.User, error) {
	userDTO := FromDomain(user)

	query, args, err := squirrel.Insert(tableName).
		Columns("id", "name", "email", "password_hash", "is_active", "created_at", "updated_at").
		Values(userDTO.ID, userDTO.Name, userDTO.Email, userDTO.PasswordHash, userDTO.IsActive, userDTO.CreatedAt, userDTO.UpdatedAt).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return nil, err
	}

	if err := r.db.QueryRow(ctx, query, args...).Scan(&userDTO.ID); err != nil {
		return nil, err
	}

	return userDTO.ToDomain(), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	query, args, err := squirrel.Select("id, name, email, is_active, created_at, updated_at").
		From(tableName).
		Where(squirrel.Eq{"email": email}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var userDTO UserDTO
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&userDTO.ID, &userDTO.Name, &userDTO.Email,
		&userDTO.IsActive, &userDTO.CreatedAt, &userDTO.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}

	return userDTO.ToDomain(), nil
}
